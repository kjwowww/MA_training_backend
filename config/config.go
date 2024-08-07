package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration represents the application configuration
type Configuration struct {
	Port         string
	Environment  string
	DBName       string
	URI          string
	Path         string
	JWTAlgorithm string
	JWTSecret    string
	JWTSignature string
}

var AppConfig Configuration
var DB *mongo.Database

func LoadConfig() {
	viper.AddConfigPath(".")    // Look for the config file in the current directory
	viper.SetConfigName(".env") // Name of the config file
	viper.SetConfigType("env")  // The type of the config file

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.SetDefault("JWT_ALGORITHM", "HS256")
	viper.SetDefault("JWT_SECRET", "your-secret-key")

	requiredKeys := []string{"port", "environment", "dbname", "uri", "path", "jwt_algorithm", "jwt_secret", "jwt_signature"}
	for _, key := range requiredKeys {
		if !viper.IsSet(key) {
			log.Fatalf("Config key %s is missing", key)
		}
	}

	AppConfig = Configuration{
		Port:         viper.GetString("port"),
		Environment:  viper.GetString("environment"),
		DBName:       viper.GetString("dbname"),
		URI:          viper.GetString("uri"),
		Path:         viper.GetString("path"),
		JWTAlgorithm: viper.GetString("jwt_algorithm"),
		JWTSecret:    viper.GetString("jwt_secret"),
		JWTSignature: viper.GetString("jwt_signature"),
	}
}

func ConnectMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(AppConfig.URI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	DB = client.Database(AppConfig.DBName)
	log.Println("Connected to MongoDB!")
}
