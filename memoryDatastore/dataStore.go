package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	db, err := sql.Open("mysql", "root|QPi+fxg74Fb_G0R4KW2KY+X&?#4a;#53@/tcp(localhost:5555)")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
