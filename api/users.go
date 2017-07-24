package api

import (
	"crypto/sha256"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"os"
	"strings"
)

func Login(auth User, r Repository) (string, error) {
	err := auth.Validate()
	if err != nil {
		return "", err
	}

	user := User{}
	encrypted := createHash(auth.Password)

	collection := r.DB(os.Getenv("DB_NAME")).C("users")
	err = collection.Find(bson.M{"email": auth.Email, "password": encrypted}).One(&user)
	if err != nil {
		log.Errorf("API - Login - Fail to find: %s", err)
		return "", err
	}
	log.Infof("API - Login - Authenticated")

	var token string
	token, err = createToken(auth.Email)
	if err != nil {
		log.Errorf("API - Login - Fail to create token: %s", err)
		return "", err
	}

	return token, nil
}

func createHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func ValidateToken(hash string) bool {
	hash = strings.Replace(hash, "Bearer ", "", -1)
	token, err := jwt.Parse(hash, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret := []byte(os.Getenv("SECRET"))
		return hmacSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	log.Errorf("API - ValidateToken - Error: %s", err)
	return false
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	hmacSecret := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
