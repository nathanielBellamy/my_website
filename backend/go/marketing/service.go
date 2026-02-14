package marketing

import (
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/models"
)

type Service interface {
	GetAllBlogPosts(page, limit int) ([]models.BlogPost, error)
	GetBlogPostByID(id string) (*models.BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error)
	GetAllHomeContent(page, limit int) ([]models.HomeContent, error)
	GetHomeContentByID(id string) (*models.HomeContent, error)
	GetAllGrooveJrContent(page, limit int) ([]models.GrooveJrContent, error)
	GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error)
	GetAllAboutContent(page, limit int) ([]models.AboutContent, error)
	GetAboutContentByID(id string) (*models.AboutContent, error)
}

type service struct {
	DB interfaces.PgxDB
}

func NewService(db interfaces.PgxDB) Service {
	return &service{DB: db}
}

func (s *service) GetAllBlogPosts(page, limit int) ([]models.BlogPost, error) {
	posts := make([]models.BlogPost, 0)
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at < NOW() AND (blog_post.deactivated_at IS NULL OR blog_post.deactivated_at > NOW())").
		Order("ordering ASC", "activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetBlogPostByID(id string) (*models.BlogPost, error) {
	var post models.BlogPost
	err := s.DB.Model(&post).
		Where("blog_post.id = ?", id).
		Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at < NOW() AND (blog_post.deactivated_at IS NULL OR blog_post.deactivated_at > NOW())").
		Relation("Author").
		Relation("Tags").
		Select()
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *service) GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error) {
	posts := make([]models.BlogPost, 0)
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Join("JOIN blog_post_tags AS bpt ON bpt.blog_post_id = blog_post.id").
		Join("JOIN tags AS t ON t.id = bpt.tag_id").
		Where("t.name = ?", tag).
		Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at < NOW() AND (blog_post.deactivated_at IS NULL OR blog_post.deactivated_at > NOW())").
		Order("ordering ASC", "activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetAllHomeContent(page, limit int) ([]models.HomeContent, error) {
	content := make([]models.HomeContent, 0)
	err := s.DB.Model(&content).
		Where("home_content.activated_at IS NOT NULL AND home_content.activated_at < NOW() AND (home_content.deactivated_at IS NULL OR home_content.deactivated_at > NOW())").
		Order("ordering ASC", "activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetHomeContentByID(id string) (*models.HomeContent, error) {
	var content models.HomeContent
	err := s.DB.Model(&content).
		Where("home_content.id = ?", id).
		Where("home_content.activated_at IS NOT NULL AND home_content.activated_at < NOW() AND (home_content.deactivated_at IS NULL OR home_content.deactivated_at > NOW())").
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *service) GetAllGrooveJrContent(page, limit int) ([]models.GrooveJrContent, error) {
	content := make([]models.GrooveJrContent, 0)
	err := s.DB.Model(&content).
		Where("groove_jr_content.activated_at IS NOT NULL AND groove_jr_content.activated_at < NOW() AND (groove_jr_content.deactivated_at IS NULL OR groove_jr_content.deactivated_at > NOW())").
		Order("ordering ASC", "activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error) {
	var content models.GrooveJrContent
	err := s.DB.Model(&content).
		Where("groove_jr_content.id = ?", id).
		Where("groove_jr_content.activated_at IS NOT NULL AND groove_jr_content.activated_at < NOW() AND (groove_jr_content.deactivated_at IS NULL OR groove_jr_content.deactivated_at > NOW())").
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *service) GetAllAboutContent(page, limit int) ([]models.AboutContent, error) {
	content := make([]models.AboutContent, 0)
	err := s.DB.Model(&content).
		Where("about_content.activated_at IS NOT NULL AND about_content.activated_at < NOW() AND (about_content.deactivated_at IS NULL OR about_content.deactivated_at > NOW())").
		Order("ordering ASC", "activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return content, err
}

func (s *service) GetAboutContentByID(id string) (*models.AboutContent, error) {
	var content models.AboutContent
	err := s.DB.Model(&content).
		Where("about_content.id = ?", id).
		Where("about_content.activated_at IS NOT NULL AND about_content.activated_at < NOW() AND (about_content.deactivated_at IS NULL OR about_content.deactivated_at > NOW())").
		Select()
	if err != nil {
		return nil, err
	}
	return &content, nil
}
