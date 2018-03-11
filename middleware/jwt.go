package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var (
	env       string
	publicKey string
)

func init() {
	env = os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	publicKey = os.Getenv("JWT_PUBLIC_KEY")
}

// JWT is a middleware that checks income request based in JWT
func JWT(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if env == "test" {
			h.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		isBearerAuth := strings.HasPrefix(authHeader, "Bearer ")

		if isBearerAuth {
			tokenString := authHeader[len("Bearer "):]
			// If something went wrong with public key, put down the server
			publicRSA, err := parseRSAPublicKey(publicKey)
			if err != nil {
				log.Fatalf("could not parse public key: %v", err)
			}

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return publicRSA, err
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if token.Valid {
				h.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func parseRSAPublicKey(keyFile string) (*rsa.PublicKey, error) {

	publicKey, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %v", err)
	}
	publicRSA, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %s. %v", publicKey, err)
	}

	return publicRSA, nil
}
