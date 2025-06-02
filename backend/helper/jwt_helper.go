package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaim struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Membuat token JWT
func GenerateJWT(id int, username string) (string, error) {
	//EXP token
	expTime := time.Now().Add(1 * time.Hour)

	claims := JWTClaim{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Membuat token dengan claim diatas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey) //Menandatangani dengan jwtKey
}

func ValidateJWT(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken, // Token yang dikirim user
		&JWTClaim{}, // Struct yang dibuat diatas untuk menampung token
		func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		},
	)

	// ok untuk cek apakah bersail konversi ke struct JWTClaim
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims, nil // balikin id dama username yang lomgin
	} else {
		return nil, err
	}
}
