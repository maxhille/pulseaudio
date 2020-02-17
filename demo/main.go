package main

import (
	"log"

	"github.com/mafik/pulseaudio"
)

func main() {
	c, err := pulseaudio.NewClient()
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}

	uc, err := c.Updates(pulseaudio.SubscriptionMaskAll)
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		ev := <-uc
		if ev.Type != pulseaudio.SubscriptionEventChange|
			pulseaudio.SubscriptionEventSink {
			continue
		}
		log.Printf("Event 'change' on sink #%d", ev.Index)
		vol, _ := c.Volume()
		log.Printf("Volume is %d%%", int(vol*100))
	}
}
