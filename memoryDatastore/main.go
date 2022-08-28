package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	Id        string `json:"id"`
	User_name string `json:"user_name"`
	Mail      string `json:"mail"`
}

type server struct {
	db *sql.DB
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.01:3306)/aiggato")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/* check db connection */
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection to db successful")
	}

	s := server{
		db: db,
	}
	r := router(s)

	log.Fatal(http.ListenAndServe(":8081", r))

}
