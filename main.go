package main

import (
	"container/list"

	"github.com/gofrs/uuid"
)

// Event type
type Event struct {
	uuid        uuid.UUID
	connections list.List
}

func uuidTest() {
	event := Event{}
	event.uuid, _ = uuid.NewV4()

	println(event.uuid.String())
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

func main() {
	uuidTest()
	listTest()
}
