package main

import (
	"container/list"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
)

// Event type
type Event struct {
	UUID  string `json:"uuid"`
	Event string `json:"event"`

	connections list.List
}

func uuidTest() {
	event := Event{}

	uuid, _ := uuid.NewV4()
	event.UUID = uuid.String()

	println(event.UUID)
}

func listTest() {
	ev := Event{}
	ev2 := Event{}

	ev2.connections.PushBack(40)
	ev2.connections.PushBack(50)

	ev.connections.PushBack(&ev2)

	ev2.connections.PushBack(60)
	ev2.connections.PushBack(70)

	ee := ev.connections.Front().Value.(*Event)

	for e := ee.connections.Front(); e != nil; e = e.Next() {
		println(e.Value.(int))
	}
}

func jsonTest() {
	jsonText := `[
		{
			"uuid": "uuid1",
			"event": "SAY"
		},
		{
			"uuid": "uuid2",
			"event": "MOVE"
		}
	]`

	var events []Event

	json.Unmarshal([]byte(jsonText), &events)

	fmt.Println(events)
}

func main() {
	uuidTest()
	listTest()
	jsonTest()
}
