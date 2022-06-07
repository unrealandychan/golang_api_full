package db

import (
	"TechSchoolGolang/util"
	"database/sql"

	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)


var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../")
	if err !=nil{
		log.Fatal("Cannot load configuration:" ,err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
