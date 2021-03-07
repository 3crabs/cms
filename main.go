package main

import (
	"cms/db"
	"cms/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Welcome to the cms from 3crabs!"))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	_ = render.RenderList(w, r, models.NewPostListResponse(db.GetPosts()))
}

func addPosts(w http.ResponseWriter, r *http.Request) {
	data := &models.PostRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	post := data.Post
	db.AddPost(post)

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, models.NewPostResponse(post))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

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
