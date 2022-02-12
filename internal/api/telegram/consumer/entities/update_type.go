package entities

type UpdateType int8

const (
	UpdateTypeUndefined UpdateType = iota
	UpdateTypeMessage
	UpdateTypeCallback
)

func (u UpdateType) String() string {
	str := []string{
		"undefined",
		"message",
		"callback",
	}
	if int(u) < len(str) {
		return ""
	}
	return str[u]
}
