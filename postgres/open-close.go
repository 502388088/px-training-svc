package postgres

import (
	//"database/sql"
	_ "encoding/json"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"github.com/jmoiron/sqlx"
	"github.com/502388088/px-training-svc/vcap"
	"fmt"
)

const ServiceKey = "VCAP_SERVICES"

var db_initialized = false
var Database *sqlx.DB
var Schema string

func Open_postgres(defaultUrl string, schema string) {
	var db_url = GetDbUrlFromPredix(defaultUrl)

	if db_url == "" {
		log.Println("using default env var")
		db_url = defaultUrl
	}

	log.Println(fmt.Sprintf("db_url: %s", db_url))

	db, err := sqlx.Connect("postgres", db_url)
	if err != nil {
		log.Fatal(err)
	}

	Database = db
	db_initialized = true

	//  configure database pooling
	mconns := os.Getenv("MAX_DB_CONNECTIONS")
	maxconns := 20
	if (mconns != "") { maxconns,_ = strconv.Atoi(mconns) }

	SetMaxIdleConns(1)
	SetMaxOpenConns(maxconns)

	//  do we need to verify that the table(s) exist??
	log.Println("Database Connection Established")

}

func GetDbUrlFromPredix(schema string) (db_url string){
	vcapServices, _ := vcap.LoadServices()
	postgresServices := vcapServices["postgres"]
	Schema = schema

	lenPostgresServices := len(postgresServices)
	if (lenPostgresServices == 1){
		vmap := postgresServices[0].Credentials
		db_url = vmap["uri"].(string)
	} else if(lenPostgresServices > 1) {
		database_name := os.Getenv("DB_POSTGRES_NAME")
		for i := range postgresServices {
			if postgresServices[i].Name == database_name {
				vmap := postgresServices[i].Credentials
				db_url = vmap["uri"].(string)
			}
		}
	}
	return
}

func Close_postgres() {

	log.Println("Closing Database Connections")

	if db_initialized {
		Database.Close()
		db_initialized = false
	}

}