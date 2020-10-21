grammar RFC;
// Internet Message Format (RFC 5322)
// extended by Internationalized Email Headers (RFC 6532) with U_REST
//
// ABNF grammar extracted from https://tools.ietf.org/html/rfc5322

// -------------------------------------------------------
// -- 3.2.1. Quoted characters
// -------------------------------------------------------

quoted_pair     :   (BACKSLASH (vchar | wsp)) | obs_qp;

// -------------------------------------------------------
// -- 3.2.2. Folding White Space and Comments
// -------------------------------------------------------

fws             :   ((wsp* crlf)? wsp+) |  obs_fws;
                                        // Folding white space

ctext           :   (EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE) |          // Printable US-ASCII
                    (ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE) |          //  characters not including
                    (RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) |         //  "(", ")", or "\"
                    obs_ctext | U_REST;

//   (3) Within a 'comment', any sequence of up to 75 printable characters
//       (not containing 'linear-white-space'), that meets the syntax
//       rules in section 2, should be treated as an 'encoded-word'.
//       Otherwise it should be treated as normal comment text.

ccontent        :   encoded_word | ctext | quoted_pair | comment; // Added RFC 2047 encoded-word

comment         :   LEFT_PAREN ((fws)? ccontent)* (fws)? RIGHT_PAREN;

cfws            :   (((fws)? comment)+ (fws)?) | fws;

// -------------------------------------------------------
// -- 3.2.3. Atom
// -------------------------------------------------------

atext           :   alpha | digit |    // Printable US-ASCII
                    EXCLAMATION | HASH |        //  characters not including
                    DOLLAR | PERCENT |        //  specials.  Used for atoms.
                    AMPERSAND | APOSTROPHE |
                    ASTERISK | PLUS |
                    DASH | SLASH |
                    EQUALS | QUESTION |
                    CARAT | UNDERSCORE |
                    ACCENT | LEFT_CURLY_BRACE |
                    PIPE | RIGHT_CURLY_BRACE |
                    TILDE | U_REST;

atom            :   (cfws)? atext+ (cfws)?;

dot_atom_text   :   atext+ (PERIOD atext+)*;

// relax: trailing period
dot_atom        :   (cfws)? dot_atom_text PERIOD? (cfws)?;

// -------------------------------------------------------
// -- 3.2.4. Quoted Strings
// -------------------------------------------------------

qtext           :   EXCLAMATION |             // Printable US-ASCII
                    (HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE) |          //  characters not including
                    (RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) |         //  "\" or the quote character
                    obs_qtext | U_REST;

qcontent        :   qtext | quoted_pair;

quoted_content	:   ((fws)? qcontent)* (fws)?;
quoted_string   :   (cfws)?
                    dquote quoted_content dquote
                    (cfws)?;

// -------------------------------------------------------
// -- 3.2.5. Miscellaneous Tokens
// -------------------------------------------------------

// relax: allow @
word            :   encoded_word | atom | quoted_string | AT; // Added RFC 2047 encoded-word

//    In this case the set of characters that may be used in a "Q"-encoded
//    'encoded-word' is restricted to: <upper and lower case ASCII
//    letters, decimal digits, "!", "*", "+", "-", "/", "=", and "_"
//    (underscore, ASCII 95.)>.  An 'encoded-word' that appears within a
//    'phrase' MUST be separated from any adjacent 'word', 'text' or
//    'special' by 'linear-white-space'.

//   (2) Any header field not defined as '*text' should be parsed
//       according to the syntax rules for that header field.  However,
//       any 'word' that appears within a 'phrase' should be treated as an
//       'encoded-word' if it meets the syntax rules in section 2.
//       Otherwise it should be treated as an ordinary 'word'.

phrase          :   (word)+ | obs_phrase;

// -------------------------------------------------------
// -- 3.3. Date and Time Specification
// -------------------------------------------------------

date_time       :   ( day_of_week COMMA )? date time (cfws)?;

day_of_week     :   ((fws)? day_name) | obs_day_of_week;

