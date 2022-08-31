package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	Id        string `json:"id"`
	User_name string `json:"user_name"`
	Mail      string `json:"mail"`
}

const (
	projectID  = "aiggato-upload"
	bucketName = "aiggato-files"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

type server struct {
	db *sql.DB
}

func main() {
	/*sql start*/
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

	//log.Fatal(http.ListenAndServe(":8081", r))
	go http.ListenAndServe(":8081", r)
	/*sql end*/

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/alex/Keys/GoogleCloud/aiggato/aiggato-upload-18942db9665f.json") // the path to the connection json

	client, err := storage.NewClient(context.Background())
	fmt.Println(client)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "upload-files/", // the path of the files in the bucket
	}

	g := gin.Default() // create a function
	g.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile("file-input") //the name of the form
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		blobFile, err := f.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = uploader.UploadFile(blobFile, f.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	g.Run(":8082") // listen and serve "localhost:8080"

}
