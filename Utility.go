package main

import (
	"database/sql"
	"fmt"
)

func CheckClientIP(clientIP string, db *sql.DB) bool {
	var result bool = false
	query := fmt.Sprintf("SELECT 1 AS result FROM listip where ip = '%s'", clientIP)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		result = true
	}

	defer rows.Close()
	return result
}
