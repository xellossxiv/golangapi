package main

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() (*sql.DB, error) {
	cfg := mysql.Config{
		// User:   "adminuser",
		// Passwd: "Panjang@dm!n12#",
		// Net:    "tcp",
		// Addr:   "172.17.32.1:3306",
		// DBName: "testapp3",
		User:                 "u216092168_targetApp01",
		Passwd:               "P@ssw0rdpanjang",
		Net:                  "tcp",
		Addr:                 "sql503.main-hosting.eu",
		DBName:               "u216092168_targetApp01",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
