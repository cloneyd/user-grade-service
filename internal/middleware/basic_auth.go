package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"

	"wb-test-task-2022/internal/convert"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			expectedUsername, expectedPassword := os.Getenv("BASIC_USERNAME"), os.Getenv("BASIC_PASSWORD")
			username, password, ok := r.BasicAuth()

			if ok {
				usernameHash := sha256.Sum256(convert.StringToBytes(username))
				passwordHash := sha256.Sum256(convert.StringToBytes(password))

				expectedUsernameHash := sha256.Sum256(convert.StringToBytes(expectedUsername))
				expectedPasswordHash := sha256.Sum256(convert.StringToBytes(expectedPassword))

				usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
				passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
}
