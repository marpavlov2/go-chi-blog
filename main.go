package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Likes    uint      `json:"likes"`
	Comments []Comment `json:"comment"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	posts     = make(map[int]Post)
	postID    = 1
	commentID = 1
)

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	postList := []Post{}
	for _, post := range posts {
		postList = append(postList, post)
	}

	// w.Write([]byte("Bog i Kiki"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(postList); error != nil {
		log.Fatalf("Can't parse JSON %v", error)
	}

}
func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post.ID = postID
	posts[post.ID] = post
	postID++

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse json %v", err)
	}
}
func main() {

	fmt.Println("Project started")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/posts", getAllPosts)
	r.Post("/posts", createPost)

	// 	post := Post{ID: 1, Title: "Touch Typing", Content: "How about NO?", Likes: 100, Comments: nil}
	// append(posts, post)
	if error := http.ListenAndServe(":8080", r); error != nil {
		log.Fatalf("Cant listen and serve %v", error)
	}
}
