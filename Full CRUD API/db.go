// db.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	// Charge les variables depuis le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Récupère les credentials depuis l'environnement
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// DSN PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	// Connexion à la DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migration des modèles
	db.AutoMigrate(&Movie{})
	fmt.Println("Database connected and migrated")
}

// Mongo DB Connection

package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client
var MoviesCollection *mongo.Collection

func mongo_init() {
	// ⚠️ Utiliser l’URI depuis un .env en prod
	uri := "mongodb+srv://flosrv123:Nesrine123@myfirstmongodbcluster.mde7n.mongodb.net/?retryWrites=true&w=majority&appName=MyFirstMongoDbCluster"

	// Contexte avec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connexion au cluster
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ Erreur connexion MongoDB:", err)
	}

	// Ping implicite avec v2
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ Ping MongoDB échoué:", err)
	}

	// Sauvegarder client + collection
	MongoClient = client
	MoviesCollection = client.Database("API_Training_Data").Collection("Movies")

	fmt.Println("✅ Connecté à MongoDB Atlas, collection: Movies")
}
