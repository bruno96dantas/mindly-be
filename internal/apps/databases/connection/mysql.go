package connection

import (
	"database/sql"

	"github.com/bruno96dantas/mindly-be/internal/apps/settings"
)

func MySQL() (*sql.DB, error) {
	var db *sql.DB
	var err error

	settings.LoadMySQLConfigs()
	db, err = sql.Open("mysql", settings.MySQLConfig.FormatDSN())

	return db, err
}
