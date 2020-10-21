
/*
Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by RFC 5322 and
extended by RFC 6532.
Notable divergences:
*/
package mail

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/textproto"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./parser"
)

var debug = debugT(false)

type debugT bool

func (d debugT) Printf(format string, args ...interface{}) {
	if d {
		log.Printf(format, args...)
	}
}

// A Message represents a parsed mail message.
type Message struct {
	Header Header
	Body   io.Reader
}

// ReadMessage reads a message from r.
// The headers are parsed, and the body of the message will be available
// for reading from msg.Body.
func ReadMessage(r io.Reader) (msg *Message, err error) {
	tp := textproto.NewReader(bufio.NewReader(r))

	hdr, err := tp.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}

	return &Message{
		Header: Header(hdr),
		Body:   tp.R,
	}, nil
}

// ParseDate parses an RFC 5322 date string.
func ParseDate(date string) (time time.Time, err error) {
  is := antlr.NewInputStream(date)
  lexer := parser.NewRFCLexer(is)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  parser := parser.NewRFCParser(stream)

  listener := rfcDateListener{}
  errorListener := rfcErrorListener{}

  parser.AddErrorListener(&errorListener)

  antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Date_time())

	if errorListener.counter > 0 {
		 err = fmt.Errorf("mail: %d syntax errors encountered", errorListener.counter)
	}
	return listener.dateTime, err
}

// A Header represents the key-value pairs in a mail message header.
type Header map[string][]string

// Get gets the first value associated with the given key.
// It is case insensitive; CanonicalMIMEHeaderKey is used
// to canonicalize the provided key.
// If there are no values associated with the key, Get returns "".
// To access multiple values of a key, or to use non-canonical keys,
// access the map directly.
func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

var ErrHeaderNotPresent = errors.New("mail: header not in message")

// Date parses the Date header field.
func (h Header) Date() (time.Time, error) {
	hdr := h.Get("Date")
	if hdr == "" {
		return time.Time{}, ErrHeaderNotPresent
	}
	return ParseDate(hdr)
}

// AddressList parses the named header field as a list of addresses.
func (h Header) AddressList(key string) ([]*Address, error) {
	hdr := h.Get(key)
	if hdr == "" {
		return nil, ErrHeaderNotPresent
	}
	return ParseAddressList(hdr)
}

// Address represents a single mail address.
// An address such as "Barry Gibbs <bg@example.com>" is represented
// as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.
type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(address string) (*Address, error) {
	return (&addrParser{s: address}).parseSingleAddress()
}

// ParseAddressList parses the given string as a list of addresses.
func ParseAddressList(list string) ([]*Address, error) {
	return (&addrParser{s: list}).parseAddressList()
}

// An AddressParser is an RFC 5322 address parser.
type AddressParser struct {
	// WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
	WordDecoder *mime.WordDecoder
}

// Parse parses a single RFC 5322 address of the
// form "Gogh Fir <gf@example.com>" or "foo@example.com".
func (p *AddressParser) Parse(address string) (*Address, error) {
	return (&addrParser{s: address, dec: p.WordDecoder}).parseSingleAddress()
}

// ParseList parses the given string as a list of comma-separated addresses
// of the form "Gogh Fir <gf@example.com>" or "foo@example.com".
func (p *AddressParser) ParseList(list string) ([]*Address, error) {
	return (&addrParser{s: list, dec: p.WordDecoder}).parseAddressList()
}

