package psql

const (
	InsertQuery = `
		INSERT INTO emotional_diary (user_id, emotion_rate)
		VALUES ($1, $2) RETURNING id;`

	SelectByUserIDQuery = `
		SELECT id, user_id, emotion_rate, refers_to_date, created_at
		FROM emotional_diary
		WHERE user_id = $1
		  AND (refers_to_date >= $2 AND refers_to_date <= $3)
		ORDER BY refers_to_date;`
)
