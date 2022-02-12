package handlers

type Callback struct {
	typ     ResponseType
	payload CallbackPayload
}

type CallbackPayload struct {
	CallbackID string
	Text       string
}

func NewCallback(payload CallbackPayload) *Callback {
	return &Callback{
		typ:     ResponseTypeCallback,
		payload: payload,
	}
}

func (m *Callback) Type() ResponseType {
	return m.typ
}

func (m *Callback) Payload() interface{} {
	return m.payload
}
