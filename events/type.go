package events

type Fetcher interface {
	Fetch() ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	UnknowType = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
