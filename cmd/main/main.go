package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-tika/tika"
	"github.com/gorilla/mux"
	"github.com/james-millner/go-lang-web-app/pkg/db"
	"github.com/james-millner/go-lang-web-app/pkg/es"
	"github.com/james-millner/go-lang-web-app/pkg/handlers"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kelseyhightower/envconfig"
	"github.com/olivere/elastic"
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
	HTTPPort   int    `default:"8092"`
	Debug      bool   `default:"false"`
	DBDialect  string `required:"false"`
	DBHost     string `required:"true" default:"localhost"`
	DBPort     int    `required:"true" default:"3306"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBDatabase string `required:"true" default:"iqblade-casestudies"`
	ElasticURL string `default:"http://elasticsearch:9200"`
	TikaPort   string `default:"9998"`
	DBDsn      string
}

func main() {

	var env Config
	err := envconfig.Process("cs", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	gormDB, err := openDBConnection(&env)
	if err != nil {
		log.Fatalf("failed to establish DB Connection: %v: ", err)
	}

	esc, err := createElasticClient(env.ElasticURL)
	if err != nil {
		log.Fatalf("failed to load elastic client: %v", err)
	}

	log.Println("Starting Tika server... ")
	tikaserver, serr := tika.NewServer("tika-server-1.14.jar", env.TikaPort)
	if serr != nil {
		log.Fatal(err)
	}

	log.Println("Tika server created... waiting to start")

	er := tikaserver.Start(context.Background())
	if er != nil {
		log.Fatal(er)
	} else {
		log.Println("Tika running")
	}

	database := db.New(gormDB)

	tc := tika.NewClient(nil, tikaserver.URL())

	es := es.New(esc)

	cs := service.New(database, tc)

	log.Println("Listening on: ", env.HTTPPort)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(env.HTTPPort),
		Handler: handlersMethod(cs, tc, es),
	}

	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		tikaErr := tikaserver.Stop()

		if tikaErr != nil {
			log.Fatal(tikaErr)
		}

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Printf("HTTP Server shutdown!")
	}
}

func handlersMethod(rs *service.CaseStudyService, tika *tika.Client, es *es.Elastic) *goji.Mux {
	router := goji.NewMux()

	user := handlers.NewCaseStudyService(rs, tika, es)
	router.HandleFunc(pat.Post("/gather-links"), user.GatherLinks())
	router.HandleFunc(pat.Post("/process-link"), user.ProcessCaseStudyLink())

	router.HandleFunc(pat.Get("/healthz"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	return router
}

func openDBConnection(config *Config) (*gorm.DB, error) {
	var gormDB *gorm.DB
	var err error

	switch config.DBDialect {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBDatabase,
		)

		db, err := sql.Open("mysql", dsn)
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
		&model.CaseStudy{},
		&model.CaseStudyOrganisations{},
		&model.CaseStudyPeople{},
		&model.CaseStudyLocations{},
	)

	return gormDB, nil
}

func createElasticClient(url string) (*elastic.Client, error) {

	c, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create new elastic client: %v", err)
	}

	return c, nil
}
