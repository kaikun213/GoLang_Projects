package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := openDB("smarterTransferDB") // alternative: smarterTestDB
	defer db.Close()
	rows, errRows := db.Query("show tables;")
	if errRows != nil {
		log.Fatalf("error in main: %s", errRows.Error())
	}
	truncateDB(rows, db)
}

func openDB(databaseName string) *sql.DB {
	// initialize database sql
	db, err := sql.Open("mysql", "root:team2@tcp(localhost:3306)/"+databaseName)
	err = db.Ping()
	if err != nil {
		log.Fatalf("error establisting connection: %s", err.Error())
	}
	log.Printf("Connections established: %d", db.Stats().OpenConnections)
	return db
}

func truncateDB(rows *sql.Rows, db *sql.DB) {
	var table string
	// unset foreign key checks
	_, err := db.Exec(`set FOREIGN_KEY_CHECKS = 0;`)
	if err != nil {
		log.Fatalf("error setting foreing key checks: %s", err.Error())
	}
	// truncate to all tables
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatalf("error getting table names: %s", err.Error())
		}
		result, errX := db.Exec("Truncate table " + table + ";")
		if errX != nil {
			log.Fatalf("error truncating tables: %s", errX.Error())
		}
		log.Printf("Connections established: %d", db.Stats().OpenConnections)
		rows, err := result.RowsAffected()
		log.Printf("Result rows affected: %d", rows)
		log.Printf("Truncate table %s;", table)
	}
}
