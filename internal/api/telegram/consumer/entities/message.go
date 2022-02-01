package entities

type Message struct {
	ID   int64
	From Sender
	Chat Chat
	Text string
}
