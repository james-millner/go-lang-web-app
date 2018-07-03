package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"database/sql"
	"log"
	"path/filepath"
	"os"

	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/james-millner/go-lang-web-app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Service struct {
	Storage *gorm.DB
	Router 	*mux.Router
	debug   bool
}

type Config struct {
	DBPort      int    	`default:"3306"`
	Debug     	bool   	`default:"false"`
	DBDialect 	string 	`required:"false"`
	Hostname	string	`default:"localhost"`
	DBDsn    	string
}

func main() {

	var env Config
	err := envconfig.Process("cs", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := openDBConnection(&env)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	service := Service{Storage: db, Router: router}

	port := ":4000"

	service.setRouters()

	fmt.Println("Listening on: ", port)
	http.ListenAndServe(port, router)
}

func (a *Service) setRouters() {
	a.Post("/gather-links", a.gatherLinks)
	a.Post("/handle-link", a.handleLinks)
}

//handler method
func (a *Service) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (s *Service) gatherLinks(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	url := r.FormValue("url")

	if url != "" {

		var links []model.Links
		var documents []model.Links

		for _, t := range getLinks(url) {
			if strings.Contains(t, ".pdf") {
				documents = append(documents, model.Links{Url: t})
			} else {
				links = append(links, model.Links{Url: t})
			}
		}

		resp := &model.Response{Links: links, Success: true, Documents: documents, Url: url}

		for l := range links {
			fmt.Println(links[l])
		}

		fmt.Println("Total links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))

		enc.Encode(resp)

		s.Storage.Save(&resp)

	} else {
		resp := &model.Response{Success: false}
		enc.Encode(resp)
	}
}

func (s *Service) handleLinks(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	selector := r.FormValue("selector")
	tag := r.FormValue("tag")

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	resp := &model.IndividualLinkResponse{Url: url, Selector: selector, Tag: tag}
	enc.Encode(resp)
}

func getLinks(url string) []string {

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		errFmt := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(errFmt)
		return nil
	}

	return web.GetPageLinks(r)
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
		&model.Links{},
	)

	return gormDB, nil
}
