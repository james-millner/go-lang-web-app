package service

import (
	"github.com/google/go-tika/tika"
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/jinzhu/gorm"
)

// CaseStudyService used for communicating with the DB
type CaseStudyService struct {
	DB         *db.DB
	TikaClient *tika.Client
}

// DB struct provides access to various helper methods for querying data from the MySQL database
type DB struct {
	db *gorm.DB
}

// New creates a new ResponseService struct for communicating with the core response service.
func New(db *db.DB, tc *tika.Client) *CaseStudyService {

	t := &CaseStudyService{
		DB:         db,
		TikaClient: tc,
	}

	return t
}
