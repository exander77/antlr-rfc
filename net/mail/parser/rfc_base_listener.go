// Code generated from RFC.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RFC

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRFCListener is a complete listener for a parse tree produced by RFCParser.
type BaseRFCListener struct{}

var _ RFCListener = &BaseRFCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRFCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRFCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRFCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRFCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuoted_pair is called when production quoted_pair is entered.
func (s *BaseRFCListener) EnterQuoted_pair(ctx *Quoted_pairContext) {}

// ExitQuoted_pair is called when production quoted_pair is exited.
func (s *BaseRFCListener) ExitQuoted_pair(ctx *Quoted_pairContext) {}

// EnterFws is called when production fws is entered.
func (s *BaseRFCListener) EnterFws(ctx *FwsContext) {}

// ExitFws is called when production fws is exited.
func (s *BaseRFCListener) ExitFws(ctx *FwsContext) {}

// EnterCtext is called when production ctext is entered.
func (s *BaseRFCListener) EnterCtext(ctx *CtextContext) {}

// ExitCtext is called when production ctext is exited.
func (s *BaseRFCListener) ExitCtext(ctx *CtextContext) {}

// EnterCcontent is called when production ccontent is entered.
func (s *BaseRFCListener) EnterCcontent(ctx *CcontentContext) {}

// ExitCcontent is called when production ccontent is exited.
func (s *BaseRFCListener) ExitCcontent(ctx *CcontentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseRFCListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseRFCListener) ExitComment(ctx *CommentContext) {}

// EnterCfws is called when production cfws is entered.
func (s *BaseRFCListener) EnterCfws(ctx *CfwsContext) {}

// ExitCfws is called when production cfws is exited.
func (s *BaseRFCListener) ExitCfws(ctx *CfwsContext) {}

// EnterAtext is called when production atext is entered.
func (s *BaseRFCListener) EnterAtext(ctx *AtextContext) {}

// ExitAtext is called when production atext is exited.
func (s *BaseRFCListener) ExitAtext(ctx *AtextContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseRFCListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseRFCListener) ExitAtom(ctx *AtomContext) {}

// EnterDot_atom_text is called when production dot_atom_text is entered.
func (s *BaseRFCListener) EnterDot_atom_text(ctx *Dot_atom_textContext) {}

// ExitDot_atom_text is called when production dot_atom_text is exited.
func (s *BaseRFCListener) ExitDot_atom_text(ctx *Dot_atom_textContext) {}

// EnterDot_atom is called when production dot_atom is entered.
func (s *BaseRFCListener) EnterDot_atom(ctx *Dot_atomContext) {}

// ExitDot_atom is called when production dot_atom is exited.
func (s *BaseRFCListener) ExitDot_atom(ctx *Dot_atomContext) {}

// EnterQtext is called when production qtext is entered.
func (s *BaseRFCListener) EnterQtext(ctx *QtextContext) {}

// ExitQtext is called when production qtext is exited.
func (s *BaseRFCListener) ExitQtext(ctx *QtextContext) {}

// EnterQcontent is called when production qcontent is entered.
func (s *BaseRFCListener) EnterQcontent(ctx *QcontentContext) {}

// ExitQcontent is called when production qcontent is exited.
func (s *BaseRFCListener) ExitQcontent(ctx *QcontentContext) {}

// EnterQuoted_content is called when production quoted_content is entered.
func (s *BaseRFCListener) EnterQuoted_content(ctx *Quoted_contentContext) {}

// ExitQuoted_content is called when production quoted_content is exited.
func (s *BaseRFCListener) ExitQuoted_content(ctx *Quoted_contentContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BaseRFCListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BaseRFCListener) ExitQuoted_string(ctx *Quoted_stringContext) {}

// EnterWord is called when production word is entered.
func (s *BaseRFCListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseRFCListener) ExitWord(ctx *WordContext) {}

// EnterPhrase is called when production phrase is entered.
func (s *BaseRFCListener) EnterPhrase(ctx *PhraseContext) {}

