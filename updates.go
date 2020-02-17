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

const (
	SubscriptionMaskNull         = 0x0000
	SubscriptionMaskSink         = 0x0001
	SubscriptionMaskSource       = 0x0002
	SubscriptionMaskSinkInput    = 0x0004
	SubscriptionMaskSourceOutput = 0x0008
	SubscriptionMaskModule       = 0x0010
	SubscriptionMaskClient       = 0x0020
	SubscriptionMaskSampleCache  = 0x0040
	SubscriptionMaskServer       = 0x0080
	SubscriptionMaskAutoload     = 0x0100
	SubscriptionMaskCard         = 0x0200
	SubscriptionMaskAll          = 0x02ff
)

type SubscriptionEvent struct {
	Type  int32
	Index int32
}

func newSubscriptionEvent(buf *bytes.Buffer) (se SubscriptionEvent, err error) {
	r := bytes.NewReader(buf.Bytes())
	_, err = r.Seek(1, io.SeekStart)
	err = binary.Read(r, binary.BigEndian, &se.Type)
	_, err = r.Seek(6, io.SeekStart)
	err = binary.Read(r, binary.BigEndian, &se.Index)
	return
}

// Updates returns a channel with PulseAudio updates.
func (c *Client) Updates(mask int32) (updates <-chan SubscriptionEvent, err error) {
	_, err = c.request(commandSubscribe, uint32Tag, mask)
	if err != nil {
		return nil, err
	}
	return c.updates, nil
}
