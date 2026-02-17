package marketing

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/models"
)

type Service interface {
	GetAllBlogPosts(page, limit int, tags []string) ([]models.BlogPost, error)
	GetBlogPostByID(id string) (*models.BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error)
	GetTags(search string, limit int) ([]models.TagWithUsage, error)
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

func (s *service) GetAllBlogPosts(page, limit int, tags []string) ([]models.BlogPost, error) {
	posts := make([]models.BlogPost, 0)
	query := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at < NOW() AND (blog_post.deactivated_at IS NULL OR blog_post.deactivated_at > NOW())")

	if len(tags) > 0 {
		// We use a subquery to find blog post IDs that match all tags.
		query.Where("blog_post.id IN (SELECT blog_post_id FROM blog_post_tags WHERE tag_id IN (?) GROUP BY blog_post_id HAVING count(distinct tag_id) = ?)", pg.In(tags), len(tags))
	}

	err := query.Order("blog_post.ordering ASC", "blog_post.activated_at DESC").
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
		Order("blog_post.ordering ASC", "blog_post.activated_at DESC").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()
	return posts, err
}

func (s *service) GetTags(search string, limit int) ([]models.TagWithUsage, error) {
	var tags []models.TagWithUsage
	// We want to select from tags and count usage in blog_post_tags.
	// Since TagWithUsage embeds Tag, we can select columns for Tag and the extra column.
	
	query := s.DB.Model(&tags).
		ColumnExpr("tag.*").
		ColumnExpr("count(bpt.blog_post_id) as usage_count").
		Join("LEFT JOIN blog_post_tags bpt ON bpt.tag_id = tag.id").
		Group("tag.id")

	if search != "" {
		query.Where("tag.name ILIKE ?", fmt.Sprintf("%%%s%%", search))
	}

	err := query.Order("usage_count DESC").
		Limit(limit).
		Select()
	
	return tags, err
}

func (s *service) GetAllHomeContent(page, limit int) ([]models.HomeContent, error) {
	content := make([]models.HomeContent, 0)
	err := s.DB.Model(&content).
		Where("home_content.activated_at IS NOT NULL AND home_content.activated_at < NOW() AND (home_content.deactivated_at IS NULL OR home_content.deactivated_at > NOW())").
		Order("home_content.ordering ASC", "home_content.activated_at DESC").
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
		Order("groove_jr_content.ordering ASC", "groove_jr_content.activated_at DESC").
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
		Order("about_content.ordering ASC", "about_content.activated_at DESC").
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