// String formats the address as a valid RFC 5322 address.
// If the address's name contains non-ASCII characters
// the name will be rendered according to RFC 2047.
func (a *Address) String() string {
	// Format address local@domain
	at := strings.LastIndex(a.Address, "@")
	var local, domain string
	if at < 0 {
		// This is a malformed address ("@" is required in addr-spec);
		// treat the whole address as local-part.
		local = a.Address
	} else {
		local, domain = a.Address[:at], a.Address[at+1:]
	}

	// Add quotes if needed
	quoteLocal := false
	for i, r := range local {
		if isAtext(r, false, false) {
			continue
		}
		if r == '.' {
			// Dots are okay if they are surrounded by atext.
			// We only need to check that the previous byte is
			// not a dot, and this isn't the end of the string.
			if i > 0 && local[i-1] != '.' && i < len(local)-1 {
				continue
			}
		}
		quoteLocal = true
		break
	}
	if quoteLocal {
		local = quoteString(local)

	}

	s := "<" + local + "@" + domain + ">"

	if a.Name == "" {
		return s
	}

	// If every character is printable ASCII, quoting is simple.
	allPrintable := true
	for _, r := range a.Name {
		// isWSP here should actually be isFWS,
		// but we don't support folding yet.
		if !isVchar(r) && !isWSP(r) || isMultibyte(r) {
			allPrintable = false
			break
		}
	}
	if allPrintable {
		return quoteString(a.Name) + " " + s
	}

	// Text in an encoded-word in a display-name must not contain certain
	// characters like quotes or parentheses (see RFC 2047 section 5.3).
	// When this is the case encode the name using base64 encoding.
	if strings.ContainsAny(a.Name, "\"#$%&'(),.:;<>@[]^`{|}~") {
		return mime.BEncoding.Encode("utf-8", a.Name) + " " + s
	}
	return mime.QEncoding.Encode("utf-8", a.Name) + " " + s
}

type addrParser struct {
	s   string
	dec *mime.WordDecoder // may be nil
}

func (p *addrParser) parseAddressList() (addressList []*Address, err error) {
  is := antlr.NewInputStream(p.s)
  lexer := parser.NewRFCLexer(is)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  parser := parser.NewRFCParser(stream)

	listener := rfcAddressListener{parser: p}
  errorListener := rfcErrorListener{}

  parser.AddErrorListener(&errorListener)

  antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Address_list())

	if errorListener.counter > 0 {
		 err = fmt.Errorf("mail: %d syntax errors encountered", errorListener.counter)
	}
	return listener.list, err
}

func (p *addrParser) parseSingleAddress() (address *Address, err error) {
  is := antlr.NewInputStream(p.s)
  lexer := parser.NewRFCLexer(is)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  parser := parser.NewRFCParser(stream)

	listener := rfcAddressListener{parser: p}
  errorListener := rfcErrorListener{}

  parser.AddErrorListener(&errorListener)

  antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Address())

	if errorListener.counter > 0 {
		err = fmt.Errorf("mail: %d syntax errors encountered", errorListener.counter)
	}
	return listener.list[0], err
}

func (p *addrParser) decodeRFC2047Word(s string) (word string, isEncoded bool, err error) {
	if p.dec != nil {
		word, err = p.dec.Decode(s)
	} else {
		word, err = rfc2047Decoder.Decode(s)
	}

	if err == nil {
		return word, true, nil
	}

	if _, ok := err.(charsetError); ok {
		return s, true, err
	}

	// Ignore invalid RFC 2047 encoded-word errors.
	return s, false, nil
}

var rfc2047Decoder = mime.WordDecoder{
	CharsetReader: func(charset string, input io.Reader) (io.Reader, error) {
		return nil, charsetError(charset)
	},
}

type charsetError string

func (e charsetError) Error() string {
	return fmt.Sprintf("charset not supported: %q", string(e))
}

// isAtext reports whether r is an RFC 5322 atext character.
// If dot is true, period is included.
// If permissive is true, RFC 5322 3.2.3 specials is included,
// except '<', '>', ':' and '"'.
func isAtext(r rune, dot, permissive bool) bool {
	switch r {
	case '.':
		return dot

	// RFC 5322 3.2.3. specials
	case '(', ')', '[', ']', ';', '@', '\\', ',':
		return permissive

	case '<', '>', '"', ':':
		return false
	}
	return isVchar(r)
}

// isQtext reports whether r is an RFC 5322 qtext character.
func isQtext(r rune) bool {
	// Printable US-ASCII, excluding backslash or quote.
	if r == '\\' || r == '"' {
		return false
	}
	return isVchar(r)
}

// quoteString renders a string as an RFC 5322 quoted-string.
func quoteString(s string) string {
	var buf strings.Builder
	buf.WriteByte('"')
	for _, r := range s {
		if isQtext(r) || isWSP(r) {
			buf.WriteRune(r)
		} else if isVchar(r) {
			buf.WriteByte('\\')
			buf.WriteRune(r)
		}
	}
	buf.WriteByte('"')
	return buf.String()
}

