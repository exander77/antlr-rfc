// Code generated from RFC.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RFC

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RFCListener is a complete listener for a parse tree produced by RFCParser.
type RFCListener interface {
	antlr.ParseTreeListener

	// EnterQuoted_pair is called when entering the quoted_pair production.
	EnterQuoted_pair(c *Quoted_pairContext)

	// EnterFws is called when entering the fws production.
	EnterFws(c *FwsContext)

	// EnterCtext is called when entering the ctext production.
	EnterCtext(c *CtextContext)

	// EnterCcontent is called when entering the ccontent production.
	EnterCcontent(c *CcontentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCfws is called when entering the cfws production.
	EnterCfws(c *CfwsContext)

	// EnterAtext is called when entering the atext production.
	EnterAtext(c *AtextContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterDot_atom_text is called when entering the dot_atom_text production.
	EnterDot_atom_text(c *Dot_atom_textContext)

	// EnterDot_atom is called when entering the dot_atom production.
	EnterDot_atom(c *Dot_atomContext)

	// EnterQtext is called when entering the qtext production.
	EnterQtext(c *QtextContext)

	// EnterQcontent is called when entering the qcontent production.
	EnterQcontent(c *QcontentContext)

	// EnterQuoted_content is called when entering the quoted_content production.
	EnterQuoted_content(c *Quoted_contentContext)

	// EnterQuoted_string is called when entering the quoted_string production.
	EnterQuoted_string(c *Quoted_stringContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterPhrase is called when entering the phrase production.
	EnterPhrase(c *PhraseContext)

	// EnterDate_time is called when entering the date_time production.
	EnterDate_time(c *Date_timeContext)

	// EnterDay_of_week is called when entering the day_of_week production.
	EnterDay_of_week(c *Day_of_weekContext)

	// EnterDay_name is called when entering the day_name production.
	EnterDay_name(c *Day_nameContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterTime_of_day is called when entering the time_of_day production.
	EnterTime_of_day(c *Time_of_dayContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterMinute is called when entering the minute production.
	EnterMinute(c *MinuteContext)

	// EnterSecond is called when entering the second production.
	EnterSecond(c *SecondContext)

	// EnterZone_hour is called when entering the zone_hour production.
	EnterZone_hour(c *Zone_hourContext)

	// EnterZone_minute is called when entering the zone_minute production.
	EnterZone_minute(c *Zone_minuteContext)

	// EnterZone is called when entering the zone production.
	EnterZone(c *ZoneContext)

	// EnterAddress is called when entering the address production.
	EnterAddress(c *AddressContext)

	// EnterMailbox is called when entering the mailbox production.
	EnterMailbox(c *MailboxContext)

	// EnterName_addr is called when entering the name_addr production.
	EnterName_addr(c *Name_addrContext)

	// EnterAngle_addr is called when entering the angle_addr production.
	EnterAngle_addr(c *Angle_addrContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterDisplay_name is called when entering the display_name production.
	EnterDisplay_name(c *Display_nameContext)

	// EnterMailbox_list is called when entering the mailbox_list production.
	EnterMailbox_list(c *Mailbox_listContext)

	// EnterAddress_list is called when entering the address_list production.
	EnterAddress_list(c *Address_listContext)

	// EnterGroup_list is called when entering the group_list production.
	EnterGroup_list(c *Group_listContext)

	// EnterAddr_spec is called when entering the addr_spec production.
	EnterAddr_spec(c *Addr_specContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterLocal_part is called when entering the local_part production.
	EnterLocal_part(c *Local_partContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterDomain_literal is called when entering the domain_literal production.
	EnterDomain_literal(c *Domain_literalContext)

	// EnterDtext is called when entering the dtext production.
	EnterDtext(c *DtextContext)

	// EnterObs_no_ws_ctl is called when entering the obs_no_ws_ctl production.
	EnterObs_no_ws_ctl(c *Obs_no_ws_ctlContext)

	// EnterObs_ctext is called when entering the obs_ctext production.
	EnterObs_ctext(c *Obs_ctextContext)

	// EnterObs_qtext is called when entering the obs_qtext production.
	EnterObs_qtext(c *Obs_qtextContext)

	// EnterObs_qp is called when entering the obs_qp production.
	EnterObs_qp(c *Obs_qpContext)

	// EnterObs_phrase is called when entering the obs_phrase production.
	EnterObs_phrase(c *Obs_phraseContext)

	// EnterObs_fws is called when entering the obs_fws production.
	EnterObs_fws(c *Obs_fwsContext)

	// EnterObs_day_of_week is called when entering the obs_day_of_week production.
	EnterObs_day_of_week(c *Obs_day_of_weekContext)

	// EnterObs_day is called when entering the obs_day production.
	EnterObs_day(c *Obs_dayContext)

	// EnterObs_year is called when entering the obs_year production.
	EnterObs_year(c *Obs_yearContext)

	// EnterObs_hour is called when entering the obs_hour production.
	EnterObs_hour(c *Obs_hourContext)

	// EnterObs_minute is called when entering the obs_minute production.
	EnterObs_minute(c *Obs_minuteContext)

	// EnterObs_second is called when entering the obs_second production.
	EnterObs_second(c *Obs_secondContext)

	// EnterObs_zone is called when entering the obs_zone production.
	EnterObs_zone(c *Obs_zoneContext)

	// EnterObs_angle_addr is called when entering the obs_angle_addr production.
	EnterObs_angle_addr(c *Obs_angle_addrContext)

	// EnterObs_route is called when entering the obs_route production.
	EnterObs_route(c *Obs_routeContext)

	// EnterObs_domain_list is called when entering the obs_domain_list production.
	EnterObs_domain_list(c *Obs_domain_listContext)

	// EnterObs_mbox_list is called when entering the obs_mbox_list production.
	EnterObs_mbox_list(c *Obs_mbox_listContext)

	// EnterObs_addr_list is called when entering the obs_addr_list production.
	EnterObs_addr_list(c *Obs_addr_listContext)

	// EnterObs_group_list is called when entering the obs_group_list production.
	EnterObs_group_list(c *Obs_group_listContext)

	// EnterObs_local_part is called when entering the obs_local_part production.
	EnterObs_local_part(c *Obs_local_partContext)

	// EnterObs_domain is called when entering the obs_domain production.
	EnterObs_domain(c *Obs_domainContext)

	// EnterObs_dtext is called when entering the obs_dtext production.
	EnterObs_dtext(c *Obs_dtextContext)

	// EnterAlpha is called when entering the alpha production.
	EnterAlpha(c *AlphaContext)

	// EnterBit is called when entering the bit production.
	EnterBit(c *BitContext)

	// EnterChar_1 is called when entering the char_1 production.
	EnterChar_1(c *Char_1Context)

	// EnterCr is called when entering the cr production.
	EnterCr(c *CrContext)

	// EnterCrlf is called when entering the crlf production.
	EnterCrlf(c *CrlfContext)

	// EnterCtl is called when entering the ctl production.
	EnterCtl(c *CtlContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterDquote is called when entering the dquote production.
	EnterDquote(c *DquoteContext)

	// EnterHexdig is called when entering the hexdig production.
	EnterHexdig(c *HexdigContext)

	// EnterHtab is called when entering the htab production.
	EnterHtab(c *HtabContext)

	// EnterLf is called when entering the lf production.
	EnterLf(c *LfContext)

	// EnterLwsp is called when entering the lwsp production.
	EnterLwsp(c *LwspContext)

	// EnterSp is called when entering the sp production.
	EnterSp(c *SpContext)

	// EnterVchar is called when entering the vchar production.
	EnterVchar(c *VcharContext)

	// EnterWsp is called when entering the wsp production.
	EnterWsp(c *WspContext)

	// EnterEncoded_word is called when entering the encoded_word production.
	EnterEncoded_word(c *Encoded_wordContext)

	// EnterAnytext_token is called when entering the anytext_token production.
	EnterAnytext_token(c *Anytext_tokenContext)

	// EnterAnytext_et is called when entering the anytext_et production.
	EnterAnytext_et(c *Anytext_etContext)

	// EnterCharset is called when entering the charset production.
	EnterCharset(c *CharsetContext)

	// EnterEncoding is called when entering the encoding production.
	EnterEncoding(c *EncodingContext)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterEncoded_text is called when entering the encoded_text production.
	EnterEncoded_text(c *Encoded_textContext)

	// ExitQuoted_pair is called when exiting the quoted_pair production.
	ExitQuoted_pair(c *Quoted_pairContext)

	// ExitFws is called when exiting the fws production.
	ExitFws(c *FwsContext)

	// ExitCtext is called when exiting the ctext production.
	ExitCtext(c *CtextContext)

	// ExitCcontent is called when exiting the ccontent production.
	ExitCcontent(c *CcontentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCfws is called when exiting the cfws production.
	ExitCfws(c *CfwsContext)

	// ExitAtext is called when exiting the atext production.
	ExitAtext(c *AtextContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitDot_atom_text is called when exiting the dot_atom_text production.
	ExitDot_atom_text(c *Dot_atom_textContext)

	// ExitDot_atom is called when exiting the dot_atom production.
	ExitDot_atom(c *Dot_atomContext)

	// ExitQtext is called when exiting the qtext production.
	ExitQtext(c *QtextContext)

	// ExitQcontent is called when exiting the qcontent production.
	ExitQcontent(c *QcontentContext)

	// ExitQuoted_content is called when exiting the quoted_content production.
	ExitQuoted_content(c *Quoted_contentContext)

	// ExitQuoted_string is called when exiting the quoted_string production.
	ExitQuoted_string(c *Quoted_stringContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitPhrase is called when exiting the phrase production.
	ExitPhrase(c *PhraseContext)

	// ExitDate_time is called when exiting the date_time production.
	ExitDate_time(c *Date_timeContext)

	// ExitDay_of_week is called when exiting the day_of_week production.
	ExitDay_of_week(c *Day_of_weekContext)

	// ExitDay_name is called when exiting the day_name production.
	ExitDay_name(c *Day_nameContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitTime_of_day is called when exiting the time_of_day production.
	ExitTime_of_day(c *Time_of_dayContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitMinute is called when exiting the minute production.
	ExitMinute(c *MinuteContext)

	// ExitSecond is called when exiting the second production.
	ExitSecond(c *SecondContext)

	// ExitZone_hour is called when exiting the zone_hour production.
	ExitZone_hour(c *Zone_hourContext)

	// ExitZone_minute is called when exiting the zone_minute production.
	ExitZone_minute(c *Zone_minuteContext)

	// ExitZone is called when exiting the zone production.
	ExitZone(c *ZoneContext)

	// ExitAddress is called when exiting the address production.
	ExitAddress(c *AddressContext)

	// ExitMailbox is called when exiting the mailbox production.
	ExitMailbox(c *MailboxContext)

	// ExitName_addr is called when exiting the name_addr production.
	ExitName_addr(c *Name_addrContext)

	// ExitAngle_addr is called when exiting the angle_addr production.
	ExitAngle_addr(c *Angle_addrContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitDisplay_name is called when exiting the display_name production.
	ExitDisplay_name(c *Display_nameContext)

	// ExitMailbox_list is called when exiting the mailbox_list production.
	ExitMailbox_list(c *Mailbox_listContext)

	// ExitAddress_list is called when exiting the address_list production.
	ExitAddress_list(c *Address_listContext)

	// ExitGroup_list is called when exiting the group_list production.
	ExitGroup_list(c *Group_listContext)

	// ExitAddr_spec is called when exiting the addr_spec production.
	ExitAddr_spec(c *Addr_specContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitLocal_part is called when exiting the local_part production.
	ExitLocal_part(c *Local_partContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitDomain_literal is called when exiting the domain_literal production.
	ExitDomain_literal(c *Domain_literalContext)

	// ExitDtext is called when exiting the dtext production.
	ExitDtext(c *DtextContext)

	// ExitObs_no_ws_ctl is called when exiting the obs_no_ws_ctl production.
	ExitObs_no_ws_ctl(c *Obs_no_ws_ctlContext)

	// ExitObs_ctext is called when exiting the obs_ctext production.
	ExitObs_ctext(c *Obs_ctextContext)

	// ExitObs_qtext is called when exiting the obs_qtext production.
	ExitObs_qtext(c *Obs_qtextContext)

	// ExitObs_qp is called when exiting the obs_qp production.
	ExitObs_qp(c *Obs_qpContext)

	// ExitObs_phrase is called when exiting the obs_phrase production.
	ExitObs_phrase(c *Obs_phraseContext)

	// ExitObs_fws is called when exiting the obs_fws production.
	ExitObs_fws(c *Obs_fwsContext)

	// ExitObs_day_of_week is called when exiting the obs_day_of_week production.
	ExitObs_day_of_week(c *Obs_day_of_weekContext)

	// ExitObs_day is called when exiting the obs_day production.
	ExitObs_day(c *Obs_dayContext)

	// ExitObs_year is called when exiting the obs_year production.
	ExitObs_year(c *Obs_yearContext)

	// ExitObs_hour is called when exiting the obs_hour production.
	ExitObs_hour(c *Obs_hourContext)

	// ExitObs_minute is called when exiting the obs_minute production.
	ExitObs_minute(c *Obs_minuteContext)

	// ExitObs_second is called when exiting the obs_second production.
	ExitObs_second(c *Obs_secondContext)

	// ExitObs_zone is called when exiting the obs_zone production.
	ExitObs_zone(c *Obs_zoneContext)

	// ExitObs_angle_addr is called when exiting the obs_angle_addr production.
	ExitObs_angle_addr(c *Obs_angle_addrContext)

	// ExitObs_route is called when exiting the obs_route production.
	ExitObs_route(c *Obs_routeContext)

	// ExitObs_domain_list is called when exiting the obs_domain_list production.
	ExitObs_domain_list(c *Obs_domain_listContext)

	// ExitObs_mbox_list is called when exiting the obs_mbox_list production.
	ExitObs_mbox_list(c *Obs_mbox_listContext)

	// ExitObs_addr_list is called when exiting the obs_addr_list production.
	ExitObs_addr_list(c *Obs_addr_listContext)

	// ExitObs_group_list is called when exiting the obs_group_list production.
	ExitObs_group_list(c *Obs_group_listContext)

	// ExitObs_local_part is called when exiting the obs_local_part production.
	ExitObs_local_part(c *Obs_local_partContext)

	// ExitObs_domain is called when exiting the obs_domain production.
	ExitObs_domain(c *Obs_domainContext)

	// ExitObs_dtext is called when exiting the obs_dtext production.
	ExitObs_dtext(c *Obs_dtextContext)

	// ExitAlpha is called when exiting the alpha production.
	ExitAlpha(c *AlphaContext)

	// ExitBit is called when exiting the bit production.
	ExitBit(c *BitContext)

	// ExitChar_1 is called when exiting the char_1 production.
	ExitChar_1(c *Char_1Context)

	// ExitCr is called when exiting the cr production.
	ExitCr(c *CrContext)

	// ExitCrlf is called when exiting the crlf production.
	ExitCrlf(c *CrlfContext)

	// ExitCtl is called when exiting the ctl production.
	ExitCtl(c *CtlContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitDquote is called when exiting the dquote production.
	ExitDquote(c *DquoteContext)

	// ExitHexdig is called when exiting the hexdig production.
	ExitHexdig(c *HexdigContext)

	// ExitHtab is called when exiting the htab production.
	ExitHtab(c *HtabContext)

	// ExitLf is called when exiting the lf production.
	ExitLf(c *LfContext)

	// ExitLwsp is called when exiting the lwsp production.
	ExitLwsp(c *LwspContext)

	// ExitSp is called when exiting the sp production.
	ExitSp(c *SpContext)

	// ExitVchar is called when exiting the vchar production.
	ExitVchar(c *VcharContext)

	// ExitWsp is called when exiting the wsp production.
	ExitWsp(c *WspContext)

	// ExitEncoded_word is called when exiting the encoded_word production.
	ExitEncoded_word(c *Encoded_wordContext)

	// ExitAnytext_token is called when exiting the anytext_token production.
	ExitAnytext_token(c *Anytext_tokenContext)

	// ExitAnytext_et is called when exiting the anytext_et production.
	ExitAnytext_et(c *Anytext_etContext)

	// ExitCharset is called when exiting the charset production.
	ExitCharset(c *CharsetContext)

	// ExitEncoding is called when exiting the encoding production.
	ExitEncoding(c *EncodingContext)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitEncoded_text is called when exiting the encoded_text production.
	ExitEncoded_text(c *Encoded_textContext)
}
