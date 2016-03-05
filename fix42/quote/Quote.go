//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	fix42.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//QuoteResponseLevel is a non-required field for Quote.
	QuoteResponseLevel *int `fix:"301"`
	//TradingSessionID is a non-required field for Quote.
	TradingSessionID *string `fix:"336"`
	//Symbol is a required field for Quote.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Quote.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Quote.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Quote.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for Quote.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for Quote.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for Quote.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for Quote.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for Quote.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for Quote.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for Quote.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for Quote.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for Quote.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Quote.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for Quote.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for Quote.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for Quote.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for Quote.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for Quote.
	EncodedSecurityDesc *string `fix:"351"`
	//BidPx is a non-required field for Quote.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for Quote.
	OfferPx *float64 `fix:"133"`
	//BidSize is a non-required field for Quote.
	BidSize *float64 `fix:"134"`
	//OfferSize is a non-required field for Quote.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for Quote.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for Quote.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for Quote.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for Quote.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for Quote.
	OfferForwardPoints *float64 `fix:"191"`
	//TransactTime is a non-required field for Quote.
	TransactTime *time.Time `fix:"60"`
	//FutSettDate is a non-required field for Quote.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for Quote.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for Quote.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for Quote.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for Quote.
	Currency *string `fix:"15"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)          { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)             { m.QuoteID = v }
func (m *Message) SetQuoteResponseLevel(v int)     { m.QuoteResponseLevel = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *Message) SetSymbol(v string)              { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)            { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *Message) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *Message) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *Message) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)              { m.Issuer = &v }
func (m *Message) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *Message) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *Message) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *Message) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *Message) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *Message) SetBidPx(v float64)              { m.BidPx = &v }
func (m *Message) SetOfferPx(v float64)            { m.OfferPx = &v }
func (m *Message) SetBidSize(v float64)            { m.BidSize = &v }
func (m *Message) SetOfferSize(v float64)          { m.OfferSize = &v }
func (m *Message) SetValidUntilTime(v time.Time)   { m.ValidUntilTime = &v }
func (m *Message) SetBidSpotRate(v float64)        { m.BidSpotRate = &v }
func (m *Message) SetOfferSpotRate(v float64)      { m.OfferSpotRate = &v }
func (m *Message) SetBidForwardPoints(v float64)   { m.BidForwardPoints = &v }
func (m *Message) SetOfferForwardPoints(v float64) { m.OfferForwardPoints = &v }
func (m *Message) SetTransactTime(v time.Time)     { m.TransactTime = &v }
func (m *Message) SetFutSettDate(v string)         { m.FutSettDate = &v }
func (m *Message) SetOrdType(v string)             { m.OrdType = &v }
func (m *Message) SetFutSettDate2(v string)        { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)          { m.OrderQty2 = &v }
func (m *Message) SetCurrency(v string)            { m.Currency = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "S", r
}
