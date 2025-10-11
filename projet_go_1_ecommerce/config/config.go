package config

import (
	"encoding/json" // Pour convertir du JSON en structures Go (équivalent à json en Python)
	"log"           // Pour afficher des messages et erreurs (équivalent à print + sys.exit)
	"os"            // Pour lire les fichiers du système

	"gorm.io/driver/postgres" // Driver PostgreSQL pour GORM
	"gorm.io/gorm"            // ORM principal (équivalent de SQLAlchemy en Python)
)

// Déclaration globale : accessible depuis tout le package
var DB *gorm.DB

func ConnectDB() {

	data, err := os.ReadFile(`C:\Users\flosr\Credentials\postgresql_creds.json`)
	// os.ReadFile lit le contenu d'un fichier et retourne deux valeurs :
	// 1. le contenu du fichier (data) de type []byte
	// 2. une erreur (err) si quelque chose ne va pas

	if err != nil {
		log.Fatal("Impossible to read the credentials file:\n", err)
	}
	// Vérifie si err est différent de nil (Go utilise nil comme None/null)
	// log.Fatal imprime le message et termine le programme si erreur

	var creds map[string]string
	// Déclare une variable creds qui est une map clé->valeur, avec clé et valeur de type string
	// La syntaxe map[cléType]valeurType est l'équivalent du dict Python

	if err := json.Unmarshal(data, &creds); err != nil {
		log.Fatal("JSON Parsing Error\n", err)
	}
	// json.Unmarshal convertit les données JSON (data) en structure Go
	// &creds passe l'adresse de la variable (passage par référence)
	// := ici crée err local pour cette ligne
	// if err != nil vérifie si parsing JSON a échoué
	// Cela correspond à json.loads(data) en Python

	connStr, ok := creds["connection_learn_golang"]
	// Récupère la valeur associée à la clé "connection_learn_golang" dans la map
	// ok est un booléen qui indique si la clé existait
	if !ok {
		log.Fatal("No connection string found in the credentials file")
	}
	// !ok est l'équivalent de "not ok" en Python

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	// gorm.Open ouvre une connexion à la DB avec le driver postgres
	// postgres.Open(connStr) fournit la string de connexion
	// &gorm.Config{} crée une instance vide de configuration (pointeur)
	// Note : DB est la variable globale déclarée plus haut

	if err != nil {
		log.Fatal("Error connecting to the database\n", err)
	}
	// Vérifie si la connexion a échoué et stop le programme si nécessaire

	log.Println("Database connection established")
	// log.Println affiche un message suivi d'un saut de ligne
}
