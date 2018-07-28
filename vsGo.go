package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Blog struct {
	Id          int    `json:"Id,omitempty"`
	Description string `json:"Description,omitempty"`
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	var blog []Blog
	db, err := sql.Open("mysql", "userConnectionString")
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	rows, err := db.Query("SELECT Id,Description FROM blog")
	if err != nil {
		fmt.Println("New Error!")
		fmt.Println(err.Error())
	}
	var rId int
	var rDescription string
	for rows.Next() {
		err = rows.Scan(&rId, &rDescription)
		if err != nil {
			log.Fatal(err)
		}
		blog = append(blog, Blog{rId, rDescription})
	}
	json.NewEncoder(w).Encode(blog)
	defer db.Close()
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/blog", GetBlog).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
