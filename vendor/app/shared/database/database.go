package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

type database struct {
	base   context.Context
	cancel func()
	db     *sql.DB
}

func Run() {

	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang")
}
