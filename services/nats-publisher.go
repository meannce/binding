package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 1; i <= 5; i++ {
		msg := []byte("Hello NATS! Message #" + string(rune(i)))
		err = nc.Publish("updates", msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Sent: %s\n", msg)
		time.Sleep(1 * time.Second)
	}
}
