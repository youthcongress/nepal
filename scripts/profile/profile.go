package profile

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/youthcongress/nepal/scripts/auth" // Import session store
)

// ProfileHandler serves profile.html and injects the logged-in username
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "session")

	// Check if the user is logged in
	userID, ok := session.Values["userID"].(int)
	username, _ := session.Values["username"].(string)
	if !ok || userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Println("✅ User logged in. Serving profile.html for:", username)
	http.ServeFile(w, r, "./public/profile.html")
}

// ProfileDataHandler sends the logged-in username as JSON
func ProfileDataHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "session")

	// Check if user is logged in
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"username": username})
}