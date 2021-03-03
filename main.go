package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Welcome to the cms from 3crabs!"))
}

func getPosts(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("get posts"))
}

func addPosts(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("add posts"))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))

	r.Get("/", index)

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", getPosts)
		r.Post("/", addPosts)
	})

	log.Println("Run server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