// ExitPhrase is called when production phrase is exited.
func (s *BaseRFCListener) ExitPhrase(ctx *PhraseContext) {}

// EnterDate_time is called when production date_time is entered.
func (s *BaseRFCListener) EnterDate_time(ctx *Date_timeContext) {}

// ExitDate_time is called when production date_time is exited.
func (s *BaseRFCListener) ExitDate_time(ctx *Date_timeContext) {}

// EnterDay_of_week is called when production day_of_week is entered.
func (s *BaseRFCListener) EnterDay_of_week(ctx *Day_of_weekContext) {}

// ExitDay_of_week is called when production day_of_week is exited.
func (s *BaseRFCListener) ExitDay_of_week(ctx *Day_of_weekContext) {}

// EnterDay_name is called when production day_name is entered.
func (s *BaseRFCListener) EnterDay_name(ctx *Day_nameContext) {}

// ExitDay_name is called when production day_name is exited.
func (s *BaseRFCListener) ExitDay_name(ctx *Day_nameContext) {}

// EnterDate is called when production date is entered.
func (s *BaseRFCListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseRFCListener) ExitDate(ctx *DateContext) {}

// EnterDay is called when production day is entered.
func (s *BaseRFCListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BaseRFCListener) ExitDay(ctx *DayContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseRFCListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseRFCListener) ExitMonth(ctx *MonthContext) {}

// EnterYear is called when production year is entered.
func (s *BaseRFCListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BaseRFCListener) ExitYear(ctx *YearContext) {}

// EnterTime is called when production time is entered.
func (s *BaseRFCListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseRFCListener) ExitTime(ctx *TimeContext) {}

// EnterTime_of_day is called when production time_of_day is entered.
func (s *BaseRFCListener) EnterTime_of_day(ctx *Time_of_dayContext) {}

// ExitTime_of_day is called when production time_of_day is exited.
func (s *BaseRFCListener) ExitTime_of_day(ctx *Time_of_dayContext) {}

// EnterHour is called when production hour is entered.
func (s *BaseRFCListener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *BaseRFCListener) ExitHour(ctx *HourContext) {}

// EnterMinute is called when production minute is entered.
func (s *BaseRFCListener) EnterMinute(ctx *MinuteContext) {}

// ExitMinute is called when production minute is exited.
func (s *BaseRFCListener) ExitMinute(ctx *MinuteContext) {}

// EnterSecond is called when production second is entered.
func (s *BaseRFCListener) EnterSecond(ctx *SecondContext) {}

// ExitSecond is called when production second is exited.
func (s *BaseRFCListener) ExitSecond(ctx *SecondContext) {}

// EnterZone_hour is called when production zone_hour is entered.
func (s *BaseRFCListener) EnterZone_hour(ctx *Zone_hourContext) {}

// ExitZone_hour is called when production zone_hour is exited.
func (s *BaseRFCListener) ExitZone_hour(ctx *Zone_hourContext) {}

// EnterZone_minute is called when production zone_minute is entered.
func (s *BaseRFCListener) EnterZone_minute(ctx *Zone_minuteContext) {}

// ExitZone_minute is called when production zone_minute is exited.
func (s *BaseRFCListener) ExitZone_minute(ctx *Zone_minuteContext) {}

// EnterZone is called when production zone is entered.
func (s *BaseRFCListener) EnterZone(ctx *ZoneContext) {}

// ExitZone is called when production zone is exited.
func (s *BaseRFCListener) ExitZone(ctx *ZoneContext) {}

// EnterAddress is called when production address is entered.
func (s *BaseRFCListener) EnterAddress(ctx *AddressContext) {}

// ExitAddress is called when production address is exited.
func (s *BaseRFCListener) ExitAddress(ctx *AddressContext) {}

// EnterMailbox is called when production mailbox is entered.
func (s *BaseRFCListener) EnterMailbox(ctx *MailboxContext) {}

// ExitMailbox is called when production mailbox is exited.
func (s *BaseRFCListener) ExitMailbox(ctx *MailboxContext) {}

