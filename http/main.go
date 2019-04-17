package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"m2i.fr/apiclient/data"

	"github.com/gorilla/mux"
)

var (
	uri string
	err error

	db   *sql.DB
	stmt *sql.Stmt
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func postGetDB(id *uint) []data.Post {
	var query string
	var posts []data.Post

	if id == nil {
		query = fmt.Sprintf("SELECT id,title,body FROM post;")
	} else {
		query = fmt.Sprintf("SELECT id,title,body FROM post WHERE id=%v LIMIT 1;", *id)
	}

	log.Println(query)
	rows, err := db.Query(query)
	post := data.Post{}

	checkError(err)

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Body)
		checkError(err)

		posts = append(posts, post)
	}

	return posts
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	buf, _ := strconv.Atoi(mux.Vars(r)["id"])
	id := uint(buf)

	posts := postGetDB(&id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	posts := postGetDB(nil)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func init() {
	uri = "root:12345@tcp(localhost:3306)/form_go"
	db, err = sql.Open("mysql", uri)
	checkError(err)
}

func main() {
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/posts/", postsHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", postHandler).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
