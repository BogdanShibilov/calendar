package migrator

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"io/fs"
)

var (
	baseFs     fs.FS
	db         *sqlx.DB
	migrations []migration
)

func SetBaseFs(fsys fs.FS) {
	baseFs = fsys
}

func SetDB(database *sqlx.DB) {
	db = database
}

func Up(dir string) error {
	err := parseAllMigrations(dir)
	if err != nil {
		return err
	}

	for _, m := range migrations {
		tx, _ := db.BeginTxx(context.TODO(), nil)
		_ = m.Up(tx)
	}
	return nil
}
