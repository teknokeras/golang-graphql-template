package auth

import(
	"net/http"
	"fmt"
	"os"
	"strings"
	"errors"
	"bytes"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/handler"
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

func isAuthRequired(r *http.Request) bool{

	buf, _ := ioutil.ReadAll(r.Body)
	
	rdrAuth := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdrTemp := ioutil.NopCloser(bytes.NewBuffer(buf))
	
	r.Body = rdrAuth

	reqOptions := handler.NewRequestOptions(r)

	queryString := strings.Split(reqOptions.Query,"{")

	isQuery := (strings.TrimSpace(queryString[0]) == "Query") || (strings.TrimSpace(queryString[0]) == "query")

	if !isQuery {
		//mutation must have auth header
		return true
	}
	
	secondQueryString := strings.TrimSpace(queryString[1])
	isLogin := strings.HasPrefix(secondQueryString,"Login")

	r.Body = rdrTemp

	if isLogin {
		return false
	}

	return true
}

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !isAuthRequired(r) {
			next.ServeHTTP(w, r)
			return
		}

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