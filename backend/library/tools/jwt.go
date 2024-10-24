package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, name, secret string, expireMinutes int64) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireMinutes) * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, fmt.Errorf("token expired: %v", err)
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func RefreshAccessToken(refreshToken, accessSecret, refreshSecret string, accessExpire int64) (string, error) {
	claims, err := VerifyToken(refreshToken, refreshSecret)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %v", err)
	}

	newAccessToken, err := GenerateToken(claims.UserID, claims.Username, accessSecret, accessExpire)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %v", err)
	}

	return newAccessToken, nil
}

func GenerateDoubleToken(userID int64, name, accessSecret, refreshSecret string, accessExpire, refreshExpire int64) (string, string, error) {
	accessToken, err := GenerateToken(userID, name, accessSecret, accessExpire)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := GenerateToken(userID, name, refreshSecret, refreshExpire)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return accessToken, refreshToken, nil
}
