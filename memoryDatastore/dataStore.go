package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx  context.Context
	db   *sql.DB
	ID   int
	Name string
	Mail string
)

func Start() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.01:3308)/testapp")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	/* check db connection */
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection to db successful")
	}

	/* select * query */
	rows, err := db.Query("select * from Contacts where ID = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &Name, &Mail)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ID, Name, Mail)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	/* insert query*/

	res, err := db.Exec("INSERT INTO Contacts(ID,Name,Mail) VALUES(?,?,?)", 2, "John", "john@mail.com")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
