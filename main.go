package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Tasks struct {
	row string
}

var logger = log.New(os.Stdout, "Shreya: ", log.Ldate+log.Ltime+log.Lshortfile)
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:nithinkumar1996@tcp(localhost:3306)/Storage")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	logger.Print("mysql connected/n")
}

func main() {
	http.HandleFunc("/api/create", insertdata)
	http.HandleFunc("/api/read", getdata)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	logger.Println("http://127.0.0.1:8050")
	http.ListenAndServe(":8050", nil)
	defer db.Close()

}