// isVchar reports whether r is an RFC 5322 VCHAR character.
func isVchar(r rune) bool {
	// Visible (printing) characters.
	return '!' <= r && r <= '~' || isMultibyte(r)
}

// isMultibyte reports whether r is a multi-byte UTF-8 character
// as supported by RFC 6532
func isMultibyte(r rune) bool {
	return r >= utf8.RuneSelf
}

// isWSP reports whether r is a WSP (white space).
// WSP is a space or horizontal tab (RFC 5234 Appendix B).
func isWSP(r rune) bool {
	return r == ' ' || r == '\t'
}

type rfcAddressListener struct {
	*parser.BaseRFCListener

	parser *addrParser
	list []*Address
	displayName []string
	addrSpec string
	collectDisplayName bool
	isPrevEncoded bool
}

//fmt.Println("C2: " + reflect.TypeOf(child).String() + " :\"" + child.(antlr.ParseTree).GetText() + "\"")
func (l *rfcAddressListener) GetText(prc antlr.ParserRuleContext) string {
	if prc.GetChildCount() == 0 {
		return ""
	}

	var s string
	for _, child := range prc.GetChildren() {
		switch child.(type) {
			case *parser.CfwsContext:
				// skip all CFWS
			case antlr.ParserRuleContext:
				s += l.GetText(child.(antlr.ParserRuleContext))
			default:
				s += child.(antlr.ParseTree).GetText()
		}
	}

	return s
}

func (l *rfcAddressListener) ExitCfws(c *parser.CfwsContext) {
}

func (l *rfcAddressListener) EnterAddress(c *parser.AddressContext) {
	l.displayName = nil
	l.addrSpec = ""
	l.collectDisplayName = false
}

func (l *rfcAddressListener) ExitAddress(c *parser.AddressContext) {
	//fmt.Printf("Adress: %v\n", l.list)
	var list = []*Address{{
      Name:		 strings.Join(l.displayName, " "),
      Address: l.addrSpec,
  }}
	l.list = append(l.list, list...)
}

func (l *rfcAddressListener) ExitAddress_list(c *parser.Address_listContext) {
	//fmt.Printf("Address-list: %v\n", l.list)
}

func (l *rfcAddressListener) EnterDisplay_name(c *parser.Display_nameContext) {
	l.collectDisplayName = true
}

func (l *rfcAddressListener) ExitDisplay_name(c *parser.Display_nameContext) {
	l.collectDisplayName = false
}

func (l *rfcAddressListener) ExitQuoted_content(c *parser.Quoted_contentContext) {
	//fmt.Println("Quoted-string :" + l.GetText(c.BaseParserRuleContext))
	if l.collectDisplayName {
		l.displayName = append(l.displayName, l.GetText(c.BaseParserRuleContext))
		l.isPrevEncoded = false
	}
}

func (l *rfcAddressListener) ExitEncoded_word(c *parser.Encoded_wordContext) {
	//fmt.Println("Encoded-word  :" + l.GetText(c.BaseParserRuleContext))
	if l.collectDisplayName {
		word, isEncoded, _ := l.parser.decodeRFC2047Word(l.GetText(c.BaseParserRuleContext))
		if l.isPrevEncoded && isEncoded {
			l.displayName[len(l.displayName)-1] += word
		} else {
			l.displayName = append(l.displayName, word)
		}
		l.isPrevEncoded = isEncoded
	}
}

func (l *rfcAddressListener) ExitAtom(c *parser.AtomContext) {
	//fmt.Println("Atom          :" + l.GetText(c.BaseParserRuleContext))
	if l.collectDisplayName {
		l.displayName = append(l.displayName, l.GetText(c.BaseParserRuleContext))
		l.isPrevEncoded = false
	}
}

func (l *rfcAddressListener) ExitAddr_spec(c *parser.Addr_specContext) {
	//fmt.Println("Addr-spec     :" + l.GetText(c.BaseParserRuleContext))
	l.addrSpec = l.GetText(c.BaseParserRuleContext)
}

type rfcErrorListener struct {
	counter int
}

