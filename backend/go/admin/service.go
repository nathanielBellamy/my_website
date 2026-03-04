package admin

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/rs/zerolog"
)

type Service interface {
	// Blog
	GetAllBlogPosts(filter models.FilterOptions) ([]models.BlogPost, int, error)
	GetBlogPostByID(id string) (*models.BlogPost, error)
	GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error)
	GetTags(search string, limit int) ([]models.TagWithUsage, error)
	CreateBlogPost(post *models.BlogPost) (*models.BlogPost, error)
	UpdateBlogPost(post *models.BlogPost) (*models.BlogPost, error)
	DeleteBlogPost(id string) error

	// Home
	GetAllHomeContent(filter models.FilterOptions) ([]models.HomeContent, int, error)
	GetHomeContentByID(id string) (*models.HomeContent, error)
	CreateHomeContent(content *models.HomeContent) (*models.HomeContent, error)
	UpdateHomeContent(content *models.HomeContent) (*models.HomeContent, error)
	DeleteHomeContent(id string) error

	// GrooveJr
	GetAllGrooveJrContent(filter models.FilterOptions) ([]models.GrooveJrContent, int, error)
	GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error)
	CreateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	UpdateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	DeleteGrooveJrContent(id string) error

	// About
	GetAllAboutContent(filter models.FilterOptions) ([]models.AboutContent, int, error)
	GetAboutContentByID(id string) (*models.AboutContent, error)
	CreateAboutContent(content *models.AboutContent) (*models.AboutContent, error)
	UpdateAboutContent(content *models.AboutContent) (*models.AboutContent, error)
	DeleteAboutContent(id string) error

	// Images
	UploadImage(filename, originalName, altText string) (*models.Image, error)
	ListImages() ([]models.Image, error)
	DeleteImage(id string) error

	// CSV Export/Import
	ExportBlogPosts() ([]models.BlogPost, error)
	ImportBlogPosts(posts []models.BlogPost) error

	ExportHomeContent() ([]models.HomeContent, error)
	ImportHomeContent(content []models.HomeContent) error

	ExportGrooveJrContent() ([]models.GrooveJrContent, error)
	ImportGrooveJrContent(content []models.GrooveJrContent) error

	ExportAboutContent() ([]models.AboutContent, error)
	ImportAboutContent(content []models.AboutContent) error

	ExportTags() ([]models.Tag, error)
	ImportTags(tags []models.Tag) error

	ExportAuthors() ([]models.Author, error)
	ImportAuthors(authors []models.Author) error
}

type service struct {
	DB  interfaces.PgxDB
	Log *zerolog.Logger
}

func NewService(db interfaces.PgxDB, log *zerolog.Logger) Service {
	return &service{DB: db, Log: log}
}

func mapSortField(field string) string {
	switch field {
	case "activatedAt":
		return "activated_at"
	case "deactivatedAt":
		return "deactivated_at"
	case "createdAt":
		return "created_at"
	case "updatedAt":
		return "updated_at"
	default:
		return field
	}
}

