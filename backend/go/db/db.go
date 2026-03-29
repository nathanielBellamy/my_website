package db

import (
	"context"
	"time"

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

func (pqa *PgQueryAdapter) ColumnExpr(expr string, params ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.ColumnExpr(expr, params...)}
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

func (pqa *PgQueryAdapter) Order(orders ...string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Order(orders...)}
}

func (pqa *PgQueryAdapter) Group(columns ...string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Group(columns...)}
}

func (pqa *PgQueryAdapter) Join(join string, params ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Join(join, params...)}
}

func (pqa *PgQueryAdapter) Select(dest ...interface{}) error {
	return pqa.Query.Select(dest...)
}

func (pqa *PgQueryAdapter) SelectAndCount(dest ...interface{}) (int, error) {
	return pqa.Query.SelectAndCount(dest...)
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

func (pqa *PgQueryAdapter) OnConflict(s string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.OnConflict(s)}
}

func (pqa *PgQueryAdapter) Set(s string) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pqa.Query.Set(s)}
}

// PgDBAdapter adapts *pg.DB to interfaces.PgxDB
type PgDBAdapter struct {
	*pg.DB
}

// Model implements interfaces.PgxDB.Model
func (pda *PgDBAdapter) Model(model ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pda.DB.Model(model...)}
}

// RunInTransaction implements interfaces.PgxDB.RunInTransaction
func (pda *PgDBAdapter) RunInTransaction(fn func(interfaces.PgxDB) error) error {
	return pda.DB.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		return fn(&PgTxAdapter{Tx: tx})
	})
}

// PgTxAdapter adapts *pg.Tx to interfaces.PgxDB
type PgTxAdapter struct {
	*pg.Tx
}

func (pta *PgTxAdapter) Model(model ...interface{}) interfaces.PgxQuerySeter {
	return &PgQueryAdapter{pta.Tx.Model(model...)}
}

func (pta *PgTxAdapter) RunInTransaction(fn func(interfaces.PgxDB) error) error {
	return pta.Tx.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		return fn(&PgTxAdapter{Tx: tx})
	})
}

func NewDBClient(cfg *config.Config) (*pg.DB, error) {
	opt, err := pg.ParseURL(cfg.Db.Url)
	if err != nil {
		return nil, err
	}

	// Add robust connection pooling to prevent silent drop hanging
	opt.PoolSize = 20
	opt.MinIdleConns = 5
	opt.IdleTimeout = 30 * time.Minute
	opt.IdleCheckFrequency = 1 * time.Minute
	opt.ReadTimeout = 10 * time.Second
	opt.WriteTimeout = 10 * time.Second
	opt.DialTimeout = 5 * time.Second

	db := pg.Connect(opt)
	return db, nil
}

// NewPgDBAdapter creates a new PgDBAdapter
func NewPgDBAdapter(db *pg.DB) interfaces.PgxDB {
	return &PgDBAdapter{DB: db}
}
