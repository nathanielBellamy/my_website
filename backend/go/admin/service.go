package admin

import (
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/models"
)

type Service interface {
	// Blog
	GetAllBlogPosts(page, limit int) ([]models.BlogPost, error)
	GetBlogPostByID(id string) (*models.BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error)
	CreateBlogPost(post *models.BlogPost) (*models.BlogPost, error)
	UpdateBlogPost(post *models.BlogPost) (*models.BlogPost, error)
	DeleteBlogPost(id string) error

	// Home
	GetAllHomeContent(page, limit int) ([]models.HomeContent, error)
	GetHomeContentByID(id string) (*models.HomeContent, error)
	CreateHomeContent(content *models.HomeContent) (*models.HomeContent, error)
	UpdateHomeContent(content *models.HomeContent) (*models.HomeContent, error)
	DeleteHomeContent(id string) error

	// GrooveJr
	GetAllGrooveJrContent(page, limit int) ([]models.GrooveJrContent, error)
	GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error)
	CreateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	UpdateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	DeleteGrooveJrContent(id string) error

	// About
	GetAllAboutContent(page, limit int) ([]models.AboutContent, error)
	GetAboutContentByID(id string) (*models.AboutContent, error)
	CreateAboutContent(content *models.AboutContent) (*models.AboutContent, error)
	UpdateAboutContent(content *models.AboutContent) (*models.AboutContent, error)
	DeleteAboutContent(id string) error
}

type service struct {
	DB interfaces.PgxDB
}

func NewService(db interfaces.PgxDB) Service {
	return &service{DB: db}
}

// Blog
func (s *service) GetAllBlogPosts(page, limit int) ([]models.BlogPost, error) {
	var posts []models.BlogPost
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetBlogPostByID(id string) (*models.BlogPost, error) {
	var post models.BlogPost
	err := s.DB.Model(&post).
		Relation("Author").
		Relation("Tags").
		Where("blog_post.id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *service) GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error) {
	var posts []models.BlogPost
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

func (s *service) CreateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	_, err := s.DB.Model(post).Insert()
	return post, err
}

func (s *service) UpdateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	// 1. Update the BlogPost itself (title, content, author_id).
	_, err := s.DB.Model(post).
		Column("title", "content", "author_id").
		Where("id = ?", post.ID).
		Update()
	if err != nil {
		return nil, err
	}

	// 2. Delete existing tag associations.
	_, err = s.DB.Model((*models.BlogPostTag)(nil)).
		Where("blog_post_id = ?", post.ID).
		Delete()
	if err != nil {
		return nil, err
	}

	// 3. Create new tag associations.
	if len(post.Tags) > 0 {
		var blogPostTags []models.BlogPostTag
		for _, tag := range post.Tags {
			blogPostTags = append(blogPostTags, models.BlogPostTag{
				BlogPostID: post.ID,
				TagID:      tag.ID,
			})
		}
		_, err = s.DB.Model(&blogPostTags).Insert()
		if err != nil {
			return nil, err
		}
	}

	return post, nil
}

func (s *service) DeleteBlogPost(id string) error {
	_, err := s.DB.Model(&models.BlogPost{}).Where("id = ?", id).Delete()
	return err
}

// Home
func (s *service) GetAllHomeContent(page, limit int) ([]models.HomeContent, error) {
	var content []models.HomeContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetHomeContentByID(id string) (*models.HomeContent, error) {
	var content models.HomeContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *service) CreateHomeContent(content *models.HomeContent) (*models.HomeContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateHomeContent(content *models.HomeContent) (*models.HomeContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteHomeContent(id string) error {
	_, err := s.DB.Model(&models.HomeContent{}).Where("id = ?", id).Delete()
	return err
}

// GrooveJr
func (s *service) GetAllGrooveJrContent(page, limit int) ([]models.GrooveJrContent, error) {
	var content []models.GrooveJrContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error) {
	var content models.GrooveJrContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *service) CreateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteGrooveJrContent(id string) error {
	_, err := s.DB.Model(&models.GrooveJrContent{}).Where("id = ?", id).Delete()
	return err
}

// About
func (s *service) GetAllAboutContent(page, limit int) ([]models.AboutContent, error) {
	var content []models.AboutContent
	err := s.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetAboutContentByID(id string) (*models.AboutContent, error) {
	var content models.AboutContent
	err := s.DB.Model(&content).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *service) CreateAboutContent(content *models.AboutContent) (*models.AboutContent, error) {
	_, err := s.DB.Model(content).Insert()
	return content, err
}

func (s *service) UpdateAboutContent(content *models.AboutContent) (*models.AboutContent, error) {
	_, err := s.DB.Model(content).Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteAboutContent(id string) error {
	_, err := s.DB.Model(&models.AboutContent{}).Where("id = ?", id).Delete()
	return err
}
