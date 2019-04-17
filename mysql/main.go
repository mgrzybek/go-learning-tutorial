package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_API_ENDPOINT = "https://jsonplaceholder.typicode.com/users"
	POST_API_ENDPOINT = "https://jsonplaceholder.typicode.com/posts"
)

type Post struct {
	UserID uint
	ID     uint
	Title  string
	Body   string
}

type User struct {
	ID                    uint
	Name, Username, Email string
}

type MultiTaskSQL struct {
	stmt sql.Stmt
	Lock sync.Mutex
	WG   sync.WaitGroup
}

func getURL(url string) []byte {
	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	checkError(err)

	return bs
}

func json2user(input []byte) User {
	var result User
	err := json.Unmarshal(input, &result)
	checkError(err)
	return result
}

func json2post(input []byte) Post {
	var result Post
	err := json.Unmarshal(input, &result)
	checkError(err)
	return result
}

// Si len == 0, alors on fait un tableau dynamique
func getPosts() []Post {
	var posts []Post

	data := getArtifact(POST_API_ENDPOINT, nil)
	err := json.Unmarshal(data, &posts)

	checkError(err)
	return posts
}

/*
 * Pour chaque post (dans le main), on crée une go routine pour récupérer le nom de la personne
 */
func getArtifact(url string, id *uint) []byte {
	if id == nil {
		return getURL(fmt.Sprintf("%v", url))
	}
	return getURL(fmt.Sprintf("%v/%v", url, *id))
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func insertPost(stmt *sql.Stmt, wg *sync.WaitGroup, mutex *sync.Mutex, errors chan error, post *Post) {
	mutex.Lock()
	_, err := stmt.Exec(post.Title, post.Body)
	if err != nil {
		errors <- err
	}
	mutex.Unlock()
	wg.Done()
}

func printErrors(errors chan error) {
	err := <-errors
	fmt.Println(err)
}

func main() {
	var semaphores MultiTaskSQL

	uri := "root:12345@tcp(localhost:3306)/form_go"
	db, err := sql.Open("mysql", uri)
	checkError(err)
	defer db.Close()

	semaphores.stmt, err = db.Prepare("INSERT INTO post (title, body) VALUES (?,?);")
	checkError(err)
	defer semaphores.stmt.Close()

	// On crée un channel pour les erreurs d'insertion
	errors := make(chan error, 10)

	// On récupère les posts
	// On insère tout en base
	posts := getPosts()

	fmt.Printf("On a %v posts à insérer", len(posts))
	semaphores.WG.Add(len(posts))
	for _, post := range posts {
		go insertPost(&semaphores, errors, &post)
	}
	go printErrors(errors)

	waitgroup.Wait()
}
