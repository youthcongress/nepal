package auth

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/youthcongress/nepal/database"
	"golang.org/x/crypto/bcrypt"
)

func Register(app *fiber.App) {
	app.Post("/register", func(c *fiber.Ctx) error {
		// Retrieve form values
		firstName := c.FormValue("register-form-first-name")
		lastName := c.FormValue("register-form-last-name")
		middleName := c.FormValue("register-form-middle-name")
		username := c.FormValue("register-form-username")
		email := c.FormValue("register-form-email")
		mobile := c.FormValue("register-form-phone")
		password := c.FormValue("register-form-password")
		rePassword := c.FormValue("register-form-repassword")

		// Basic validation
		if firstName == "" || lastName == "" || username == "" || email == "" || mobile == "" || password == "" || rePassword == "" {
			return c.Render("register", fiber.Map{"error": "All fields are required"})
		}

		// Check if passwords match
		if password != rePassword {
			return c.Render("register", fiber.Map{"error": "Passwords do not match"})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password:", err)
			return c.Render("register", fiber.Map{"error": "Failed to process registration"})
		}

		// Capture current timestamp
		registrationTime := time.Now().Format("2006-01-02 15:04:05") // Format: YYYY-MM-DD HH:MM:SS

		// Open database connection
		db, err := database.Connection()
		if err != nil {
			log.Println("Database connection failed:", err)
			return c.Render("register", fiber.Map{"error": "Database connection failed"})
		}
		defer db.Close()

		// Insert user into database with timestamp
		query := "INSERT INTO members (first_name, last_name, middle_name, username, email, mobile, password, time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(query, firstName, lastName, middleName, username, email, mobile, string(hashedPassword), registrationTime)
		if err != nil {
			log.Println("Error inserting user:", err)
			return c.Render("register", fiber.Map{"error": "Failed to register user"})
		}

		// Redirect to login page after successful registration
		return c.Redirect("/login")
	})
}