day_name        :   ((CAP_M | M) (CAP_O | O) (CAP_N | N)) | ((CAP_T | T) (CAP_U | U) (CAP_E | E)) | ((CAP_W | W) (CAP_E | E) (CAP_D | D)) | ((CAP_T | T) (CAP_H | H) (CAP_U | U)) |
                    ((CAP_F | F) (CAP_R | R) (CAP_I | I)) | ((CAP_S | S) (CAP_A | A) (CAP_T | T)) | ((CAP_S | S) (CAP_U | U) (CAP_N | N));

date            :   day month year;

day             :   ((fws)? (digit  digit?) fws) | obs_day;

month           :   (((CAP_J | J) (CAP_A | A) (CAP_N | N)) | ((CAP_F | F) (CAP_E | E) (CAP_B | B)) | ((CAP_M | M) (CAP_A | A) (CAP_R | R)) | ((CAP_A | A) (CAP_P | P) (CAP_R | R)) |
                     ((CAP_M | M) (CAP_A | A) (CAP_Y | Y)) | ((CAP_J | J) (CAP_U | U) (CAP_N | N)) | ((CAP_J | J) (CAP_U | U) (CAP_L | L)) | ((CAP_A | A) (CAP_U | U) (CAP_G | G)) |
                     ((CAP_S | S) (CAP_E | E) (CAP_P | P)) | ((CAP_O | O) (CAP_C | C) (CAP_T | T)) | ((CAP_N | N) (CAP_O | O) (CAP_V | V)) | ((CAP_D | D) (CAP_E | E) (CAP_C | C)));

year            :   (fws (digit digit digit digit+) fws) | obs_year;

// relax: zone is not required assume GMT +0000
time            :   time_of_day zone?;

time_of_day     :   hour COLON minute ( COLON second )?;

// relax: single digit
hour            :   (digit digit?) | obs_hour;

//relax: single digit
minute          :   (digit digit?) | obs_minute;

//relax: single digit
second :   (digit digit?) | obs_second;

// split
zone_hour				:		digit digit;

zone_minute			:		digit digit;

zone            :   (fws sign=( PLUS | DASH ) (zone_hour zone_minute)) | obs_zone;

// -------------------------------------------------------
// -- 3.4. Address Specification
// -------------------------------------------------------

address         :   mailbox | group;

mailbox         :   addr_spec | name_addr;

// relax: allow addr_spec instead of angle_addr
// name_addr       :   (display_name)? angle_addr;
name_addr       :   (display_name)? (angle_addr | cfws addr_spec);

// relax: do not require addr_spec
angle_addr      :   ((cfws)? LESS_THAN (addr_spec)? GREATER_THAN (cfws)?) |
                    obs_angle_addr;

// relax: group list without semicolon
group           :   display_name COLON (group_list)? (SEMICOLON)? (cfws)?;

display_name    :   phrase;

mailbox_list    :   (mailbox (COMMA mailbox)*) | obs_mbox_list;
// relax: allow addr_spec

address_list    :   (address (COMMA address)*) | obs_addr_list;

group_list      :   mailbox_list | cfws | obs_group_list;

// -------------------------------------------------------
// -- 3.4.1. Addr-Spec Specification
// -------------------------------------------------------

// relax: allow local_part only, allow port
addr_spec       :   local_part (AT domain)? (COLON port cfws?)?;

// relax: new rule for port
port            :		digit+;

local_part      :   dot_atom | quoted_string | obs_local_part;

domain          :   dot_atom | domain_literal | obs_domain;

domain_literal  :   (cfws)? LEFT_BRACE ((fws)? dtext)* (fws)? RIGHT_BRACE (cfws)?;

dtext           :   (EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z) |          // Printable US-ASCII
                    (CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) |         //  characters not including
                    obs_dtext | U_REST;          //  "[", "]", or "\"

// -------------------------------------------------------
// -- 4.1. Miscellaneous Obsolete Tokens
// -------------------------------------------------------

obs_no_ws_ctl   :   (U_0001 | U_0002 | U_0003 | U_0004 | U_0005 | U_0006 | U_0007 | U_0008) |            // US-ASCII control
                    U_000B |             //  characters that do not
                    U_000C |             //  include the carriage
                    (U_000E | U_000F | U_0010 | U_0011 | U_0012 | U_0013 | U_0014 | U_0015 | U_0016 | U_0017 | U_0018 | U_0019 | U_001A | U_001B | U_001C | U_001D | U_001E | U_001F) |          //  return, line feed, and
                    U_007F;              //  white space characters

