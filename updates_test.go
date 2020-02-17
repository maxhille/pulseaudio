package pulseaudio

import (
	"bytes"
	"testing"
)

func TestNewSubscriptionEvent(t *testing.T) {
	buf := bytes.NewBuffer([]byte{
		0x4c,
		0x00, 0x00, 0x00, 0x25,
		0x4c,
		0x00, 0x00, 0x00, 0x62,
	})
	se, err := newSubscriptionEvent(buf)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if se.Type != SubscriptionEventRemove|SubscriptionEventClient {
		t.Errorf("expected SubscriptionEventChange, but got %x",
			se.Type)
	}

	if se.Index != 98 {
		t.Errorf("expected #98, but got %d", se.Index)
	}
}
