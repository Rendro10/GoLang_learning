package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", " D:/sqlite_db/check.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable(db *sql.DB) error {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            age INTEGER NOT NULL
        );`

	_, err := db.Exec(createTableSQL)
	return err
}

type User struct {
	ID   int
	Name string
	Age  int
}

// Create
func createUser(db *sql.DB, user User) (int64, error) {
	insertSQL := `INSERT INTO users (name, age) VALUES (?, ?)`
	result, err := db.Exec(insertSQL, user.Name, user.Age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Read
func getUserByID(db *sql.DB, id int) (User, error) {
	querySQL := `SELECT id, name, age FROM users WHERE id = ?`
	row := db.QueryRow(querySQL, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Update
func updateUser(db *sql.DB, user User) error {
	updateSQL := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err := db.Exec(updateSQL, user.Name, user.Age, user.ID)
	return err
}

// Delete
func deleteUser(db *sql.DB, id int) error {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(deleteSQL, id)
	return err
}

func main() {
	// Connect to the database
	db, err := connectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Create the users table
	err = createTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Perform CRUD operations
	user := User{Name: "John Doe", Age: 30}
	userID, err := createUser(db, user)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// Read the user
	fetchedUser, err := getUserByID(db, int(userID))
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}
	fmt.Println("Fetched user:", fetchedUser)

	// Update the user
	fetchedUser.Name = "Jane Doe"
	fetchedUser.Age = 28
	err = updateUser(db, fetchedUser)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}

	// Read the updated user
	updatedUser, err := getUserByID(db, int(userID))
	if err != nil {
		fmt.Println("Error fetching updated user:", err)
		return
	}
	fmt.Println("Updated user:", updatedUser)

	// Delete the user
	err = deleteUser(db, int(userID))
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
}
