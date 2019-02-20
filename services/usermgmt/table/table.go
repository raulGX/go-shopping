package table

import (
	"database/sql"

	"github.com/pkg/errors"
)

func EnsureUsersTableExists(db *sql.DB) error {
	const qry = `
CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY,
	username VARCHAR (500) UNIQUE NOT NULL,
	password VARCHAR (500) NOT NULL,
	created_at timestamp with time zone DEFAULT current_timestamp,
	last_login TIMESTAMP
)`

	if _, err := db.Exec(qry); err != nil {
		err = errors.Wrapf(err,
			"Users table creation query failed (%s)",
			qry)
		return err
	}

	return nil
}
