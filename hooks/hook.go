package main

import (
	"database/sql"
	"fmt"
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
		log.Fatalf("Error opening file: %v", err.Error())
	}
	err = f.Truncate(0)
	if err != nil {
		log.Fatalf("Error truncating the file: %v", err.Error())
	}
	defer f.Close()
	log.SetOutput(f)

	// initialize database sql
	db, err := sql.Open("mysql", "root:team2@tcp(localhost:3306)/smarterTransferDB")
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
	}

	// close after all other statements (defer)
	defer db.Close()

	h.BeforeAll(func(t []*trans.Transaction) {
		log.Printf("Truncate Function called from Transaction: %s ", "BeforeAll")
		truncateDB(db)
	})

	/* Merchants */
	h.Before("Merchants > GET Merchants", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})
	h.Before("Merchant > GET Merchant", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})
	h.Before("Merchant > UPDATE Merchant > Example 2", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})
	h.Before("Merchant > DELETE Merchant", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMerchant(db)
	})

	/* Users */
	h.Before("Users > ADD User", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		_, err := db.Exec(`insert into THEME (themeId) values (1);`)
		if err != nil {
			log.Fatalf("Error adding a user: %v", err.Error())
		}
	})
	h.Before("User > GET User", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addUser(db)
	})
	h.Before("User > UPDATE User > Example 2", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addUser(db)
	})
	h.Before("User > DELETE User", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addUser(db)
	})

	/* Items */
	h.Before("Items > ADD Item", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		/* add Merchant for dependency */
		addMerchant(db)
	})
	h.Before("Items > GET Items", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addItem(db)
	})
	h.Before("Item > GET Item", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addItem(db)
	})
	h.Before("Item > UPDATE Item > Example 2", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addItem(db)
	})
	h.Before("Item > DELETE Item", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addItem(db)
	})
	/* Menus */
	h.Before("Menus > ADD Menu", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		/* add Merchant for dependency */
		addMerchant(db)
	})
	h.Before("Menus > GET Menus", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMenu(db)
	})
	h.Before("Menu > GET Menu", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMenu(db)
	})
	h.Before("Menu > UPDATE Menu > Example 2", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMenu(db)
	})
	h.Before("Menu > DELETE Menu", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addMenu(db)
	})
	/* POS */
	h.Before("PointOfSales > ADD POS", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		/* add  Menu for dependency */
		addMenu(db)
	})
	h.Before("PointOfSales > GET POSs", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addPOS(db)
	})
	h.Before("POS > GET POS", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addPOS(db)
	})
	h.Before("POS > UPDATE POS > Example 2", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addPOS(db)
	})
	h.Before("POS > DELETE POS", func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s", "Before", t.Name)
		addPOS(db)
	})

	// clean database after each Transaction
	h.AfterEach(func(t *trans.Transaction) {
		log.Printf("Function: %s called from Transaction: %s \n\n", "AfterEach", t.Name)
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
	result, err := db.Exec(`insert into ADDRESS (country, city) values ("testCountry", "testCity");`)
	if err != nil {
		log.Fatalf("Error adding a merchant: %v", err.Error())
	}
	id, _ := result.LastInsertId()
	result, err = db.Exec(fmt.Sprintf(`insert into MERCHANT (keshId, companyName, phoneNumber, email,  addressId,
																					ustId, logoId, websiteURL, shopURL, ticketURL) values
																					(1, "testCompany", "testPhoneNumber", "testEmail", %d, "testUstId", 1,
																					"testWebsiteURL", "testShopURL", "testTicketURL");`, id))
	if err != nil {
		log.Fatalf("Error adding a merchant: %v", err.Error())
	}
	id, _ = result.LastInsertId()

	_, err = db.Exec(fmt.Sprintf(`insert into CONTACT (merchantId, forename, surname, email) values
																					(%d, "testContactForename", "testContactSurname", "testContactEmail");`, id))
	if err != nil {
		log.Fatalf("Error adding a merchant: %v", err.Error())
	}
}

// "userId": 1,
// "keshId": 1,
// "name": "testerInput",
// "location":{
// 	"lon": 12.21,
// 	"lat": 32.1
// },
// "deviceId": "testDevice",
// "themeId": 1
func addUser(db *sql.DB) {
	result, err := db.Exec(`insert into THEME (themeId) values (1);`)
	if err != nil {
		log.Fatalf("Error adding a user: %v", err.Error())
	}
	id, _ := result.LastInsertId()

	_, err = db.Exec(fmt.Sprintf(`insert into USER (keshId, name, longitude, latitude,themeId,deviceId) values
																											(1, "testName", 12.21, 32.1, %d, "testDeviceId");`, id))
	if err != nil {
		log.Fatalf("Error adding a user: %v", err.Error())
	}
}

// "merchantId":1,
// "itemId":1,
// "name":"testItem",
// "description":"testDescription",
// "price":20.0
func addItem(db *sql.DB) {
	/* Add merchant first for dependency */
	addMerchant(db)

	_, err := db.Exec(`insert into ITEM (merchantId, name, description, price) values (1, "testItem", "testDescription", 20.0);`)
	if err != nil {
		log.Fatalf("Error adding an item: %v", err.Error())
	}
}

// "merchantId": 1,
// "posId":1,
// "menuId": 1,
// "location":{
// 	"lon": 12.21,
// 	"lat": 32.1
// }
func addPOS(db *sql.DB) {
	/* Add menu & merchant first for dependency */
	addMenu(db)

	_, err := db.Exec(`insert into POS (merchantId, menuId, longitude, latitude) values (1, 1, 12.21, 31.1);`)
	if err != nil {
		log.Fatalf("Error adding a POS: %v", err.Error())
	}
}

// "merchantId": 1,
// "menuId": 1,
// "name": "testMenu",
// "items": [{
// 	"merchantId":1,
// 	"itemId":1,
// 	"name":"testItem",
// 	"description":"testDescription",
// 	"price":20.0
// 	}]
func addMenu(db *sql.DB) {
	/* add Merchant */
	//	addMerchant(db)
	/* add Menu */
	result, err := db.Exec(`insert into MENU (merchantId, name) values (1, "testMenu");`)
	if err != nil {
		log.Fatalf("Error adding a menu: %v", err.Error())
	}
	menuId, _ := result.LastInsertId()
	/* add Item */
	addItem(db)
	/* add Item to Menu */
	_, err = db.Exec(fmt.Sprintf(`insert into MENU_ITEM (menuId, itemId) values (%d, 1);`, menuId))
	if err != nil {
		log.Fatalf("Error adding an item to menu: %v", err.Error())
	}
}

func truncateDB(db *sql.DB) {
	var table string
	// Query all tables of the DB
	rows, errRows := db.Query("show tables;")
	if errRows != nil {
		log.Fatalf("Error: %v", errRows.Error())
	}
	// unset foreign key checks
	_, err := db.Exec(`set FOREIGN_KEY_CHECKS = 0;`)
	if err != nil {
		log.Fatalf("Error setting foreing key checks: %s", err.Error())
	}
	// truncate to all tables
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatalf("Error getting table names: %s", err.Error())
		}
		_, errX := db.Exec("Truncate table " + table + ";")
		if errX != nil {
			log.Fatalf("Error truncating tables.\nError: %s", errX.Error())
		}
		_, errX = db.Exec("ALTER table " + table + " AUTO_INCREMENT = 1;")
		if errX != nil {
			log.Fatalf("Error resetting AUTO_INCREMENT.\nError: %s", errX.Error())
		}
	}
}
