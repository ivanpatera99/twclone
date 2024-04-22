package middleware

import (
	"context"
	"net/http"

	"github.com/ivanpatera/twclone/pkg/auth"
)



func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userId := r.Header.Get("x-user-id")
		username := r.Header.Get("x-username")
		
		user, err := auth.GetUser(userId, username)
		if err != nil {
			if err.Error() == "INVALID_USER" {
				http.Error(w, "INVALID_USER", http.StatusForbidden)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Enrich the request context with the user data
		ctx := context.WithValue(r.Context(), auth.UserIDKey, user.ID)
		ctx = context.WithValue(ctx, auth.UsernameKey, user.Username)

		// Call the next handler with the updated request context
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}