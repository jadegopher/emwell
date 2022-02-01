package entities

type UpdateType int8

const (
	UpdateTypeUndefined UpdateType = iota
	UpdateTypeMessage
)

func (u UpdateType) String() string {
	str := []string{
		"undefined",
		"message",
	}
	if int(u) < len(str) {
		return ""
	}
	return str[u]
}
