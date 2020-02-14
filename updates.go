package pulseaudio

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	SubscriptionEventSink         = 0x0000
	SubscriptionEventSource       = 0x0001
	SubscriptionEventSinkInput    = 0x0002
	SubscriptionEventSourceOutput = 0x0003
	SubscriptionEventModule       = 0x0004
	SubscriptionEventClient       = 0x0005
	SubscriptionEventSampleCache  = 0x0006
	SubscriptionEventServer       = 0x0007
	SubscriptionEventAutoload     = 0x0008
	SubscriptionEventCard         = 0x0009
	SubscriptionEventFacilityMask = 0x000F
	SubscriptionEventNew          = 0x0000
	SubscriptionEventChange       = 0x0010
	SubscriptionEventRemove       = 0x0020
	SubscriptionEventTypeMask     = 0x0030
)

type SubscriptionEvent struct {
	Type   int32
	Client int32
}

func newSubscriptionEvent(buf *bytes.Buffer) (se SubscriptionEvent, err error) {
	r := bytes.NewReader(buf.Bytes())
	_, err = r.Seek(1, io.SeekStart)
	err = binary.Read(r, binary.BigEndian, &se.Type)
	_, err = r.Seek(6, io.SeekStart)
	err = binary.Read(r, binary.BigEndian, &se.Client)
	return
}

// Updates returns a channel with PulseAudio updates.
func (c *Client) Updates() (updates <-chan SubscriptionEvent, err error) {
	const subscriptionMaskAll = 0x02ff
	_, err = c.request(commandSubscribe, uint32Tag, uint32(subscriptionMaskAll))
	if err != nil {
		return nil, err
	}
	return c.updates, nil
}
