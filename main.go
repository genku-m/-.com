package main

import (
	"database/sql"
	"time"

	"github.com/genku-m/upsider-cording-test/guid"
	"github.com/genku-m/upsider-cording-test/invoice/repository"
	invoice_usecase "github.com/genku-m/upsider-cording-test/invoice/usecase"
	"github.com/genku-m/upsider-cording-test/server"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

func main() {
	cfg := server.NewConfig()
	dbDriver := "mysql"
	c := mysql.Config{
		DBName:    "test_main",
		User:      "test_user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.UTC,
	}
	db, err := sql.Open(dbDriver, c.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	svr := server.NewServer(invoice_usecase.NewInvoiceUsecase(guid.New(), repository.NewInvoiceRepository(db)), cfg)
	svr.Listen()
}
