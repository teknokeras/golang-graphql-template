package passwordutils

import (
	"os"
	"time"
	"errors"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("DEFAULT_JWT_SIGNING_KEY"))

func GenerateJwtToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	    "authorized": true,
	    "client": email,
	    "exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	if tokenString, err := token.SignedString(signingKey); err != nil {
		return "", err
	}else{
		return tokenString, nil
	}
}

func ValidateJwtToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("There was an error parsing the token")
		}
		return signingKey, nil
	})

	if err != nil {
		return false, errors.New("Cannot parse token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Print(claims)
		return true, nil
	} else {
		return false, nil
	}
}

