package entites

import (
	"bytes"
	"errors"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var (
	ErrInvalidUserID        = errors.New("invalid userID")
	ErrInvalidEmotionalRate = errors.New("invalid emotional rate")
)

const (
	MinEmotionalRate     = -1000
	WorstEmotionalRate   = -750
	WorseEmotionalRate   = -500
	BadEmotionalRate     = -250
	NeutralEmotionalRate = 0
	GoodEmotionalRate    = 250
	BetterEmotionalRate  = 500
	BestEmotionalRate    = 750
	MaxEmotionalRate     = 1000
)

type EmotionalInfo struct {
	id            int64
	UserID        int64
	EmotionalRate int32
	ReferToDate   time.Time
	createdAt     time.Time
}

func NewEmotionalDiaryEntity(id int64, createdAt time.Time, entity EmotionalInfo) EmotionalInfo {
	return EmotionalInfo{
		id:            id,
		UserID:        entity.UserID,
		EmotionalRate: entity.EmotionalRate,
		ReferToDate:   entity.ReferToDate,
		createdAt:     createdAt,
	}
}

func (e *EmotionalInfo) ID() int64 {
	return e.id
}

func (e *EmotionalInfo) CreatedAt() time.Time {
	return e.createdAt
}

func (e *EmotionalInfo) Validate() error {
	if e.UserID <= 0 {
		return ErrInvalidUserID
	}

	if e.EmotionalRate < MinEmotionalRate || e.EmotionalRate > MaxEmotionalRate {
		return ErrInvalidEmotionalRate
	}

	return nil
}

type EmotionalInfos []EmotionalInfo

const dateFormat = "2006-01-02"

func (e EmotionalInfos) Visualize() ([]byte, error) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "",
			Start:      0,
			End:        0,
			Throttle:   0,
			XAxisIndex: nil,
			YAxisIndex: nil,
		}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}))
	line.SetXAxis(getAxisMarks(e)).
		AddSeries("Emotional wellness", dataToLineData(e)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	b := bytes.NewBuffer([]byte{})
	if err := line.Render(b); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func getAxisMarks(data []EmotionalInfo) []string {
	result := make([]string, 0, len(data))
	for _, elem := range data {
		result = append(result, elem.ReferToDate.Format(dateFormat))
	}
	return result
}

func dataToLineData(data []EmotionalInfo) []opts.LineData {
	result := make([]opts.LineData, 0)
	for _, elem := range data {
		result = append(result, opts.LineData{Value: elem.EmotionalRate})
	}
	return result
}