obs_ctext       :   obs_no_ws_ctl;

obs_qtext       :   obs_no_ws_ctl;

obs_qp          :   BACKSLASH (U_0000 | obs_no_ws_ctl | lf | cr);

obs_phrase      :   word (word | PERIOD | cfws)*;

// -------------------------------------------------------
// -- 4.2. Obsolete Folding White Space
// -------------------------------------------------------

obs_fws         :   ((crlf)? wsp)+;

// -------------------------------------------------------
// -- 4.3. Obsolete Date and Time
// -------------------------------------------------------

obs_day_of_week :   (cfws)? day_name (cfws)?;

obs_day         :   (cfws)? (digit digit?) (cfws)?;

obs_year        :   (cfws)? (digit digit+) (cfws)?;

// relax: single digit
obs_hour        :   (cfws)? (digit digit?) (cfws)?;

// relax: single digit
obs_minute      :   (cfws)? (digit digit?) (cfws)?;

// relax: single digit
obs_second      :   (cfws)? (digit digit?) (cfws)?;

obs_zone        :   ((CAP_U | U) (CAP_T | T)) | ((CAP_G | G) (CAP_M | M) (CAP_T | T)) |     // Universal Time
                                        // North American UT
                                        // offsets
                    ((CAP_E | E) (CAP_S | S) (CAP_T | T)) | ((CAP_E | E) (CAP_D | D) (CAP_T | T)) |    // Eastern:  - 5/ - 4
                    ((CAP_C | C) (CAP_S | S) (CAP_T | T)) | ((CAP_C | C) (CAP_D | D) (CAP_T | T)) |    // Central:  - 6/ - 5
                    ((CAP_M | M) (CAP_S | S) (CAP_T | T)) | ((CAP_M | M) (CAP_D | D) (CAP_T | T)) |    // Mountain: - 7/ - 6
                    ((CAP_P | P) (CAP_S | S) (CAP_T | T)) | ((CAP_P | P) (CAP_D | D) (CAP_T | T)) |    // Pacific:  - 8/ - 7
                                        //
                    (CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I) |          // Military zones - "A"
                    (CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z) |          // through "I" and "K"
                    (A | B | C | D | E | F | G | H | I) |         // through "Z", both
                    (K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z);          // upper and lower case

// -------------------------------------------------------
// -- 4.4. Obsolete Addressing
// -------------------------------------------------------

obs_angle_addr  :   (cfws)? LESS_THAN obs_route addr_spec GREATER_THAN (cfws)?;

obs_route       :   obs_domain_list COLON;

obs_domain_list :   (cfws | COMMA)* AT domain
                    (COMMA (cfws)? (AT domain)?)*;

obs_mbox_list   :   ((cfws)? COMMA)* mailbox (COMMA (mailbox | cfws)?)*;

obs_addr_list   :   ((cfws)? COMMA)* address (COMMA (address | cfws)?)*;

obs_group_list  :   ((cfws)? COMMA)+ (cfws)?;

obs_local_part  :   word (PERIOD word)*;

obs_domain      :   atom (PERIOD atom)*;

obs_dtext       :   obs_no_ws_ctl | quoted_pair;


// -------------------------------------------------------
// -- RFC 5234
// -- B.1. Core Rules
// -------------------------------------------------------

alpha          :  (CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z) | (A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z);   // A-Z / a-z

bit            :  ZERO | ONE;

char_1           :  (U_0001 | U_0002 | U_0003 | U_0004 | U_0005 | U_0006 | U_0007 | U_0008 | TAB | LF | U_000B | U_000C | CR | U_000E | U_000F | U_0010 | U_0011 | U_0012 | U_0013 | U_0014 | U_0015 | U_0016 | U_0017 | U_0018 | U_0019 | U_001A | U_001B | U_001C | U_001D | U_001E | U_001F | SPACE | EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE | U_007F);
                    // any 7-bit US-ASCII character,
                    //  excluding NUL

cr             :  CR;
                    // carriage return

crlf           :  cr lf;
                    // Internet standard newline

