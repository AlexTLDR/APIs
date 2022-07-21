package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	db   *sql.DB
// 	ID   int
// 	Name string
// 	Mail string
// )

func Start() {
	ctx := context.Background()
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
	rows, err := db.Query("select * from Contacts where ID = ?", 2)
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

	err = db.QueryRowContext(ctx, "SELECT Name, Mail from Contacts WHERE ID = ?", 1).Scan(&Name, &Mail)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", ID)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("name is %q and mail is %q\n", Name, Mail)
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

	/* update query*/

	res, err = db.Exec("UPDATE Contacts set Name =?,Mail =? where ID =?", "Johnny", "john@mail.com", 2)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	/* delete query */
	_, err = db.Exec("delete from Contacts where ID = ?", 5)
	if err != nil {
		panic(err)
	}
}
