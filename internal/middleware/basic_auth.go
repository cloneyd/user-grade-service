package middleware

import (
	"net/http"
	"os"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			expectedUsername, expectedPassword := os.Getenv("BASIC_USERNAME"), os.Getenv("BASIC_PASSWORD")
			username, password, ok := r.BasicAuth()

			if ok {
				usernameMatch := username == expectedUsername
				passwordMatch := password == expectedPassword

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
}