ctl            :  (U_0000 | U_0001 | U_0002 | U_0003 | U_0004 | U_0005 | U_0006 | U_0007 | U_0008 | TAB | LF | U_000B | U_000C | CR | U_000E | U_000F | U_0010 | U_0011 | U_0012 | U_0013 | U_0014 | U_0015 | U_0016 | U_0017 | U_0018 | U_0019 | U_001A | U_001B | U_001C | U_001D | U_001E | U_001F) | U_007F;
                    // controls

digit          :  (ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE);
                    // 0-9

dquote         :  QUOTE;
                    // " (Double Quote)

hexdig         :  digit | (CAP_A | A) | (CAP_B | B) | (CAP_C | C) | (CAP_D | D) | (CAP_E | E) | (CAP_F | F);

htab           :  TAB;
                    // horizontal tab

lf             :  LF;
                    // linefeed

lwsp           :  (wsp | (crlf wsp))*;
                    // Use of this linear-white-space rule
                    //  permits lines containing only white
                    //  space that are no longer legal in
                    //  mail headers and have caused
                    //  interoperability problems in other
                    //  contexts.
                    // Do not use when defining mail
                    //  headers and use with caution in
                    //  other contexts.

sp             :  SPACE;

vchar          :  (EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) | U_REST;
                    // visible (printing) characters

wsp            :  sp | htab;
                    // white space

// Multipurpose Internet Mail Extensions (RFC 2047)
//
// ABNF grammar extracted from https://tools.ietf.org/html/rfc2047

// -------------------------------------------------------
// -- 2. Syntax of encoded-words
// -------------------------------------------------------

//   An 'encoded-word' may not be more than 75 characters long, including
//   'charset', 'encoding', 'encoded-text', and delimiters.  If it is
//   desirable to encode more text than will fit in an 'encoded-word' of
//   75 characters, multiple 'encoded-word's (separated by CRLF SPACE) may
//   be used.

//   While there is no limit to the length of a multiple-line header
//   field, each line of a header field that contains one or more
//   'encoded-word's is limited to 76 characters.

//   The length restrictions are included both to ease interoperability
//   through internetwork mail gateways, and to impose a limit on the
//   amount of lookahead a header parser must employ (while looking for a
//   final ?= delimiter) before it can decide whether a token is an
//   "encoded-word" or something else.

//   DIFFERENCE: I would not impose this.
//   NOTE: Encoded-word is RFC 822 antom, not RFC 5322 atom.


//   (1) Any message or body part header field defined as '*text', or any
//       user-defined header field, should be parsed as follows: Beginning
//       at the start of the field-body and immediately following each
//       occurrence of 'linear-white-space', each sequence of up to 75
//       printable characters (not containing any 'linear-white-space')
//       should be examined to see if it is an 'encoded-word' according to
//       the syntax rules in section 2.  Any other sequence of printable
//       characters should be treated as ordinary ASCII text.

//   It is possible that an 'encoded-word' that is legal according to the
//   syntax defined in section 2, is incorrectly formed according to the
//   rules for the encoding being used.   For example:
//
//   (1) An 'encoded-word' which contains characters which are not legal
//       for a particular encoding (for example, a "-" in the "B"
//       encoding, or a SPACE or HTAB in either the "B" or "Q" encoding),
//       is incorrectly formed.
//
//   (2) Any 'encoded-word' which encodes a non-integral number of
//       characters or octets is incorrectly formed.
//
//   A mail reader need not attempt to display the text associated with an
//   'encoded-word' that is incorrectly formed.  However, a mail reader
//   MUST NOT prevent the display or handling of a message because an
//   'encoded-word' is incorrectly formed.


encoded_word           : (cfws)? (EQUALS QUESTION) charset QUESTION encoding QUESTION encoded_text (QUESTION EQUALS) (cfws)?;

// <Any CHAR except SPACE, CTLs, and especials> especials = "(" / ")" / "<" / ">" / "@" / "," / ";" / ":" / "<" / ">" / "/" / "[" / "]" / "?" / "." / "=" 
anytext_token  :  (EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE) | (ASTERISK | PLUS) | DASH | (ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE) | (CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z) | BACKSLASH | (CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE);

