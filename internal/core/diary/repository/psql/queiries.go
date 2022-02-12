package psql

const (
	InsertQuery = `
		INSERT INTO emotional_diary (user_id, emotion_rate)
		VALUES ($1, $2) RETURNING id;`
)
