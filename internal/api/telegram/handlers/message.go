package handlers

type Message struct {
	payload MessagePayload
}

type MessagePayload struct {
	ToChatID       int64
	Text           string
	InlineKeyboard [][]Button
}

type Button struct {
	Text string
	Data string
}

func NewMessage(payload MessagePayload) *Message {
	return &Message{
		payload: payload,
	}
}

func (m *Message) Type() ResponseType {
	return ResponseTypeMessage
}

func (m *Message) Payload() interface{} {
	return m.payload
}
