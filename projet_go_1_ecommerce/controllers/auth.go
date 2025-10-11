package controllers

import (
	"net/http"

	"go_learn_ecommerce/config"
	"go_learn_ecommerce/models"
	"go_learn_ecommerce/utils/hash"
	"go_learn_ecommerce/utils/jwt"

	"github.com/gin-gonic/gin"
)

// Register : crée un nouvel utilisateur
func Register(c *gin.Context) {
	var user models.User

	// 1️⃣ Bind JSON du client dans le struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2️⃣ Hash le mot de passe
	hashedPwd, err := hash.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hash du mot de passe"})
		return
	}
	user.Password = hashedPwd

	// 3️⃣ Enregistre l'utilisateur dans la DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
		return
	}

	// 4️⃣ Retour succès
	c.JSON(http.StatusOK, gin.H{
		"message": "Utilisateur créé avec succès",
		"user_id": user.ID,
	})
}

// Login : vérifie email/password et génère un JWT
func Login(c *gin.Context) {
	var input models.UserLogin

	// 1️⃣ Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2️⃣ Cherche l'utilisateur par email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	// 3️⃣ Vérifie le mot de passe
	if !hash.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	// 4️⃣ Génère le token JWT
	token, err := jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de générer le token"})
		return
	}

	// 5️⃣ Retourne succès avec token
	c.JSON(http.StatusOK, gin.H{
		"message": "Connexion réussie",
		"token":   token,
	})
}
