package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces" // Import interfaces package
)

// PgQueryAdapter adapts *pg.Query to interfaces.PgxQuerySeter
type PgQueryAdapter struct {
	*pg.Query
}

func (pqa *PgQueryAdapter) Relation(name string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Relation(name)}
}

func (pqa *PgQueryAdapter) Column(columns ...string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Column(columns...)}
}

func (pqa *PgQueryAdapter) Limit(count int) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Limit(count)}
}

func (pqa *PgQueryAdapter) Offset(offset int) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Offset(offset)}
}

func (pqa *PgQueryAdapter) Where(query string, params ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Where(query, params...)}
}

func (pqa *PgQueryAdapter) Join(join string, params ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Join(join, params...)}
}

func (pqa *PgQueryAdapter) Select(dest ...interface{}) error {
	return pqa.Query.Select(dest...)
}

func (pqa *PgQueryAdapter) Insert(dest ...interface{}) (pg.Result, error) {
	return pqa.Query.Insert(dest...)
}

func (pqa *PgQueryAdapter) Update(dest ...interface{}) (pg.Result, error) {
	return pqa.Query.Update(dest...)
}

func (pqa *PgQueryAdapter) Delete(dest ...interface{}) (pg.Result, error) {
	return pqa.Query.Delete(dest...)
}

// PgDBAdapter adapts *pg.DB to interfaces.PgxDB
type PgDBAdapter struct {
	*pg.DB
}

// Model implements interfaces.PgxDB.Model
func (pda *PgDBAdapter) Model(model ...interface{}) interfaces.PgxQuerySeter {
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
func NewPgDBAdapter(db *pg.DB) interfaces.PgxDB {
	return &PgDBAdapter{DB: db}
}
