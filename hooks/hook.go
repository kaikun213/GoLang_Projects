package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(h)

	// initialize log file
	f, err := os.OpenFile("hooksLogFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err.Error())
	}
	err = f.Truncate(0)
	if err != nil {
		log.Fatalf("error truncating the file: %v", err.Error())
	}
	defer f.Close()
	log.SetOutput(f)

	// initialize database sql
	db, err := sql.Open("mysql", "root:team2@tcp(localhost:3306)/smarterTransferDB")
	err = db.Ping()
	if err != nil {
		log.Fatalf("error connection to database: %v", err)
	}

	// close after all other statements (defer)
	defer db.Close()

	// Define all your event callbacks here
	h.Before("Merchants > GET Merchants", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})

	h.Before("Items > GET Items", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})

	// clean database after each Transaction
	h.AfterEach(func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "AfterEach", t.Name)
		truncateDB(db)
	})

	// server.Serve() will block and allow the goodman server to run your defined
	// event callbacks
	server.Serve()
	// You must close the listener at end of main()
	defer server.Listener.Close()
}

// "merchantId": 1,
// "keshId": 1,
// "companyName": "testCompany",
// "phoneNumber": "testPhoneNumber",
// "email": "testEmail",
// "contactForename": "testContactForename",
// "contactSurname": "testContactSurname",
// "contactEmail": "testContactEmail",
// "country": "testCountry",
// "city": "testCity",
// "ustId": "testUstId",
// "logoId": 1,
// "websiteURL": "testWebsiteURL",
// "shopURL": "testShopURL",
// "ticketURL": "testTicketURL"
func addMerchant(db *sql.DB) {
	var err error
	_, err = db.Exec(`insert into MERCHANT (keshId, companyName, phoneNumber, email,
		 																			contactForename, contactSurname, contactEmail, country, city,
																					ustId, logoId, websiteURL, shopURL, ticketURL) values
																					(1, "testCompany", "testPhoneNumber", "testEmail", "testContactForename"
																					"testContactSurname", "testContactEmail", "testCountry", "testCity", "testUstId", 1
																					"testWebsiteURL", "testShopURL", "testTicketURL");`)
	if err != nil {
		log.Fatalf("error adding a merchant: %v", err.Error())
	}
}

func addUser(db *sql.DB) {

}

func addItem(db *sql.DB) {
	var err error
	_, err = db.Exec(`insert into ITEM (name, description, price) values ("hooksTestItem", "description", 10.0);`)
	if err != nil {
		log.Fatalf("error adding an item: %v", err.Error())
	}
}

func truncateDB(db *sql.DB) {
	var table string
	// Query all tables of the DB
	rows, errRows := db.Query("show tables;")
	if errRows != nil {
		log.Fatalf("error: %v", errRows.Error())
	}
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
