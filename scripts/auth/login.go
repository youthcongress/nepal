package auth

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/youthcongress/nepal/database"
	"golang.org/x/crypto/bcrypt"
)

func Login(app *fiber.App) {
	app.Post("/login", func(c *fiber.Ctx) error {
		// Retrieve form values
		username := c.FormValue("login-form-username")
		password := c.FormValue("login-form-password")

		// Basic validation
		if username == "" || password == "" {
			return c.Render("login", fiber.Map{"error": "Both fields are required"})
		}

		// Open database connection
		db, err := database.Connection()
		if err != nil {
			log.Println("Database connection failed:", err)
			return c.Render("login", fiber.Map{"error": "Database connection failed"})
		}
		defer db.Close()

		// Retrieve stored user details
		var storedPassword string
		query := "SELECT password FROM members WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Render("login", fiber.Map{"error": "Invalid username or password"})
			}
			log.Println("Error retrieving user:", err)
			return c.Render("login", fiber.Map{"error": "Failed to process login"})
		}

		// Compare hashed password
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			return c.Render("login", fiber.Map{"error": "Invalid username or password"})
		}

		// Successful login
		return c.Redirect("/") // Adjust the redirect URL as needed
	})
}