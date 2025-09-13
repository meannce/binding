package main

import (
	"log"
	"github.com/nats-io/nats.go"
)

func main() {
	// connect to nas server
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// subscribe to a subject
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening to subject 'updates'...")

	for {
		msg, err := sub.NextMsg(0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received: %s\n", string(msg.Data))
	}
}
