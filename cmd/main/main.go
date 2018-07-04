package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/james-millner/go-lang-web-app/pkg/handlers"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kelseyhightower/envconfig"
)

type Service struct {
	Storage *gorm.DB
	Router  *mux.Router
	debug   bool
}

type Config struct {
	DBPort    int    `default:"3306"`
	Debug     bool   `default:"false"`
	DBDialect string `required:"false"`
	Hostname  string `default:"localhost"`
	DBDsn     string
}

func main() {

	var env Config
	err := envconfig.Process("cs", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	gormDB, err := openDBConnection(&env)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	service := Service{Storage: gormDB, Router: router}

	port := ":4000"

	database := db.New(service.Storage)

	tw := handlers.New(database)

	service.setRouters()

	fmt.Println("Listening on: ", port)
	http.ListenAndServe(port, router)
}

func (a *Service) setRouters() {
	a.Post("/gather-links", handlers)
}

//handler method
func (a *Service) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func openDBConnection(config *Config) (*gorm.DB, error) {
	var gormDB *gorm.DB
	var err error

	switch config.DBDialect {
	case "mysql":
		dbDSN := config.DBDsn

		if config.DBDsn == "" {
			dbDSN = "root@tcp(" + config.Hostname + ":3306)/iqblade-casestudy?charset=utf8&parseTime=True"
		}

		db, err := sql.Open("mysql", dbDSN)
		if err != nil {
			log.Fatalf("Failed to load mysql driver: %v", err)
		}

		gormDB, err = gorm.Open("mysql", db)
	default:
		gormDB, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to database: %v", err)
	}

	gormDB.AutoMigrate(
		&model.Response{},
	)

	return gormDB, nil
}
