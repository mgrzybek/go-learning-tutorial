package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// go routine pour récupérer les info de https://jsonplaceholder.typicode.com/photos

// channel pour

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

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getURL(url string) []byte {
	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	checkError(err)

	return bs
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

func mergePostUser() {

}

// Si len == 0, alors on fait un tableau dynamique
func getPosts() []Post {
	var posts []Post

	data := getArtifact(POST_API_ENDPOINT, nil)
	err := json.Unmarshal(data, &posts)

	checkError(err)
	return posts
}

func getPostsNumber() uint {
	return uint(len(getPosts(0)))
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

func getUser(id uint) User {
	var cache map[uint]User

	if user, ok := cache[id]; ok {
		return user
	} else {
		return func() {
			// Appel à l'API
			user := json2user(getArtifact(USER_API_ENDPOINT, id))
			// Mise à jour du cache
			cache[id] = user

			return user
		}(id)
	}
}

func main() {
	//var i uint = 1
	//fmt.Println(json2user(getArtifact(USER_API_ENDPOINT, &i)))
	//fmt.Println(json2post(getArtifact(POST_API_ENDPOINT, &i)))

	// On récupère le nombre de posts pour créer un tableau de taille fixe
	var postsNumber uint = getPostsNumber()

	// On récupère les posts
	posts := getPosts()
	// Pour chaque post :
	// - on récupère les informations de l'utilisateur
	// - on joint l'id du post et de l'utilisateur

	for post := range posts {
		user := json2user(getArtifact(USER_API_ENDPOINT, post.UserID))
	}
}
