IQBlade - Case Study Microservice
----------------------------

[![CircleCI](https://circleci.com/gh/james-millner/go-lang-web-app/tree/master.svg?style=svg&circle-token=b97b68792491c3010205c810362a0e99b1b81db4)](https://circleci.com/gh/james-millner/go-lang-web-app/tree/master)

A simple Case Study Microservice in GoLang. The app boots up and provides an API to interact and use the service. 

##### Pre Requisites 

- The application requries elastic search to be running somewhere on your machine / environment. You can pass as a environment variable to the application: `ElasticURL` which needs to be the full reference to elasticsearch, such as: `https://localhost:9200`

- The applicaiton will also boot up an instance of Apache Tika, please note you can configure the port this runs on by following the Build & Run app stage.

The application has two endpoints: 

- Gather links; `/gather-links` requires a POST Request with a parameter of `url`.
This endpoint will go to the passed URL and scan for web content / links and return possible documents and links as JSON. It will also be stored locally in a DB for the service to log.

- Process Link; `/process-link` requires a POST request with two parameters of; `url` & `company_number` 
This endpoint will go to a found case study URL provided as a p arameter and bind it to the company number parameter provided. The service will then try to extract text from the document and analyze company related text. This will then be bound to a DTO and exported into an Elasticsearch Client.

Warning Beginner Gopher here.

Build & Run app
---------------
```
	HTTPPort   int    `default:"8811"`
	DBPort     int    `default:"3306"`
	Debug      bool   `default:"false"`
	DBDialect  string `required:"false"`
	Hostname   string `default:"localhost"`
	ElasticURL string `default:"http://localhost:9200"`
	TikaPort   string `default:"9998"`
	DBDsn      string
```

The applicaiton has the above environment variables that can be configured as the program is built & run. 

To run the application simply type the following run command:

1. `make build` - This will build the binary app and run all associated tests.
2. Configure the environment variables by prefixing each variable with `CS_`. For example `CS_DBDIALECT=mysql main`, or for multiple: `CS_DBDIALECT=mysql CS_ELASTICURL=https://localhost:4242 main`
3. Watch the logs for application boot. 

