package timer

import "time"

type Timer struct {
}

func (t *Timer) Now() time.Time {
	return time.Now().UTC()
}
