package database

import (
	"assignment2/models"
	_ "database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "macbookpro"
	password = "postgress"
	dbname   = "assignment_2"
	dialect  = "postgres"
)

var (
// db *sql.DB
// err error
)

// GORM
func handleDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error occured while trying to connect to database:", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
	return db, nil
}

func GetDatabaseInstance() *gorm.DB {
	db, err := handleDatabaseConnection()
	if err != nil {
		log.Panic(err)
	}
	return db
}

// // Using SQL
// func handleDatabaseConnection() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open(dialect, psqlInfo)
// 	if err != nil {
// 		log.Panic("Error occured while trying to validate database arguments:", err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Panic("Error occured while trying to connect to database:", err)
// 	}
// }

// func handleTableCreate() {
// 	orderTable := `
// 	CREATE TABLE IF NOT EXISTS "orders" (
// 		order_id SERIAL PRIMARY KEY,
// 		customer_name VARCHAR(255) NOT NULL,
// 		ordered_at timestamptz DEFAULT now(),
// 		created_at timestamptz DEFUALT now(),
// 		updated_at timestamptz DEFAULT now()
// 	);
// 	`

// 	itemTable := `
// 	CREATE TABLE IF NOT EXISTS "items" (
// 		item_id SERIAL PRIMARY KEY,
// 		item_code VARCHAR(191) NOT NULL,
// 		quantity INT NOT NULL,
// 		description TEXT NOT NULL,
// 		order_id INT NOT NULL,
// 		created_at timestamptz DEFAULT now(),
// 		updated_at timestamptz DEFAULT now(),
// 		CONSTRAINT items_order_id_fk FOREIGN KEY(order_id) REFERENCES orders(order_id) ON DELETE CASCADE
// 	);
// 	`

// 	_, err = db.Exec(orderTable)
// 	if err != nil {
// 		log.Panic("Error occured while trying to create order table:", err)
// 	}

// 	_, err = db.Exec(itemTable)
// 	if err != nil {
// 		log.Panic("Error occured while trying to create item table:", err)
// 	}
// }

func InitializeDatabase() {
	handleDatabaseConnection()
	// handleTableCreate()
}

// func GetDatabaseInstance() *sql.DB {
// 	if db == nil {
// 		log.Panic("database instance nil")
// 	}
// 	return db
// }
