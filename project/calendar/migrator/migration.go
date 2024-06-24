package migrator

import "github.com/jmoiron/sqlx"

type migration struct {
	Up      func(tx *sqlx.Tx) error
	Down    func(tx *sqlx.Tx) error
	Version int
}

func newMigration(version int, upQuery, downQuery string) *migration {
	queryFn := func(query string) func(tx *sqlx.Tx) error {
		return func(tx *sqlx.Tx) error {
			_, err := tx.Exec(query)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			_ = tx.Commit()
			return err
		}
	}

	return &migration{
		Up:      queryFn(upQuery),
		Down:    queryFn(downQuery),
		Version: version,
	}
}
