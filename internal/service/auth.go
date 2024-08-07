package service

import (
	"context"
	"errors"
	"time"

	"ma-backend-training/internal/repository"
	"ma-backend-training/internal/utils"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser creates a new user in the database
func CreateUser(firstName, lastName, username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := repository.User{
		ID:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return repository.CreateUser(context.Background(), user)
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (repository.User, error) {
	return repository.GetUserByUsername(context.Background(), username)
}

// CheckPassword compares the provided password with the hashed password
func CheckPassword(providedPassword, hashedPassword string) bool {
	return utils.CheckPasswordHash(providedPassword, hashedPassword)
}

// GenerateJWT generates a JWT token for the given username
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("your-secret-key") // Replace with your actual secret key
	return token.SignedString(jwtSecret)
}

// ParseJWT parses a JWT token and returns the username
func ParseJWT(tokenString string) (string, error) {
	jwtSecret := []byte("your-secret-key") // Replace with your actual secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return username, nil
}
