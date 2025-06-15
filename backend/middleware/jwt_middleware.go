package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/fzndps/mini-social-media/backend/constant"
	"github.com/fzndps/mini-social-media/backend/helper"

	"github.com/julienschmidt/httprouter"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("📥 Masuk ke middleware ProtectedRoute")

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
		ctx := context.WithValue(r.Context(), constant.UserInfoKey, claims)
		log.Println("✅ Token valid, lanjut ke handler...")

		next.ServeHTTP(w, r.WithContext(ctx)) // meneruskan ke request handler berikutnya
	})
}

// Bungkus endpoint yang butuh JWT
func ProtectedRoute(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, ps) // tetap passing ps ke handler
		})).ServeHTTP(w, r)
	}
}
