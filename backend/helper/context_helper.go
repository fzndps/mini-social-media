package helper

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/fzndps/mini-social-media/backend/constant"
)

type contextKey string

const UserInfoKey contextKey = "userInfo"

type UserInfo struct {
	Id       int
	Username string
}

// Ambil userId dari context
func GetUserIdFromRequest(r *http.Request) (int, error) {
	claimsRaw := r.Context().Value(constant.UserInfoKey)
	claims, ok := claimsRaw.(*JWTClaim)
	if !ok {
		log.Println("❌ GAGAL: JWT context bukan *JWTClaim")
		return 0, errors.New("id not found") // bisa ganti dengan custom error kalau mau
	}

	log.Println("✅ User ID dari JWT:", claims.Id)
	return claims.Id, nil
}

// Ambil username dari context
func GetUsernameFromContext(ctx context.Context) string {
	claims, ok := ctx.Value(UserInfoKey).(*JWTClaim)
	if !ok || claims == nil {
		return ""
	}
	return claims.Username
}

// Fungsi gabungan kalau mau sekaligus
func GetUserInfoFromContext(ctx context.Context) UserInfo {
	claims, ok := ctx.Value(UserInfoKey).(*JWTClaim)
	if !ok || claims == nil {
		return UserInfo{}
	}
	return UserInfo{
		Id:       claims.Id,
		Username: claims.Username,
	}
}
