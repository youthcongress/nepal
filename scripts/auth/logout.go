package auth

import "net/http"

// LogoutHandler clears the session and redirects to login
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	// Clear session
	session.Options.MaxAge = -1
	session.Save(r, w)

	// Redirect to login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
