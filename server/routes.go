package server

import (
	"net/http"

	"github.com/youthcongress/nepal/scripts/auth"
	"github.com/youthcongress/nepal/scripts/profile"
)

// Route mapping
var routes = map[string]string{
	"/login":   "./public/login.html",
	
}

// SetupRoutes registers all routes dynamically
func SetupRoutes() {
	// Register static file routes
	for route, filePath := range routes {
		http.HandleFunc(route, func(fp string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, fp)
			}
		}(filePath))
	}

	// Register API routes
	http.HandleFunc("/auth", auth.LoginHandler)   // Handles login authentication
	http.HandleFunc("/profile", profile.ProfileHandler) // Profile page with session
	http.HandleFunc("/logout", auth.LogoutHandler) // Logout
	http.HandleFunc("/profile/data", profile.ProfileDataHandler) // Serve username as JSON

}
