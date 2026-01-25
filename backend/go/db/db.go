package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/marketing" // Import marketing package
)

// PgQueryAdapter adapts *pg.Query to marketing.PgxQuerySeter
type PgQueryAdapter struct {
	*pg.Query
}

func (pqa *PgQueryAdapter) Relation(name string) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Relation(name)}
}

func (pqa *PgQueryAdapter) Limit(count int) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Limit(count)}
}

func (pqa *PgQueryAdapter) Offset(offset int) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Offset(offset)}
}

func (pqa *PgQueryAdapter) Where(query string, params ...interface{}) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Where(query, params...)}
}

func (pqa *PgQueryAdapter) Join(join string, params ...interface{}) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Join(join, params...)}
}

func (pqa *PgQueryAdapter) Select(dest ...interface{}) error {
	return pqa.Query.Select(dest...)
}

// PgDBAdapter adapts *pg.DB to marketing.PgxDB
type PgDBAdapter struct {
	*pg.DB
}

// Model implements marketing.PgxDB.Model
func (pda *PgDBAdapter) Model(model ...interface{}) marketing.PgxQuerySeter {
	return &PgQueryAdapter{pda.DB.Model(model...)}
}

func NewDBClient(cfg *config.Config) (*pg.DB, error) {
	opt, err := pg.ParseURL(cfg.Db.Url)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	return db, nil
}

// NewPgDBAdapter creates a new PgDBAdapter
func NewPgDBAdapter(db *pg.DB) marketing.PgxDB {
	return &PgDBAdapter{DB: db}
}
