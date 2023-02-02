package database

import (
	"github.com/jmoiron/sqlx"
)

var (
	// DB is the database connection
	DB *sqlx.DB
)
