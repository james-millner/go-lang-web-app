package db

import (
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/jinzhu/gorm"
)

// User provides access to Users via various methods
type Response interface {
	NewRecord(r *model.Response) bool
	FindBySourceURL(sourceUrl string) []*model.Response
	FindAll() []*model.Response
	Save(r *model.Response) *model.Response
}

type DB struct {
	db *gorm.DB
}

// New creates a new instance of DB and returns a reference to it
func New(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) NewRecord(u *model.Response) bool {
	return d.db.NewRecord(u)
}

func (d *DB) FindBySourceURL(sourceUrl string) []*model.Response {
	var r model.Response
	r.SourceURL = sourceUrl
	d.db.Where(&r).First(&r)

	return nil
}

func (d *DB) FindAll() []*model.Response {
	var responses []*model.Response
	d.db.Find(&responses)

	return responses
}

func (d *DB) Save(r *model.Response) *model.Response {
	if d.db.NewRecord(r) {
		d.db.Create(&r)
	} else {
		d.db.Save(&r)
	}

	return r
}
