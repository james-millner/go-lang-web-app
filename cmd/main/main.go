package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/james-millner/go-lang-web-app/pkg/handlers"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kelseyhightower/envconfig"
	goji "goji.io"
	"goji.io/pat"
)

//Service struct for holding core service related dependencies/
type Service struct {
	Storage *gorm.DB
	Router  *mux.Router
	debug   bool
}

//Config struct for holding environment variables.
type Config struct {
	HTTPPort  int    `default:"8811"`
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

	database := db.New(gormDB)

	rs := service.New(database)

	fmt.Println("Listening on: ", env.HTTPPort)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(env.HTTPPort),
		Handler: handlersMethod(rs),
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Printf("HTTP Server shutdown!")
	}
}

func handlersMethod(rs *service.ResponseService) *goji.Mux {
	router := goji.NewMux()

	user := handlers.NewResponseService(rs)
	router.HandleFunc(pat.Post("/gather-links"), user.GatherLinks())

	return router
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
