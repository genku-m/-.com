package repository_test

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	testDB   *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	dbDriver := "mysql"
	c := mysql.Config{
		DBName:               "test_main",
		User:                 "test_user",
		Passwd:               "password",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	testDB, err = sql.Open(dbDriver, c.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	fixtures, err = testfixtures.New(
		testfixtures.Database(testDB),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("./testfixtures"),
	)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
