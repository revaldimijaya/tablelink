package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Create tables
	createTables(db)
	return db
}

func createTables(db *sqlx.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS tm_ingredient (
            uuid TEXT PRIMARY KEY,
            name TEXT NOT NULL UNIQUE,
            cause_alergy INTEGER,
            type INTEGER,
            status INTEGER,
            created_at DATETIME,
            updated_at DATETIME,
            deleted_at DATETIME
        );`,
		`CREATE TABLE IF NOT EXISTS tm_item (
            uuid TEXT PRIMARY KEY,
            name TEXT NOT NULL UNIQUE,
            price REAL,
            status INTEGER,
            created_at DATETIME,
            updated_at DATETIME,
            deleted_at DATETIME
        );`,
		`CREATE TABLE IF NOT EXISTS tm_item_ingredient (
            uuid_item TEXT,
            uuid_ingredient TEXT,
            FOREIGN KEY(uuid_item) REFERENCES tm_item(uuid),
            FOREIGN KEY(uuid_ingredient) REFERENCES tm_ingredient(uuid),
            PRIMARY KEY(uuid_item, uuid_ingredient)
        );`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}

	// Insert sample data
	insertSampleData(db)
}

func insertSampleData(db *sqlx.DB) {
	sampleDataQueries := []string{
		`INSERT OR IGNORE INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at) 
         VALUES ('8666d298-517b-45f3-8566-378cb5c8738c', 'Chicken', 0, 0, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at) 
         VALUES ('e97c6a5f-c541-4f3f-84d4-953c3eabe686', 'Pork', 0, 0, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at) 
         VALUES ('9a3a33cf-7144-4c5d-a0c6-fd8e894a0db5', 'Radish', 0, 2, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at) 
         VALUES ('b2752259-090e-4e6e-a9a1-2f47d538d833', 'Egg', 1, 1, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_item (uuid, name, price, status, created_at) 
         VALUES ('07419b87-4702-49f9-83aa-f9b489f64b14', 'Chicken Pork', 30000.00, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_item (uuid, name, price, status, created_at) 
         VALUES ('d290fc2b-6a32-4bfe-98c7-25b9884c5245', 'Chicken Pork with Radish', 35000.00, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_item (uuid, name, price, status, created_at) 
         VALUES ('7cc760a3-393b-493c-b780-3cfd7afd1cf9', 'Salad Egg', 20000.00, 1, datetime('now'))`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('07419b87-4702-49f9-83aa-f9b489f64b14', '8666d298-517b-45f3-8566-378cb5c8738c')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('07419b87-4702-49f9-83aa-f9b489f64b14', 'e97c6a5f-c541-4f3f-84d4-953c3eabe686')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('d290fc2b-6a32-4bfe-98c7-25b9884c5245', '8666d298-517b-45f3-8566-378cb5c8738c')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('d290fc2b-6a32-4bfe-98c7-25b9884c5245', 'e97c6a5f-c541-4f3f-84d4-953c3eabe686')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('d290fc2b-6a32-4bfe-98c7-25b9884c5245', '9a3a33cf-7144-4c5d-a0c6-fd8e894a0db5')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('7cc760a3-393b-493c-b780-3cfd7afd1cf9', '9a3a33cf-7144-4c5d-a0c6-fd8e894a0db5')`,
		`INSERT OR IGNORE INTO tm_item_ingredient (uuid_item, uuid_ingredient) 
         VALUES ('7cc760a3-393b-493c-b780-3cfd7afd1cf9', 'b2752259-090e-4e6e-a9a1-2f47d538d833')`,
	}

	for _, query := range sampleDataQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error inserting sample data: %v", err)
		}
	}
}
