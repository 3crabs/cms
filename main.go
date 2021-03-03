package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Welcome to the cms from 3crabs!"))
}

func GetPosts(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))
	r.Get("/", Index)
	r.Get("/posts", GetPosts)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