// EnterName_addr is called when production name_addr is entered.
func (s *BaseRFCListener) EnterName_addr(ctx *Name_addrContext) {}

// ExitName_addr is called when production name_addr is exited.
func (s *BaseRFCListener) ExitName_addr(ctx *Name_addrContext) {}

// EnterAngle_addr is called when production angle_addr is entered.
func (s *BaseRFCListener) EnterAngle_addr(ctx *Angle_addrContext) {}

// ExitAngle_addr is called when production angle_addr is exited.
func (s *BaseRFCListener) ExitAngle_addr(ctx *Angle_addrContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseRFCListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseRFCListener) ExitGroup(ctx *GroupContext) {}

// EnterDisplay_name is called when production display_name is entered.
func (s *BaseRFCListener) EnterDisplay_name(ctx *Display_nameContext) {}

// ExitDisplay_name is called when production display_name is exited.
func (s *BaseRFCListener) ExitDisplay_name(ctx *Display_nameContext) {}

// EnterMailbox_list is called when production mailbox_list is entered.
func (s *BaseRFCListener) EnterMailbox_list(ctx *Mailbox_listContext) {}

// ExitMailbox_list is called when production mailbox_list is exited.
func (s *BaseRFCListener) ExitMailbox_list(ctx *Mailbox_listContext) {}

// EnterAddress_list is called when production address_list is entered.
func (s *BaseRFCListener) EnterAddress_list(ctx *Address_listContext) {}

// ExitAddress_list is called when production address_list is exited.
func (s *BaseRFCListener) ExitAddress_list(ctx *Address_listContext) {}

// EnterGroup_list is called when production group_list is entered.
func (s *BaseRFCListener) EnterGroup_list(ctx *Group_listContext) {}

// ExitGroup_list is called when production group_list is exited.
func (s *BaseRFCListener) ExitGroup_list(ctx *Group_listContext) {}

// EnterAddr_spec is called when production addr_spec is entered.
func (s *BaseRFCListener) EnterAddr_spec(ctx *Addr_specContext) {}

// ExitAddr_spec is called when production addr_spec is exited.
func (s *BaseRFCListener) ExitAddr_spec(ctx *Addr_specContext) {}

// EnterPort is called when production port is entered.
func (s *BaseRFCListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseRFCListener) ExitPort(ctx *PortContext) {}

// EnterLocal_part is called when production local_part is entered.
func (s *BaseRFCListener) EnterLocal_part(ctx *Local_partContext) {}

// ExitLocal_part is called when production local_part is exited.
func (s *BaseRFCListener) ExitLocal_part(ctx *Local_partContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseRFCListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseRFCListener) ExitDomain(ctx *DomainContext) {}

// EnterDomain_literal is called when production domain_literal is entered.
func (s *BaseRFCListener) EnterDomain_literal(ctx *Domain_literalContext) {}

// ExitDomain_literal is called when production domain_literal is exited.
func (s *BaseRFCListener) ExitDomain_literal(ctx *Domain_literalContext) {}

// EnterDtext is called when production dtext is entered.
func (s *BaseRFCListener) EnterDtext(ctx *DtextContext) {}

// ExitDtext is called when production dtext is exited.
func (s *BaseRFCListener) ExitDtext(ctx *DtextContext) {}

// EnterObs_no_ws_ctl is called when production obs_no_ws_ctl is entered.
func (s *BaseRFCListener) EnterObs_no_ws_ctl(ctx *Obs_no_ws_ctlContext) {}

// ExitObs_no_ws_ctl is called when production obs_no_ws_ctl is exited.
func (s *BaseRFCListener) ExitObs_no_ws_ctl(ctx *Obs_no_ws_ctlContext) {}

// EnterObs_ctext is called when production obs_ctext is entered.
func (s *BaseRFCListener) EnterObs_ctext(ctx *Obs_ctextContext) {}

// ExitObs_ctext is called when production obs_ctext is exited.
func (s *BaseRFCListener) ExitObs_ctext(ctx *Obs_ctextContext) {}

