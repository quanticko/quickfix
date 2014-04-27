package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RequestForPositionsAck msg type = AO.
type RequestForPositionsAck struct {
	message.Message
}

//RequestForPositionsAckBuilder builds RequestForPositionsAck messages.
type RequestForPositionsAckBuilder struct {
	message.MessageBuilder
}

//CreateRequestForPositionsAckBuilder returns an initialized RequestForPositionsAckBuilder with specified required fields.
func CreateRequestForPositionsAckBuilder(
	posmaintrptid field.PosMaintRptID,
	posreqresult field.PosReqResult,
	posreqstatus field.PosReqStatus) RequestForPositionsAckBuilder {
	var builder RequestForPositionsAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("AO"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(posreqresult)
	builder.Body.Set(posreqstatus)
	return builder
}

//PosMaintRptID is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosMaintRptID() (*field.PosMaintRptID, errors.MessageRejectError) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosMaintRptID(f *field.PosMaintRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumPosReports is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TotalNumPosReports() (*field.TotalNumPosReports, errors.MessageRejectError) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumPosReports reads a TotalNumPosReports from RequestForPositionsAck.
func (m RequestForPositionsAck) GetTotalNumPosReports(f *field.TotalNumPosReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqResult is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqResult() (*field.PosReqResult, errors.MessageRejectError) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqResult reads a PosReqResult from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqResult(f *field.PosReqResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqStatus is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqStatus() (*field.PosReqStatus, errors.MessageRejectError) {
	f := new(field.PosReqStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqStatus reads a PosReqStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqStatus(f *field.PosReqStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from RequestForPositionsAck.
func (m RequestForPositionsAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from RequestForPositionsAck.
func (m RequestForPositionsAck) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from RequestForPositionsAck.
func (m RequestForPositionsAck) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from RequestForPositionsAck.
func (m RequestForPositionsAck) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from RequestForPositionsAck.
func (m RequestForPositionsAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqType() (*field.PosReqType, errors.MessageRejectError) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqType(f *field.PosReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}
