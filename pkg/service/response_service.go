package service

import (
	"github.com/google/go-tika/tika"
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/jinzhu/gorm"
)

// ResponseService used for communicating with the DB
type DBService struct {
	DB         db.Response
	TikaClient *tika.Client
}

// DB struct provides access to various helper methods for querying data from the Twitter services database
type DB struct {
	db *gorm.DB
}

// New creates a new ResponseService struct for communicating with the core response service.
func New(db db.Response, tc *tika.Client) *DBService {

	t := &DBService{
		DB:         db,
		TikaClient: tc,
	}

	return t
}