// EnterObs_qtext is called when production obs_qtext is entered.
func (s *BaseRFCListener) EnterObs_qtext(ctx *Obs_qtextContext) {}

// ExitObs_qtext is called when production obs_qtext is exited.
func (s *BaseRFCListener) ExitObs_qtext(ctx *Obs_qtextContext) {}

// EnterObs_qp is called when production obs_qp is entered.
func (s *BaseRFCListener) EnterObs_qp(ctx *Obs_qpContext) {}

// ExitObs_qp is called when production obs_qp is exited.
func (s *BaseRFCListener) ExitObs_qp(ctx *Obs_qpContext) {}

// EnterObs_phrase is called when production obs_phrase is entered.
func (s *BaseRFCListener) EnterObs_phrase(ctx *Obs_phraseContext) {}

// ExitObs_phrase is called when production obs_phrase is exited.
func (s *BaseRFCListener) ExitObs_phrase(ctx *Obs_phraseContext) {}

// EnterObs_fws is called when production obs_fws is entered.
func (s *BaseRFCListener) EnterObs_fws(ctx *Obs_fwsContext) {}

// ExitObs_fws is called when production obs_fws is exited.
func (s *BaseRFCListener) ExitObs_fws(ctx *Obs_fwsContext) {}

// EnterObs_day_of_week is called when production obs_day_of_week is entered.
func (s *BaseRFCListener) EnterObs_day_of_week(ctx *Obs_day_of_weekContext) {}

// ExitObs_day_of_week is called when production obs_day_of_week is exited.
func (s *BaseRFCListener) ExitObs_day_of_week(ctx *Obs_day_of_weekContext) {}

// EnterObs_day is called when production obs_day is entered.
func (s *BaseRFCListener) EnterObs_day(ctx *Obs_dayContext) {}

// ExitObs_day is called when production obs_day is exited.
func (s *BaseRFCListener) ExitObs_day(ctx *Obs_dayContext) {}

// EnterObs_year is called when production obs_year is entered.
func (s *BaseRFCListener) EnterObs_year(ctx *Obs_yearContext) {}

// ExitObs_year is called when production obs_year is exited.
func (s *BaseRFCListener) ExitObs_year(ctx *Obs_yearContext) {}

// EnterObs_hour is called when production obs_hour is entered.
func (s *BaseRFCListener) EnterObs_hour(ctx *Obs_hourContext) {}

// ExitObs_hour is called when production obs_hour is exited.
func (s *BaseRFCListener) ExitObs_hour(ctx *Obs_hourContext) {}

// EnterObs_minute is called when production obs_minute is entered.
func (s *BaseRFCListener) EnterObs_minute(ctx *Obs_minuteContext) {}

// ExitObs_minute is called when production obs_minute is exited.
func (s *BaseRFCListener) ExitObs_minute(ctx *Obs_minuteContext) {}

// EnterObs_second is called when production obs_second is entered.
func (s *BaseRFCListener) EnterObs_second(ctx *Obs_secondContext) {}

// ExitObs_second is called when production obs_second is exited.
func (s *BaseRFCListener) ExitObs_second(ctx *Obs_secondContext) {}

// EnterObs_zone is called when production obs_zone is entered.
func (s *BaseRFCListener) EnterObs_zone(ctx *Obs_zoneContext) {}

// ExitObs_zone is called when production obs_zone is exited.
func (s *BaseRFCListener) ExitObs_zone(ctx *Obs_zoneContext) {}

// EnterObs_angle_addr is called when production obs_angle_addr is entered.
func (s *BaseRFCListener) EnterObs_angle_addr(ctx *Obs_angle_addrContext) {}

// ExitObs_angle_addr is called when production obs_angle_addr is exited.
func (s *BaseRFCListener) ExitObs_angle_addr(ctx *Obs_angle_addrContext) {}

// EnterObs_route is called when production obs_route is entered.
func (s *BaseRFCListener) EnterObs_route(ctx *Obs_routeContext) {}

