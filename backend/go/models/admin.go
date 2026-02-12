package models

type AuthorDTO struct {
	Name string `json:"name"`
}

type TagDTO struct {
	Name string `json:"name"`
}

type CreateBlogPostDTO struct {
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Author  *AuthorDTO  `json:"author"`
	Tags    []*TagDTO   `json:"tags"`
}
