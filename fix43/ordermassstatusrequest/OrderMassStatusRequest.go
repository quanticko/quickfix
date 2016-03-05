//Package ordermassstatusrequest msg type = AF.
package ordermassstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/underlyinginstrument"
)

//Message is a OrderMassStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AF"`
	fix43.Header
	//MassStatusReqID is a required field for OrderMassStatusRequest.
	MassStatusReqID string `fix:"584"`
	//MassStatusReqType is a required field for OrderMassStatusRequest.
	MassStatusReqType int `fix:"585"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for OrderMassStatusRequest.
	Account *string `fix:"1"`
	//TradingSessionID is a non-required field for OrderMassStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	instrument.Instrument
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//Side is a non-required field for OrderMassStatusRequest.
	Side *string `fix:"54"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMassStatusReqID(v string)     { m.MassStatusReqID = v }
func (m *Message) SetMassStatusReqType(v int)      { m.MassStatusReqType = v }
func (m *Message) SetAccount(v string)             { m.Account = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *Message) SetSide(v string)                { m.Side = &v }

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
	return enum.BeginStringFIX43, "AF", r
}
