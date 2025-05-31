package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/fzndps/mini-social-media/backend/helper"
)

type contextKey string

const userInfoKey contextKey = "userInfo"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") { // Periksa bearer apakah ada
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ") // hapus bearer
		claims, err := helper.ValidateJWT(tokenString)           // Validasi token
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Simpan data user dari token ke context le
		ctx := context.WithValue(r.Context(), userInfoKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx)) // meneruskan ke request handler berikutnya
	})
}
