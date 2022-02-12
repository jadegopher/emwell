package entities

type Callback struct {
	ID              string
	From            Sender
	Message         Message
	InlineMessageID string
	Data            string
}
