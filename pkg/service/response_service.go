package service

import (
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/jinzhu/gorm"
)

// Twitter is used for communicating with Twitters APIs
type ResponseService struct {
	DB db.Response
}

// DB struct provides access to various helper methods for querying data from the Twitter services database
type DB struct {
	db *gorm.DB
}

// New creates a new Twitter struct for communicating with Twitter APIs
func New(db db.Response) *ResponseService {

	t := &ResponseService{
		DB: db,
	}

	return t
}
