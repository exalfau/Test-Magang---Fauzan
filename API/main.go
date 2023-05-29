package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	// Membuat koneksi ke MySQL
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/login")
	if err != nil {
		log.Fatal(err)
	}

	// Memeriksa koneksi database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	// Membuat tabel users jika belum ada
	createTable(db)

	// Menginisialisasi rute-rute API
	e.GET("/users", getUsers(db))
	e.GET("/users/:id", getUser(db))
	e.POST("/users", createUser(db))
	e.PUT("/users/:id", updateUser(db))
	e.DELETE("/users/:id", deleteUser(db))
	e.GET("/resetid", resetID(db))

	// Menjalankan server pada port 8080
	e.Start(":8080")
}

// Membuat tabel users jika belum ada
func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		name VARCHAR(50),
		email VARCHAR(50),
		PRIMARY KEY (id)
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler untuk mendapatkan semua pengguna
func getUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to get users")
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				log.Fatal(err)
				return c.JSON(http.StatusInternalServerError, "Failed to get users")
			}
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	}
}

// Handler untuk mendapatkan pengguna berdasarkan ID
func getUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

		var user User
		err = row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, "User not found")
			}
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to get user")
		}

		return c.JSON(http.StatusOK, user)
	}
}

// Handler untuk membuat pengguna baru
func createUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user User
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to create user")
		}

		id, _ := result.LastInsertId()
		user.ID = int(id)

		return c.JSON(http.StatusCreated, user)
	}
}

// Handler untuk mengupdate pengguna
func updateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		var user User
		err = c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		result, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to update user")
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, "User not found")
		}

		return c.JSON(http.StatusOK, user)
	}
}

// Handler untuk menghapus pengguna
func deleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to delete user")
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, "User not found")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// Handler untuk mereset ID menjadi 1
func resetID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := db.Exec("TRUNCATE TABLE users")
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to reset ID")
		}

		_, err = db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "Failed to reset ID")
		}

		return c.JSON(http.StatusOK, "ID reset to 1")
	}
}
