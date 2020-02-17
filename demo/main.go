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
		log.Printf("Event: %x on client #%d", ev.Type, ev.Client)
	}
}
