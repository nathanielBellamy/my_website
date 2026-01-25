package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/config"
)

func NewDBClient(cfg *config.Config) (*pg.DB, error) {
	opt, err := pg.ParseURL(cfg.Db.Url)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	return db, nil
}
