//Package securitylistrequest msg type = x.
package securitylistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
)

//Message is a SecurityListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"x"`
	fix43.Header
	//SecurityReqID is a required field for SecurityListRequest.
	SecurityReqID string `fix:"320"`
	//SecurityListRequestType is a required field for SecurityListRequest.
	SecurityListRequestType int `fix:"559"`
	//Instrument Component
	instrument.Instrument
	//Currency is a non-required field for SecurityListRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityListRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityListRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityListRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityListRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityListRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for SecurityListRequest.
	SubscriptionRequestType *string `fix:"263"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)           { m.SecurityReqID = v }
func (m *Message) SetSecurityListRequestType(v int)    { m.SecurityListRequestType = v }
func (m *Message) SetCurrency(v string)                { m.Currency = &v }
func (m *Message) SetText(v string)                    { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)             { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)             { m.EncodedText = &v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }

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
	return enum.BeginStringFIX43, "x", r
}
