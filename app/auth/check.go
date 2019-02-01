package auth

import(
	"net/http"
	"fmt"
	"os"
	"strings"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("DEFAULT_JWT_SIGNING_KEY"))

func getToken(r *http.Request) (string, error){
	
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) == 2 && (auth[0] != "Bearer" || auth[0] != "bearer"){
		return auth[1], nil
	}else{
		return "", errors.New("Authorization header does not contain correct token format. Format should be 'bearer <JWT Token>'")
    }
}

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tokenString, err := getToken(r); err!=nil {
			fmt.Fprintf(w, "Not Authorized")
		}else{
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		}
	})
}