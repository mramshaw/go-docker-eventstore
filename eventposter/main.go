package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type event struct {
	EventId   string `json:"eventId"`
	EventType string `json:"eventType"`
	Data      string `json:"data"`
}

type events []*event

const (
	someEventType = "some-event-type"
	numberEvents  = 10000
)

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {

	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func main() {

	count := 1
	for count < numberEvents {

		data := fmt.Sprintf("{ value : %v }", count)
		uuid, _ := newUUID()
		event := event{
			EventId:   uuid,
			EventType: someEventType,
			Data:      data,
		}

		events := events{&event}
		body, _ := json.Marshal(events)
		reader := bytes.NewReader(body)

		resp, err := http.Post("http://eventstore:2113/streams/newstream", "application/vnd.eventstore.events+json", reader)
		if err != nil {
			log.Println("Error sending event", err.Error())
		}
		log.Println(resp.Status)

		count++
	}
	log.Printf("Have now created %d events\n", numberEvents)
}
