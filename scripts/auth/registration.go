package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/youthcongress/nepal/database"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles the registration form submission
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	username := r.FormValue("register-form-username")
	email := r.FormValue("register-form-email")
	mobile := r.FormValue("register-form-phone")
	password := r.FormValue("register-form-password")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Connect to the database
	db, err := database.Connection() // ✅ Correct: Handle both return values
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close() // ✅ Close DB when done


	// Insert user into `members` table
	query := "INSERT INTO members (username, email, mobile, password) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("❌ Database prepare error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, email, mobile, string(hashedPassword))
	if err != nil {
		log.Println("❌ Insert error:", err)
		http.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "✅ Registration successful!")
}