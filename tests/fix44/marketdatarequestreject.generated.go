package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
	"github.com/b2broker/simplefix-go/session/messages"
)

const MsgTypeMarketDataRequestReject = "Y"

type MarketDataRequestReject struct {
	*fix.Message
}

func makeMarketDataRequestReject() *MarketDataRequestReject {
	msg := &MarketDataRequestReject{
		Message: fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeMarketDataRequestReject).
			SetBody(
				fix.NewKeyValue(FieldMDReqID, &fix.String{}),
				fix.NewKeyValue(FieldMDReqRejReason, &fix.String{}),
				NewAltMDSourceGrp().Group,
				fix.NewKeyValue(FieldText, &fix.String{}),
				fix.NewKeyValue(FieldEncodedTextLen, &fix.Int{}),
				fix.NewKeyValue(FieldEncodedText, &fix.String{}),
			),
	}

	msg.SetHeader(makeHeader().AsComponent())
	msg.SetTrailer(makeTrailer().AsComponent())

	return msg
}

func NewMarketDataRequestReject(mDReqID string) *MarketDataRequestReject {
	msg := makeMarketDataRequestReject().
		SetMDReqID(mDReqID)

	return msg
}

func ParseMarketDataRequestReject(data []byte) (*MarketDataRequestReject, error) {
	msg := fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeMarketDataRequestReject).
		SetBody(makeMarketDataRequestReject().Body()...).
		SetHeader(makeHeader().AsComponent()).
		SetTrailer(makeTrailer().AsComponent())

	if err := msg.Unmarshal(data); err != nil {
		return nil, err
	}

	return &MarketDataRequestReject{
		Message: msg,
	}, nil
}

func (marketDataRequestReject *MarketDataRequestReject) Header() *Header {
	header := marketDataRequestReject.Message.Header()

	return &Header{header}
}

func (marketDataRequestReject *MarketDataRequestReject) HeaderBuilder() messages.HeaderBuilder {
	return marketDataRequestReject.Header()
}

func (marketDataRequestReject *MarketDataRequestReject) Trailer() *Trailer {
	trailer := marketDataRequestReject.Message.Trailer()

	return &Trailer{trailer}
}

func (marketDataRequestReject *MarketDataRequestReject) MDReqID() string {
	kv := marketDataRequestReject.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataRequestReject *MarketDataRequestReject) SetMDReqID(mDReqID string) *MarketDataRequestReject {
	kv := marketDataRequestReject.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(mDReqID)
	return marketDataRequestReject
}

func (marketDataRequestReject *MarketDataRequestReject) MDReqRejReason() string {
	kv := marketDataRequestReject.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataRequestReject *MarketDataRequestReject) SetMDReqRejReason(mDReqRejReason string) *MarketDataRequestReject {
	kv := marketDataRequestReject.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(mDReqRejReason)
	return marketDataRequestReject
}

func (marketDataRequestReject *MarketDataRequestReject) AltMDSourceGrp() *AltMDSourceGrp {
	group := marketDataRequestReject.Get(2).(*fix.Group)

	return &AltMDSourceGrp{group}
}

func (marketDataRequestReject *MarketDataRequestReject) SetAltMDSourceGrp(noAltMDSource *AltMDSourceGrp) *MarketDataRequestReject {
	marketDataRequestReject.Set(2, noAltMDSource.Group)

	return marketDataRequestReject
}

func (marketDataRequestReject *MarketDataRequestReject) Text() string {
	kv := marketDataRequestReject.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataRequestReject *MarketDataRequestReject) SetText(text string) *MarketDataRequestReject {
	kv := marketDataRequestReject.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(text)
	return marketDataRequestReject
}

func (marketDataRequestReject *MarketDataRequestReject) EncodedTextLen() int {
	kv := marketDataRequestReject.Get(4)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (marketDataRequestReject *MarketDataRequestReject) SetEncodedTextLen(encodedTextLen int) *MarketDataRequestReject {
	kv := marketDataRequestReject.Get(4).(*fix.KeyValue)
	_ = kv.Load().Set(encodedTextLen)
	return marketDataRequestReject
}

func (marketDataRequestReject *MarketDataRequestReject) EncodedText() string {
	kv := marketDataRequestReject.Get(5)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataRequestReject *MarketDataRequestReject) SetEncodedText(encodedText string) *MarketDataRequestReject {
	kv := marketDataRequestReject.Get(5).(*fix.KeyValue)
	_ = kv.Load().Set(encodedText)
	return marketDataRequestReject
}
