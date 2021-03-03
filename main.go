package main

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

var posts = []*Post{
	{Text: "first post", Created: 1},
	{Text: "second post", Created: 2},
}

type Post struct {
	Text    string `json:"text"`
	Created int    `json:"created"`
}

type PostResponse struct {
	*Post
}

type PostRequest struct {
	*Post
}

func (rd *PostResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func newPostResponse(post *Post) *PostResponse {
	return &PostResponse{Post: post}
}

func newPostListResponse(posts []*Post) []render.Renderer {
	var list []render.Renderer
	for _, post := range posts {
		list = append(list, newPostResponse(post))
	}
	return list
}

func (a *PostRequest) Bind(_ *http.Request) error {
	if a.Post == nil {
		return errors.New("missing required Post fields")
	}
	return nil
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Welcome to the cms from 3crabs!"))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	_ = render.RenderList(w, r, newPostListResponse(posts))
}

func addPosts(w http.ResponseWriter, r *http.Request) {
	data := &PostRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	post := data.Post
	posts = append(posts, post)

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, newPostResponse(post))
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
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