// ExitObs_route is called when production obs_route is exited.
func (s *BaseRFCListener) ExitObs_route(ctx *Obs_routeContext) {}

// EnterObs_domain_list is called when production obs_domain_list is entered.
func (s *BaseRFCListener) EnterObs_domain_list(ctx *Obs_domain_listContext) {}

// ExitObs_domain_list is called when production obs_domain_list is exited.
func (s *BaseRFCListener) ExitObs_domain_list(ctx *Obs_domain_listContext) {}

// EnterObs_mbox_list is called when production obs_mbox_list is entered.
func (s *BaseRFCListener) EnterObs_mbox_list(ctx *Obs_mbox_listContext) {}

// ExitObs_mbox_list is called when production obs_mbox_list is exited.
func (s *BaseRFCListener) ExitObs_mbox_list(ctx *Obs_mbox_listContext) {}

// EnterObs_addr_list is called when production obs_addr_list is entered.
func (s *BaseRFCListener) EnterObs_addr_list(ctx *Obs_addr_listContext) {}

// ExitObs_addr_list is called when production obs_addr_list is exited.
func (s *BaseRFCListener) ExitObs_addr_list(ctx *Obs_addr_listContext) {}

// EnterObs_group_list is called when production obs_group_list is entered.
func (s *BaseRFCListener) EnterObs_group_list(ctx *Obs_group_listContext) {}

// ExitObs_group_list is called when production obs_group_list is exited.
func (s *BaseRFCListener) ExitObs_group_list(ctx *Obs_group_listContext) {}

// EnterObs_local_part is called when production obs_local_part is entered.
func (s *BaseRFCListener) EnterObs_local_part(ctx *Obs_local_partContext) {}

// ExitObs_local_part is called when production obs_local_part is exited.
func (s *BaseRFCListener) ExitObs_local_part(ctx *Obs_local_partContext) {}

// EnterObs_domain is called when production obs_domain is entered.
func (s *BaseRFCListener) EnterObs_domain(ctx *Obs_domainContext) {}

// ExitObs_domain is called when production obs_domain is exited.
func (s *BaseRFCListener) ExitObs_domain(ctx *Obs_domainContext) {}

// EnterObs_dtext is called when production obs_dtext is entered.
func (s *BaseRFCListener) EnterObs_dtext(ctx *Obs_dtextContext) {}

// ExitObs_dtext is called when production obs_dtext is exited.
func (s *BaseRFCListener) ExitObs_dtext(ctx *Obs_dtextContext) {}

// EnterAlpha is called when production alpha is entered.
func (s *BaseRFCListener) EnterAlpha(ctx *AlphaContext) {}

// ExitAlpha is called when production alpha is exited.
func (s *BaseRFCListener) ExitAlpha(ctx *AlphaContext) {}

// EnterBit is called when production bit is entered.
func (s *BaseRFCListener) EnterBit(ctx *BitContext) {}

// ExitBit is called when production bit is exited.
func (s *BaseRFCListener) ExitBit(ctx *BitContext) {}

// EnterChar_1 is called when production char_1 is entered.
func (s *BaseRFCListener) EnterChar_1(ctx *Char_1Context) {}

// ExitChar_1 is called when production char_1 is exited.
func (s *BaseRFCListener) ExitChar_1(ctx *Char_1Context) {}

// EnterCr is called when production cr is entered.
func (s *BaseRFCListener) EnterCr(ctx *CrContext) {}

// ExitCr is called when production cr is exited.
func (s *BaseRFCListener) ExitCr(ctx *CrContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BaseRFCListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BaseRFCListener) ExitCrlf(ctx *CrlfContext) {}

// EnterCtl is called when production ctl is entered.
func (s *BaseRFCListener) EnterCtl(ctx *CtlContext) {}

// ExitCtl is called when production ctl is exited.
func (s *BaseRFCListener) ExitCtl(ctx *CtlContext) {}

// EnterDigit is called when production digit is entered.
func (s *BaseRFCListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BaseRFCListener) ExitDigit(ctx *DigitContext) {}

