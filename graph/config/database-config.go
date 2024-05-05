package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMysqlDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func SetupMongoDBConnection() (*mongo.Client, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		return nil, fmt.Errorf("failed to load env: %v", errEnv)
	}

	MONGO_URI := os.Getenv("MONGODB_URI")
	if MONGO_URI == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	return client, nil
}

func CloseMysqlDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection")
	}
	dbSQL.Close()
}

func CloseMongoDBConnection(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %v", err)
	}
	return nil
}
