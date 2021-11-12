package database

import (
	"bookstoreapi/app/queries"
)

type Queries struct {
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {
	db,err:=PostgreSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		BookQueries:&queries.BookQueries{DB:db},
	},nil
}

