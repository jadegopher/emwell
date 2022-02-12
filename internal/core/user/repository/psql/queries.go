package psql

const (
	InsertQuery = `
		INSERT INTO users (name, language, telegram_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (telegram_id) DO UPDATE SET name     = excluded.name,
		                                        language = excluded.language
		RETURNING id;`
)