// <Any printable ASCII character other than "?" or SPACE> (but see "Use of encoded-words in message headers", section 5)
anytext_et     :  (EXCLAMATION | QUOTE | HASH | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN) | (AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE);

//  Both 'encoding' and 'charset' names are case-independent.  Thus the
//   charset name "ISO-8859-1" is equivalent to "iso-8859-1", and the
//   encoding named "Q" may be spelled either "Q" or "q".

charset        : token;    // see section 3

encoding       : token;   // see section 4

token          : anytext_token+;

encoded_text   : anytext_et+;

////////////////////////////////////////////////////////////////////////////////////////////
// Lexer rules generated for each distinct character in original grammar
// Simplified character names based on Unicode (http://www.unicode.org/charts/PDF/U0000.pdf)
////////////////////////////////////////////////////////////////////////////////////////////

TAB : '\u0009';
LF : '\u000A';
CR : '\u000D';
SPACE : ' ';
EXCLAMATION : '!';
QUOTE : '"';
HASH : '#';
DOLLAR : '$';
PERCENT : '%';
AMPERSAND : '&';
APOSTROPHE : '\'';
LEFT_PAREN : '(';
RIGHT_PAREN : ')';
ASTERISK : '*';
PLUS : '+';
COMMA : ',';
DASH : '-';
PERIOD : '.';
SLASH : '/';
ZERO : '0';
ONE : '1';
TWO : '2';
THREE : '3';
FOUR : '4';
FIVE : '5';
SIX : '6';
SEVEN : '7';
EIGHT : '8';
NINE : '9';
COLON : ':';
SEMICOLON : ';';
LESS_THAN : '<';
EQUALS : '=';
GREATER_THAN : '>';
QUESTION : '?';
AT : '@';
CAP_A : 'A';
CAP_B : 'B';
CAP_C : 'C';
CAP_D : 'D';
CAP_E : 'E';
CAP_F : 'F';
CAP_G : 'G';
CAP_H : 'H';
CAP_I : 'I';
CAP_J : 'J';
CAP_K : 'K';
CAP_L : 'L';
CAP_M : 'M';
CAP_N : 'N';
CAP_O : 'O';
CAP_P : 'P';
CAP_Q : 'Q';
CAP_R : 'R';
CAP_S : 'S';
CAP_T : 'T';
CAP_U : 'U';
CAP_V : 'V';
CAP_W : 'W';
CAP_X : 'X';
CAP_Y : 'Y';
CAP_Z : 'Z';
LEFT_BRACE : '[';
BACKSLASH : '\\';
RIGHT_BRACE : ']';
CARAT : '^';
UNDERSCORE : '_';
ACCENT : '`';
A : 'a';
B : 'b';
C : 'c';
D : 'd';
E : 'e';
F : 'f';
G : 'g';
H : 'h';
I : 'i';
J : 'j';
K : 'k';
L : 'l';
M : 'm';
N : 'n';
O : 'o';
P : 'p';
Q : 'q';
R : 'r';
S : 's';
T : 't';
U : 'u';
V : 'v';
W : 'w';
X : 'x';
Y : 'y';
Z : 'z';
LEFT_CURLY_BRACE : '{';
PIPE : '|';
RIGHT_CURLY_BRACE : '}';
TILDE : '~';
U_0000 : '\u0000';
U_0001 : '\u0001';
U_0002 : '\u0002';
U_0003 : '\u0003';
U_0004 : '\u0004';
U_0005 : '\u0005';
U_0006 : '\u0006';
U_0007 : '\u0007';
U_0008 : '\u0008';
U_000B : '\u000B';
U_000C : '\u000C';
U_000E : '\u000E';
U_000F : '\u000F';
U_0010 : '\u0010';
U_0011 : '\u0011';
U_0012 : '\u0012';
U_0013 : '\u0013';
U_0014 : '\u0014';
U_0015 : '\u0015';
U_0016 : '\u0016';
U_0017 : '\u0017';
U_0018 : '\u0018';
U_0019 : '\u0019';
U_001A : '\u001A';
U_001B : '\u001B';
U_001C : '\u001C';
U_001D : '\u001D';
U_001E : '\u001E';
U_001F : '\u001F';
U_007F : '\u007F';
U_REST : '\u0080'..'\uFFFF';
