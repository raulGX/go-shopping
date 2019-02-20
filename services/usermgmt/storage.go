package usermgmt

import (
	sql "database/sql"
	"fmt"

	table "github.com/raulGX/go-shopping/services/usermgmt/table"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "makethisanenvvariable"
	dbname   = "usermgmt"
)

func NewPostgresConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = table.EnsureUsersTableExists(db)
	if err != nil {
		panic(err)
	}

	return db
}
