package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// GenerateJWT generates a JWT token
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	var token *jwt.Token
	if viper.GetString("jwt_algorithm") == "RS256" {
		privateKey := viper.GetString("jwt_private_key")
		key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
		if err != nil {
			return "", err
		}
		token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		return token.SignedString(key)
	}

	secret := viper.GetString("jwt_secret")
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT validates a JWT token
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := viper.GetString("jwt_secret")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
