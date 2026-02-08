package admin

import (
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
)

type Service interface {
	// Blog
	GetAllBlogPosts(page, limit int) ([]BlogPost, error)
	GetBlogPostByID(id string) (*BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]BlogPost, error)
	CreateBlogPost(post *BlogPost) (*BlogPost, error)
	UpdateBlogPost(post *BlogPost) (*BlogPost, error)
	DeleteBlogPost(id string) error

	// Home
	GetAllHomeContent(page, limit int) ([]HomeContent, error)
	GetHomeContentByID(id string) (*HomeContent, error)
	CreateHomeContent(content *HomeContent) (*HomeContent, error)
	UpdateHomeContent(content *HomeContent) (*HomeContent, error)
	DeleteHomeContent(id string) error

	// GrooveJr
	GetAllGrooveJrContent(page, limit int) ([]GrooveJrContent, error)
	GetGrooveJrContentByID(id string) (*GrooveJrContent, error)
	CreateGrooveJrContent(content *GrooveJrContent) (*GrooveJrContent, error)
	UpdateGrooveJrContent(content *GrooveJrContent) (*GrooveJrContent, error)
	DeleteGrooveJrContent(id string) error

	// About
	GetAllAboutContent(page, limit int) ([]AboutContent, error)
	GetAboutContentByID(id string) (*AboutContent, error)
	CreateAboutContent(content *AboutContent) (*AboutContent, error)
	UpdateAboutContent(content *AboutContent) (*AboutContent, error)
	DeleteAboutContent(id string) error
}

type service struct {
	DB interfaces.PgxDB
}

func NewService(db interfaces.PgxDB) Service {
	return &service{DB: db}
}

// Blog
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
	if err != nil {
		return nil, err
	}
	return &post, nil
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

func (s *service) CreateBlogPost(post *BlogPost) (*BlogPost, error) {
	_, err := s.DB.Model(post).Insert()
	return post, err
}

func (s *service) UpdateBlogPost(post *BlogPost) (*BlogPost, error) {
	_, err := s.DB.Model(post).Where("id = ?", post.ID).Update()
	return post, err
}

func (s *service) DeleteBlogPost(id string) error {
	_, err := s.DB.Model(&BlogPost{}).Where("id = ?", id).Delete()
	return err
}

// Home
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
	if err != nil {
		return nil, err
	}
	return &content, err
}

func (s *service) CreateHomeContent(content *HomeContent) (*HomeContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateHomeContent(content *HomeContent) (*HomeContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteHomeContent(id string) error {
	_, err := s.DB.Model(&HomeContent{}).Where("id = ?", id).Delete()
	return err
}

// GrooveJr
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
	if err != nil {
		return nil, err
	}
	return &content, err
}

func (s *service) CreateGrooveJrContent(content *GrooveJrContent) (*GrooveJrContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateGrooveJrContent(content *GrooveJrContent) (*GrooveJrContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteGrooveJrContent(id string) error {
	_, err := s.DB.Model(&GrooveJrContent{}).Where("id = ?", id).Delete()
	return err
}

// About
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
	if err != nil {
		return nil, err
	}
	return &content, err
}

func (s *service) CreateAboutContent(content *AboutContent) (*AboutContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateAboutContent(content *AboutContent) (*AboutContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteAboutContent(id string) error {
	_, err := s.DB.Model(&AboutContent{}).Where("id = ?", id).Delete()
	return err
}