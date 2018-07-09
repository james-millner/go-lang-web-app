package db

import (
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/jinzhu/gorm"
)

// Response interface for getting responses from MySQL.
type Response interface {
	FindBySourceURLAndURLFound(source string, url string) *model.Response
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

//FindBySourceURLAndURLFound method
func (d *DB) FindBySourceURLAndURLFound(sourceurl string, urlfound string) *model.Response {
	var c model.Response
	c.URLFound = urlfound
	c.SourceURL = sourceurl
	d.db.Where(&c).First(&c)

	return &c
}

//FindBySourceURL method
func (d *DB) FindBySourceURL(sourceUrl string) []*model.Response {
	var r model.Response
	r.SourceURL = sourceUrl
	d.db.Where(&r).First(&r)

	return nil
}

//FindAll method
func (d *DB) FindAll() []*model.Response {
	var responses []*model.Response
	d.db.Find(&responses)

	return responses
}

//Save method
func (d *DB) Save(r *model.Response) *model.Response {
	if d.db.NewRecord(r) {
		d.db.Create(&r)
	} else {
		d.db.Save(&r)
	}

	return r
}
