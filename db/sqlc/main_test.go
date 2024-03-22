package db

import (
	"database/sql"
	"log"
	"os"
	"simple_bank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}
