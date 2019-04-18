package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	useCache *string
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

type UserCache map[uint]User

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

func mergePostUser(posts *[]Post, ptrCache *UserCache) {
	for _, post := range *posts {
		user := getUser(ptrCache, post.UserID)
		fmt.Printf("Post: %v - User: %v\n", post.ID, user.Name)
	}
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
	return uint(len(getPosts()))
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

func getUser(cache *UserCache, id uint) User {
	if cache == nil {
		return json2user(getArtifact(USER_API_ENDPOINT, &id))
	}
	if user, ok := (*cache)[id]; ok {
		//log.Printf("User cache HIT: %v", id)
		return user
	} else {
		//log.Printf("User cache MISS: %v", id)
		// Appel à l'API
		// Mise à jour du cache
		(*cache)[id] = json2user(getArtifact(USER_API_ENDPOINT, &id))
		return (*cache)[id]
	}
}

// On gère les options au lancement du programme
func init() {
	useCache = flag.String("cache", "true", "Use User cache")
	flag.Parse()
}

func runTimedFunc(f func(posts *[]Post, ptrCache *UserCache), posts *[]Post, ptrCache *UserCache) time.Duration {
	start := time.Now()
	f(posts, ptrCache)
	t := time.Now()

	return t.Sub(start)
}

func main() {
	var userCache UserCache
	var ptrCache *UserCache

	if *useCache == "true" {
		ptrCache = &userCache
		userCache = make(UserCache)
	} else {
		ptrCache = nil
	}
	//var i uint = 1
	//fmt.Println(json2user(getArtifact(USER_API_ENDPOINT, &i)))
	//fmt.Println(json2post(getArtifact(POST_API_ENDPOINT, &i)))

	// On récupère le nombre de posts pour créer un tableau de taille fixe
	//var postsNumber uint = getPostsNumber()

	// On récupère les posts
	posts := getPosts()

	// Pour chaque post :
	// - on récupère les informations de l'utilisateur
	// - on joint l'id du post et de l'utilisateur
	elapsed := runTimedFunc(mergePostUser, &posts, ptrCache)
	fmt.Println("Duration:", elapsed)
}
