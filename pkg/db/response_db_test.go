package db

import (
	"fmt"
	"math/rand"

	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/jinzhu/gorm"
)

// setupTestDB creates a in memory test sqlite3 database for testing
func setupTestDB(src rand.Source) (*gorm.DB, error) {
	r1 := rand.New(src)
	r1.Int63()

	db, err := gorm.Open("sqlite3", fmt.Sprintf("file:%d?mode=memory", r1.Int63()))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.Response{},
	)

	return db, nil
}

// func TestNewRecord(t *testing.T) {
// 	src := rand.NewSource(rand.Int63())

// 	db, err := setupTestDB(src)
// 	assert.NoError(t, err)

// 	user1 := &model.Response{
// 		SourceURL: "https://news.ycombinator.com/news",
// 	}

// 	db.Create(&user1)
// }
