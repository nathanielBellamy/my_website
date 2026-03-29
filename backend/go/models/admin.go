package models

import "time"

type AuthorDTO struct {
	Name string `json:"name"`
}

type TagDTO struct {
	Name string `json:"name"`
}

type CreateBlogPostDTO struct {
	Title         string     `json:"title"`
	Order         int        `json:"order"`
	Content       string     `json:"content"`
	Author        *AuthorDTO `json:"author"`
	Tags          []*TagDTO  `json:"tags"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
}
