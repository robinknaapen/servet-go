package service

//go:generate enumer -type State -json -sql -linecomment

// State enum
type State uint8

// Available States
const (
	StateDead    State = iota // dead
	StateExited               // exited
	StateWaiting              // waiting
	StateRunning              // running
	StateFailed               // failed
)

// Server representatation of a service
type Service struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State State  `json:"state"`

	ObjectPath string `json:"object_path"`
}

// Subscription contains a service id
type Subscription struct {
	ID        int `json:"id"`
	ServiceID int `json:"service_id"`
}
