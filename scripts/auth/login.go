package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/youthcongress/nepal/database"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles user authentication and starts a session
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
	username := r.FormValue("login-form-username")
	password := r.FormValue("login-form-password")

	// Connect to database
	db, err := database.Connection()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Retrieve user ID and hashed password from database
	var userID int
	var storedHashedPassword string
	query := "SELECT id, password FROM members WHERE username = ?"
	err = db.QueryRow(query, username).Scan(&userID, &storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			log.Println("❌ Database query error:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a session and store user ID & username
	session, _ := Store.Get(r, "session")
	session.Values["userID"] = userID
	session.Values["username"] = username
	session.Options.MaxAge = 3600 // 1-hour session timeout
	session.Save(r, w)

	// Redirect to /profile after login
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
	fmt.Println("✅ Login successful! Redirecting to profile.")
}
