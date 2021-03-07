package models

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type Post struct {
	Text    string `json:"text"`
	Created int    `json:"created"`
}

type PostResponse struct {
	*Post
}

func (rd *PostResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type PostRequest struct {
	*Post
}

func (a *PostRequest) Bind(_ *http.Request) error {
	if a.Post == nil {
		return errors.New("missing required Post fields")
	}
	return nil
}

func NewPostResponse(post *Post) *PostResponse {
	return &PostResponse{Post: post}
}

func NewPostListResponse(posts []*Post) []render.Renderer {
	var list []render.Renderer
	for _, post := range posts {
		list = append(list, NewPostResponse(post))
	}
	return list
}
