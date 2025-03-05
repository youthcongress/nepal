package auth

import (
	"github.com/gorilla/sessions"
)

// Global session store
var Store = sessions.NewCookieStore([]byte("super-secret-key"))
