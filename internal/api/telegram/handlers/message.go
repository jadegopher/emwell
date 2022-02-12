package handlers

type Message struct {
	typ     ResponseType
	payload MessagePayload
}

type MessagePayload struct {
	ToChatID int64
	Text     string
	Buttons  [][]Button
}

type Button struct {
	Text string
	Data string
}

func NewMessage(payload MessagePayload) *Message {
	return &Message{
		typ:     ResponseTypeMessage,
		payload: payload,
	}
}

func (m *Message) Type() ResponseType {
	return m.typ
}

func (m *Message) Payload() interface{} {
	return m.payload
}