// EnterDquote is called when production dquote is entered.
func (s *BaseRFCListener) EnterDquote(ctx *DquoteContext) {}

// ExitDquote is called when production dquote is exited.
func (s *BaseRFCListener) ExitDquote(ctx *DquoteContext) {}

// EnterHexdig is called when production hexdig is entered.
func (s *BaseRFCListener) EnterHexdig(ctx *HexdigContext) {}

// ExitHexdig is called when production hexdig is exited.
func (s *BaseRFCListener) ExitHexdig(ctx *HexdigContext) {}

// EnterHtab is called when production htab is entered.
func (s *BaseRFCListener) EnterHtab(ctx *HtabContext) {}

// ExitHtab is called when production htab is exited.
func (s *BaseRFCListener) ExitHtab(ctx *HtabContext) {}

// EnterLf is called when production lf is entered.
func (s *BaseRFCListener) EnterLf(ctx *LfContext) {}

// ExitLf is called when production lf is exited.
func (s *BaseRFCListener) ExitLf(ctx *LfContext) {}

// EnterLwsp is called when production lwsp is entered.
func (s *BaseRFCListener) EnterLwsp(ctx *LwspContext) {}

// ExitLwsp is called when production lwsp is exited.
func (s *BaseRFCListener) ExitLwsp(ctx *LwspContext) {}

// EnterSp is called when production sp is entered.
func (s *BaseRFCListener) EnterSp(ctx *SpContext) {}

// ExitSp is called when production sp is exited.
func (s *BaseRFCListener) ExitSp(ctx *SpContext) {}

// EnterVchar is called when production vchar is entered.
func (s *BaseRFCListener) EnterVchar(ctx *VcharContext) {}

// ExitVchar is called when production vchar is exited.
func (s *BaseRFCListener) ExitVchar(ctx *VcharContext) {}

// EnterWsp is called when production wsp is entered.
func (s *BaseRFCListener) EnterWsp(ctx *WspContext) {}

// ExitWsp is called when production wsp is exited.
func (s *BaseRFCListener) ExitWsp(ctx *WspContext) {}

// EnterEncoded_word is called when production encoded_word is entered.
func (s *BaseRFCListener) EnterEncoded_word(ctx *Encoded_wordContext) {}

// ExitEncoded_word is called when production encoded_word is exited.
func (s *BaseRFCListener) ExitEncoded_word(ctx *Encoded_wordContext) {}

// EnterAnytext_token is called when production anytext_token is entered.
func (s *BaseRFCListener) EnterAnytext_token(ctx *Anytext_tokenContext) {}

// ExitAnytext_token is called when production anytext_token is exited.
func (s *BaseRFCListener) ExitAnytext_token(ctx *Anytext_tokenContext) {}

// EnterAnytext_et is called when production anytext_et is entered.
func (s *BaseRFCListener) EnterAnytext_et(ctx *Anytext_etContext) {}

// ExitAnytext_et is called when production anytext_et is exited.
func (s *BaseRFCListener) ExitAnytext_et(ctx *Anytext_etContext) {}

// EnterCharset is called when production charset is entered.
func (s *BaseRFCListener) EnterCharset(ctx *CharsetContext) {}

// ExitCharset is called when production charset is exited.
func (s *BaseRFCListener) ExitCharset(ctx *CharsetContext) {}

// EnterEncoding is called when production encoding is entered.
func (s *BaseRFCListener) EnterEncoding(ctx *EncodingContext) {}

// ExitEncoding is called when production encoding is exited.
func (s *BaseRFCListener) ExitEncoding(ctx *EncodingContext) {}

// EnterToken is called when production token is entered.
func (s *BaseRFCListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseRFCListener) ExitToken(ctx *TokenContext) {}

// EnterEncoded_text is called when production encoded_text is entered.
func (s *BaseRFCListener) EnterEncoded_text(ctx *Encoded_textContext) {}

// ExitEncoded_text is called when production encoded_text is exited.
func (s *BaseRFCListener) ExitEncoded_text(ctx *Encoded_textContext) {}
