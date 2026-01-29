package marketing

import (
	"github.com/go-pg/pg/v10"
)

type Service interface {
	GetAllBlogPosts(page, limit int) ([]BlogPost, error)
	GetBlogPostByID(id string) (*BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]BlogPost, error)
	GetAllHomeContent(page, limit int) ([]HomeContent, error)
	GetHomeContentByID(id string) (*HomeContent, error)
	GetAllGrooveJrContent(page, limit int) ([]GrooveJrContent, error)
	GetGrooveJrContentByID(id string) (*GrooveJrContent, error)
	GetAllAboutContent(page, limit int) ([]AboutContent, error)
	GetAboutContentByID(id string) (*AboutContent, error)
}

type service struct {
	DB PgxDB
}

func NewService(db PgxDB) Service {
	return &service{DB: db}
}

func (s *service) GetAllBlogPosts(page, limit int) ([]BlogPost, error) {
	var posts []BlogPost
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetBlogPostByID(id string) (*BlogPost, error) {
	var post BlogPost
	err := s.DB.Model(&post).
		Where("id = ?", id).
		Relation("Author").
		Relation("Tags").
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &post, err
}

func (s *service) GetBlogPostsByTag(tag string, page, limit int) ([]BlogPost, error) {
	var posts []BlogPost
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Join("JOIN blog_post_tags AS bpt ON bpt.blog_post_id = blog_post.id").
		Join("JOIN tags AS t ON t.id = bpt.tag_id").
		Where("t.name = ?", tag).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetAllHomeContent(page, limit int) ([]HomeContent, error) {
	var content []HomeContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetHomeContentByID(id string) (*HomeContent, error) {
	var content HomeContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &content, err
}

func (s *service) GetAllGrooveJrContent(page, limit int) ([]GrooveJrContent, error) {
	var content []GrooveJrContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetGrooveJrContentByID(id string) (*GrooveJrContent, error) {
	var content GrooveJrContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &content, err
}

func (s *service) GetAllAboutContent(page, limit int) ([]AboutContent, error) {
	var content []AboutContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetAboutContentByID(id string) (*AboutContent, error) {
	var content AboutContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &content, err
}
