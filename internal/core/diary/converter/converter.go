package converter

import (
	"time"

	"emwell/internal/core/diary/entites"
)

type Converter struct {
}

func (c *Converter) ConvertToPoints(rawData entites.EmotionalInfos) entites.EmotionalInfos {
	result := make(entites.EmotionalInfos, 0, len(rawData))

	count := int32(0)
	for _, elem := range rawData {
		if elem.EmotionalRate == 0 {
			continue
		}

		if len(result) == 0 {
			result = append(result, elem)
			count++
			continue
		}

		if isDatesEqual(result[len(result)-1].ReferToDate, elem.ReferToDate) {
			result[len(result)-1].EmotionalRate += elem.EmotionalRate
			count++
			continue
		}

		result[len(result)-1].EmotionalRate /= count
		count = 0

		if result[len(result)-1].EmotionalRate == 0 {
			result[len(result)-1] = elem
		} else {
			result = append(result, elem)
		}
		count++
	}

	if count > 0 {
		result[len(result)-1].EmotionalRate /= count
		count = 0
	}

	return result
}

func isDatesEqual(t, t1 time.Time) bool {
	y, m, d := t.Date()
	y1, m1, d1 := t1.Date()
	return y == y1 && m == m1 && d == d1
}