func (l *rfcErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.counter += 1
}
func (l *rfcErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ambiguity")
}
func (l *rfcErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.counter -= 1
}
func (l *rfcErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//fmt.Println("context")
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

const (
	stdMonth = iota
	stdDay
	stdHour
	stdMinute
	stdSecond
	stdYear
	stdZoneHour
	stdZoneMinute
)

// match reports whether s1 and s2 match ignoring case.
// It is assumed s1 and s2 are the same length.
func match(s1, s2 string) bool {
  for i := 0; i < len(s1); i++ {
    c1 := s1[i]
    c2 := s2[i]
    if c1 != c2 {
      // Switch to lower-case; 'a'-'A' is known to be a single bit.
      c1 |= 'a' - 'A'
      c2 |= 'a' - 'A'
      if c1 != c2 || c1 < 'a' || c1 > 'z' {
        return false
      }
    }
  }
  return true
}

var errBad = errors.New("bad value for field") // placeholder not passed to user

func lookup(tab []string, val string) (int, string, error) {
  for i, v := range tab {
    if len(val) >= len(v) && match(val[0:len(v)], v) {
      return i, val[len(v):], nil
    }
  }
  return -1, val, errBad
}

type rfcDateListener struct {
	*parser.BaseRFCListener

	year       int
	month      int// = -1
	day        int// = -1
	hour       int
	min        int
	sec        int
	z					 *time.Location
	zoneHour	 int
	zoneMinute int
	zoneOffset int// = -1
	zoneName   string

	std				int

	dateTime	time.Time
}

func (l *rfcDateListener) EnterDate_time(c *parser.Date_timeContext) {
	l.month = -1
	l.day = -1
	l.zoneOffset = -1
}

func (l *rfcDateListener) ExitDate_time(c *parser.Date_timeContext) {
	l.z = time.FixedZone("", l.zoneOffset)
	var t = time.Date(l.year, time.Month(l.month), l.day, l.hour, l.min, l.sec, 0, l.z)

	fmt.Println(t)
}

func (l *rfcDateListener) ExitDay_name(c *parser.Day_nameContext) {
	// ignore
}

func (l *rfcDateListener) ExitDigit(c *parser.DigitContext) {
	fmt.Println(c.GetText())
	switch l.std {
	case stdDay:
		l.day					= l.day*10 + int(c.GetText()[0])-'0'
		l.day++
	case stdYear:
		l.year				= l.year*10 + int(c.GetText()[0])-'0'
	case stdHour:
		l.hour				= l.hour*10 + int(c.GetText()[0])-'0'
	case stdMinute:
		l.min					= l.min*10 + int(c.GetText()[0])-'0'
	case stdSecond:
		l.sec					= l.sec*10 + int(c.GetText()[0])-'0'
	case stdZoneHour:
		l.zoneHour		= l.zoneHour*10 + int(c.GetText()[0])-'0'
	case stdZoneMinute:
		l.zoneMinute	= l.zoneMinute*10 + int(c.GetText()[0])-'0'
	}
}

func (l *rfcDateListener) ExitMonth(c *parser.MonthContext) {
	l.month, _, _ = lookup(shortMonthNames, c.GetText())
	l.month++
}

func (l *rfcDateListener) EnterDay(c *parser.DayContext) {
	l.std = stdDay
}

func (l *rfcDateListener) EnterYear(c *parser.YearContext) {
	l.std = stdYear
}

func (l *rfcDateListener) EnterHour(c *parser.HourContext) {
	l.std = stdHour
}

func (l *rfcDateListener) EnterMinute(c *parser.MinuteContext) {
	l.std = stdMinute
}

func (l *rfcDateListener) EnterSecond(c *parser.SecondContext) {
	l.std = stdSecond
}

func (l *rfcDateListener) EnterZone_hour(c *parser.Zone_hourContext) {
	l.std = stdZoneHour
}

func (l *rfcDateListener) EnterZone_minute(c *parser.Zone_minuteContext) {
	l.std = stdZoneMinute
}

func (l *rfcDateListener) ExitZone(c *parser.ZoneContext) {
	l.zoneOffset = (l.zoneHour*60 + l.zoneMinute)*60
	switch c.GetSign().GetTokenType() {
	case parser.RFCParserDASH:
		l.zoneOffset = -l.zoneOffset
	case parser.RFCParserPLUS:
	default:
	}
}
