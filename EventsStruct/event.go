package eventsStruct

type Eventer interface {
	ProcessEvent(e interface{})
}

var subscribers []Eventer

func AddSubscribers(subscriber Eventer) {
	subscribers = append(subscribers, subscriber)
}

func PushEvent(e interface{}) {
	for _, obj := range subscribers {
		go (obj).ProcessEvent(e)
	}
}
