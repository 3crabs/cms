package db

import (
	"cms/models"
)

var posts []*models.Post

func GetPosts() []*models.Post {
	return posts
}

func AddPost(p *models.Post) {
	_ = append(posts, p)
}
