package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Define your database credentials
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Rickarka@445"
	dbname   = "CheckGoLang"
)

// Define a struct to represent the data you want to store in the table
type Item struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	// Create the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Ensure the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	// Create the table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			name TEXT,
			price NUMERIC
		)
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// CRUD Operations

	// Create operation

	// item1 := Item{Name: "Apple", Price: 80}

	// err = createItem(db, &item1)
	// if err != nil {
	// 	log.Fatal("Error creating item1:", err)
	// }

	// Read operation

	items, err := readItems(db)
	if err != nil {
		log.Fatal("Error reading items:", err)
	}

	fmt.Println("Items:")
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
	}

	// Update operation

	// itemToUpdate := Item{ID: 1, Name: "Updated Item", Price: 15.99}
	// err = updateItem(db, &itemToUpdate)
	// if err != nil {
	// 	log.Fatal("Error updating item:", err)
	// }

	// Read the updated item

	// updatedItem, err := readItemByID(db, itemToUpdate.ID)
	// if err != nil {
	// 	log.Fatal("Error querying the updated item:", err)
	// }

	// fmt.Println("Updated Item:")
	// fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", updatedItem.ID, updatedItem.Name, updatedItem.Price)

	// Delete operation
	itemToDeleteID := 6
	err = deleteItem(db, itemToDeleteID)
	if err != nil {
		log.Fatal("Error deleting item:", err)
	}

	fmt.Printf("Item with ID %d has been deleted.\n", itemToDeleteID)
}

// CRUD Functions

// Create a new item
func createItem(db *sql.DB, item *Item) error {
	err := db.QueryRow("INSERT INTO items (name, price) VALUES ($1, $2) RETURNING id", item.Name, item.Price).Scan(&item.ID)
	if err != nil {
		return err
	}
	return nil
}

// Read all items
func readItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, price FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		items = append(items, item)
	}

	return items, nil
}

// Read an item by its ID
func readItemByID(db *sql.DB, id int) (Item, error) {
	var item Item
	err := db.QueryRow("SELECT id, name, price FROM items WHERE id=$1", id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

// Update an item
func updateItem(db *sql.DB, item *Item) error {
	_, err := db.Exec("UPDATE items SET name=$1, price=$2 WHERE id=$3", item.Name, item.Price, item.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete an item by its ID
func deleteItem(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