// Blog
func (s *service) GetAllBlogPosts(filter models.FilterOptions) ([]models.BlogPost, int, error) {
	var posts []models.BlogPost
	query := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags")

	switch filter.Status {
	case "current":
		query.Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at <= NOW() AND (blog_post.deactivated_at IS NULL OR blog_post.deactivated_at > NOW())")
	case "inactive":
		query.Where("blog_post.activated_at IS NULL AND blog_post.deactivated_at IS NULL")
	case "past":
		query.Where("blog_post.deactivated_at IS NOT NULL AND blog_post.deactivated_at < NOW()")
	case "future":
		query.Where("blog_post.activated_at IS NOT NULL AND blog_post.activated_at > NOW()")
	}

	if len(filter.Tags) > 0 {
		query.Where("blog_post.id IN (SELECT blog_post_id FROM blog_post_tags WHERE tag_id IN (?) GROUP BY blog_post_id HAVING count(distinct tag_id) = ?)", pg.In(filter.Tags), len(filter.Tags))
	}

	if filter.SortField != "" {
		order := "ASC"
		if filter.SortOrder == "desc" || filter.SortOrder == "DESC" {
			order = "DESC"
		}
		field := mapSortField(filter.SortField)
		if field == "activated_at" || field == "deactivated_at" || field == "created_at" || field == "updated_at" {
			field = "blog_post." + field
		}
		query.Order(field + " " + order)
	} else {
		query.Order("blog_post.ordering ASC", "blog_post.activated_at DESC")
	}

	count, err := query.Limit(filter.Limit).
		Offset((filter.Page - 1) * filter.Limit).
		SelectAndCount()
	return posts, count, err
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

func (s *service) GetTags(search string, limit int) ([]models.TagWithUsage, error) {
	var tags []models.TagWithUsage
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

func (s *service) CreateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	s.Log.Info().Interface("post", post).Msg("Initial post")

	// 1. Handle Author
	if post.Author != nil {
		// If author has an ID, try to find it. If not, create it.
		if post.Author.ID == "" {
			// Check if author with this name already exists
			var existingAuthor models.Author
			err := s.DB.Model(&existingAuthor).Where("name = ?", post.Author.Name).Select()
			if err == nil {
				post.Author = &existingAuthor
			} else {
				_, err := s.DB.Model(post.Author).Insert()
				if err != nil {
					s.Log.Error().Err(err).Msg("Error inserting author")
					return nil, err
				}
			}
		}
		post.AuthorID = post.Author.ID
	}
	s.Log.Info().Interface("post", post).Msg("Post after author handling")

	// 2. Insert the BlogPost itself to generate the ID.
	_, err := s.DB.Model(post).Insert()
	if err != nil {
		s.Log.Error().Err(err).Msg("Error inserting blog post")
		return nil, err
	}
	s.Log.Info().Interface("post", post).Msg("Post after insert")

	// 3. Handle Tags
	if len(post.Tags) > 0 {
		var newTags []*models.Tag
		for _, tag := range post.Tags {
			var existingTag models.Tag
			err := s.DB.Model(&existingTag).Where("name = ?", tag.Name).Select()
			if err == nil {
				newTags = append(newTags, &existingTag)
			} else {
				newTag := models.Tag{Name: tag.Name}
				_, err := s.DB.Model(&newTag).Insert()
				if err != nil {
					s.Log.Error().Err(err).Msg("Error inserting tag")
					return nil, err
				}
				newTags = append(newTags, &newTag)
			}
		}
		post.Tags = newTags
		s.Log.Info().Interface("post", post).Msg("Post after tag handling")

		// 4. Create new tag associations.
		if len(post.Tags) > 0 {
			var blogPostTags []models.BlogPostTag
			for _, tag := range post.Tags {
				blogPostTags = append(blogPostTags, models.BlogPostTag{
					BlogPostID: post.ID,
					TagID:      tag.ID,
				})
			}
			s.Log.Info().Interface("tags", blogPostTags).Msg("BlogPostTags to insert")
			s.Log.Info().Msg("Attempting to insert BlogPostTags")
			_, err = s.DB.Model(&blogPostTags).Insert()
			if err != nil {
				s.Log.Error().Err(err).Msg("Error inserting blog post tags")
				return nil, err
			}
			s.Log.Info().Msg("Successfully inserted BlogPostTags")
		}
	}

	return post, nil
}

func (s *service) UpdateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	// 1. Handle Author
	if post.Author != nil {
		// If author has an ID, try to find it. If not, create it.
		if post.Author.ID == "" {
			// Check if author with this name already exists
			var existingAuthor models.Author
			err := s.DB.Model(&existingAuthor).Where("name = ?", post.Author.Name).Select()
			if err == nil {
				post.Author = &existingAuthor
			} else {
				_, err := s.DB.Model(post.Author).Insert()
				if err != nil {
					return nil, err
				}
			}
		}
		post.AuthorID = post.Author.ID
	}

	// 2. Update the BlogPost itself (title, content, author_id, updated_at, activated_at, deactivated_at).
	_, err := s.DB.Model(post).
		Column("title", "content", "author_id", "updated_at", "activated_at", "deactivated_at", "ordering").
		Where("id = ?", post.ID).
		Update()
	if err != nil {
		return nil, err
	}

	// 3. Delete existing tag associations.
	_, err = s.DB.Model((*models.BlogPostTag)(nil)).
		Where("blog_post_id = ?", post.ID).
		Delete()
	if err != nil {
		return nil, err
	}

	// 4. Handle Tags
	if len(post.Tags) > 0 {
		var newTags []*models.Tag
		for _, tag := range post.Tags {
			var existingTag models.Tag
			err := s.DB.Model(&existingTag).Where("name = ?", tag.Name).Select()
			if err == nil {
				newTags = append(newTags, &existingTag)
			} else {
				newTag := models.Tag{Name: tag.Name}
				_, err := s.DB.Model(&newTag).Insert()
				if err != nil {
					return nil, err
				}
				newTags = append(newTags, &newTag)
			}
		}
		post.Tags = newTags

		// 5. Create new tag associations.
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
func (s *service) GetAllHomeContent(filter models.FilterOptions) ([]models.HomeContent, int, error) {
	var content []models.HomeContent
	query := s.DB.Model(&content)

	switch filter.Status {
	case "current":
		query.Where("activated_at IS NOT NULL AND activated_at <= NOW() AND (deactivated_at IS NULL OR deactivated_at > NOW())")
	case "inactive":
		query.Where("activated_at IS NULL AND deactivated_at IS NULL")
	case "past":
		query.Where("deactivated_at IS NOT NULL AND deactivated_at < NOW()")
	case "future":
		query.Where("activated_at IS NOT NULL AND activated_at > NOW()")
	}

	if filter.SortField != "" {
		order := "ASC"
		if filter.SortOrder == "desc" || filter.SortOrder == "DESC" {
			order = "DESC"
		}
		query.Order(mapSortField(filter.SortField) + " " + order)
	} else {
		query.Order("ordering ASC", "activated_at DESC")
	}

	count, err := query.Limit(filter.Limit).
		Offset((filter.Page - 1) * filter.Limit).
		SelectAndCount()
	return content, count, err
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
	_, err := s.DB.Model(content).Column("title", "content", "activated_at", "deactivated_at", "ordering").Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteHomeContent(id string) error {
	_, err := s.DB.Model(&models.HomeContent{}).Where("id = ?", id).Delete()
	return err
}

// GrooveJr
func (s *service) GetAllGrooveJrContent(filter models.FilterOptions) ([]models.GrooveJrContent, int, error) {
	var content []models.GrooveJrContent
	query := s.DB.Model(&content)

	switch filter.Status {
	case "current":
		query.Where("activated_at IS NOT NULL AND activated_at <= NOW() AND (deactivated_at IS NULL OR deactivated_at > NOW())")
	case "inactive":
		query.Where("activated_at IS NULL AND deactivated_at IS NULL")
	case "past":
		query.Where("deactivated_at IS NOT NULL AND deactivated_at < NOW()")
	case "future":
		query.Where("activated_at IS NOT NULL AND activated_at > NOW()")
	}

	if filter.SortField != "" {
		order := "ASC"
		if filter.SortOrder == "desc" || filter.SortOrder == "DESC" {
			order = "DESC"
		}
		query.Order(mapSortField(filter.SortField) + " " + order)
	} else {
		query.Order("ordering ASC", "activated_at DESC")
	}

	count, err := query.Limit(filter.Limit).
		Offset((filter.Page - 1) * filter.Limit).
		SelectAndCount()
	return content, count, err
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
	_, err := s.DB.Model(content).Column("title", "content", "activated_at", "deactivated_at", "ordering").Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteGrooveJrContent(id string) error {
	_, err := s.DB.Model(&models.GrooveJrContent{}).Where("id = ?", id).Delete()
	return err
}

// About
func (s *service) GetAllAboutContent(filter models.FilterOptions) ([]models.AboutContent, int, error) {
	var content []models.AboutContent
	query := s.DB.Model(&content)

	switch filter.Status {
	case "current":
		query.Where("activated_at IS NOT NULL AND activated_at <= NOW() AND (deactivated_at IS NULL OR deactivated_at > NOW())")
	case "inactive":
		query.Where("activated_at IS NULL AND deactivated_at IS NULL")
	case "past":
		query.Where("deactivated_at IS NOT NULL AND deactivated_at < NOW()")
	case "future":
		query.Where("activated_at IS NOT NULL AND activated_at > NOW()")
	}

	if filter.SortField != "" {
		order := "ASC"
		if filter.SortOrder == "desc" || filter.SortOrder == "DESC" {
			order = "DESC"
		}
		query.Order(mapSortField(filter.SortField) + " " + order)
	} else {
		query.Order("ordering ASC", "activated_at DESC")
	}

	count, err := query.Limit(filter.Limit).
		Offset((filter.Page - 1) * filter.Limit).
		SelectAndCount()
	return content, count, err
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
	_, err := s.DB.Model(content).Column("title", "content", "activated_at", "deactivated_at", "ordering").Where("id = ?", content.ID).Update()
	return content, err
}

func (s *service) DeleteAboutContent(id string) error {
	_, err := s.DB.Model(&models.AboutContent{}).Where("id = ?", id).Delete()
	return err
}

// Images

func (s *service) UploadImage(filename, originalName, altText string) (*models.Image, error) {
	image := &models.Image{
		Filename:     filename,
		OriginalName: originalName,
		AltText:      altText,
	}
	_, err := s.DB.Model(image).Insert()
	return image, err
}

func (s *service) ListImages() ([]models.Image, error) {
	var images []models.Image
	err := s.DB.Model(&images).Order("created_at DESC").Select()
	return images, err
}

func (s *service) DeleteImage(id string) error {
	var image models.Image
	err := s.DB.Model(&image).Where("id = ?", id).Select()
	if err != nil {
		return err
	}

	// Delete from DB
	_, err = s.DB.Model(&image).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}

	// Delete from Disk
	uploadDir := "uploads/images"
	return os.Remove(fmt.Sprintf("%s/%s", uploadDir, image.Filename))
}

// CSV Export/Import Implementation

func (s *service) ExportBlogPosts() ([]models.BlogPost, error) {
	var posts []models.BlogPost
	err := s.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Order("ordering ASC", "activated_at DESC").
		Select()
	return posts, err
}

func (s *service) ImportBlogPosts(posts []models.BlogPost) error {
	// Using transaction for bulk operation
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, post := range posts {
			// Handle Author
			if post.Author != nil {
				if post.Author.ID != "" {
					post.AuthorID = post.Author.ID
					_, err := tx.Model(post.Author).
						OnConflict("(id) DO UPDATE").
						Set("name = EXCLUDED.name, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
						Insert()
					if err != nil {
						return err
					}
				} else if post.Author.Name != "" {
					// Check if author exists by name
					var existing models.Author
					err := tx.Model(&existing).Where("name = ?", post.Author.Name).Select()
					if err == nil {
						post.AuthorID = existing.ID
						post.Author = &existing
					} else {
						_, err := tx.Model(post.Author).Insert()
						if err != nil {
							return err
						}
						post.AuthorID = post.Author.ID
					}
				}
			}

			// Handle Tags
			var tagIDs []string
			if len(post.Tags) > 0 {
				for _, tag := range post.Tags {
					if tag.ID != "" {
						_, err := tx.Model(tag).
							OnConflict("(id) DO UPDATE").
							Set("name = EXCLUDED.name, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
							Insert()
						if err != nil {
							return err
						}
						tagIDs = append(tagIDs, tag.ID)
					} else if tag.Name != "" {
						var existing models.Tag
						err := tx.Model(&existing).Where("name = ?", tag.Name).Select()
						if err == nil {
							tagIDs = append(tagIDs, existing.ID)
						} else {
							_, err := tx.Model(tag).Insert()
							if err != nil {
								return err
							}
							tagIDs = append(tagIDs, tag.ID)
						}
					}
				}
			}

			// Upsert Post
			if post.ID != "" {
				_, err := tx.Model(&post).
					OnConflict("(id) DO UPDATE").
					Set("title = EXCLUDED.title, content = EXCLUDED.content, author_id = EXCLUDED.author_id, ordering = EXCLUDED.ordering, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&post).Insert()
				if err != nil {
					return err
				}
			}

			// Handle Post-Tags relations
			// First, remove existing relations if we are updating, or just ensure we add new ones.
			// Ideally, we sync: delete those not in list, add those in list.
			// For simplicity: delete all for this post, then add all.
			_, err := tx.Model((*models.BlogPostTag)(nil)).
				Where("blog_post_id = ?", post.ID).
				Delete()
			if err != nil {
				return err
			}

			if len(tagIDs) > 0 {
				var blogPostTags []models.BlogPostTag
				for _, tID := range tagIDs {
					blogPostTags = append(blogPostTags, models.BlogPostTag{
						BlogPostID: post.ID,
						TagID:      tID,
					})
				}
				_, err = tx.Model(&blogPostTags).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *service) ExportHomeContent() ([]models.HomeContent, error) {
	var content []models.HomeContent
	err := s.DB.Model(&content).Order("ordering ASC", "activated_at DESC").Select()
	return content, err
}

func (s *service) ImportHomeContent(content []models.HomeContent) error {
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, item := range content {
			if item.ID != "" {
				_, err := tx.Model(&item).
					OnConflict("(id) DO UPDATE").
					Set("title = EXCLUDED.title, content = EXCLUDED.content, ordering = EXCLUDED.ordering, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&item).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *service) ExportGrooveJrContent() ([]models.GrooveJrContent, error) {
	var content []models.GrooveJrContent
	err := s.DB.Model(&content).Order("ordering ASC", "activated_at DESC").Select()
	return content, err
}

func (s *service) ImportGrooveJrContent(content []models.GrooveJrContent) error {
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, item := range content {
			if item.ID != "" {
				_, err := tx.Model(&item).
					OnConflict("(id) DO UPDATE").
					Set("title = EXCLUDED.title, content = EXCLUDED.content, ordering = EXCLUDED.ordering, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&item).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *service) ExportAboutContent() ([]models.AboutContent, error) {
	var content []models.AboutContent
	err := s.DB.Model(&content).Order("ordering ASC", "activated_at DESC").Select()
	return content, err
}

func (s *service) ImportAboutContent(content []models.AboutContent) error {
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, item := range content {
			if item.ID != "" {
				_, err := tx.Model(&item).
					OnConflict("(id) DO UPDATE").
					Set("title = EXCLUDED.title, content = EXCLUDED.content, ordering = EXCLUDED.ordering, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&item).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *service) ExportTags() ([]models.Tag, error) {
	var tags []models.Tag
	err := s.DB.Model(&tags).Order("name ASC").Select()
	return tags, err
}

func (s *service) ImportTags(tags []models.Tag) error {
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, tag := range tags {
			if tag.ID != "" {
				_, err := tx.Model(&tag).
					OnConflict("(id) DO UPDATE").
					Set("name = EXCLUDED.name, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&tag).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *service) ExportAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := s.DB.Model(&authors).Order("name ASC").Select()
	return authors, err
}

func (s *service) ImportAuthors(authors []models.Author) error {
	return s.DB.RunInTransaction(func(tx interfaces.PgxDB) error {
		for _, author := range authors {
			if author.ID != "" {
				_, err := tx.Model(&author).
					OnConflict("(id) DO UPDATE").
					Set("name = EXCLUDED.name, activated_at = EXCLUDED.activated_at, deactivated_at = EXCLUDED.deactivated_at").
					Insert()
				if err != nil {
					return err
				}
			} else {
				_, err := tx.Model(&author).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
