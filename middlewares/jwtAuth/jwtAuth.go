package jwtAuth

import (
	"LearningProgramming/api/accountapi"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecretKey"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("key")

		if tokenString == "" {
			accountapi.ResponseWithError(res, http.StatusUnauthorized, "Unauthorized")
		} else {
			result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err == nil && result.Valid {
				next.ServeHTTP(res, req)
			} else {
				accountapi.ResponseWithError(res, http.StatusUnauthorized, "Unauthorized")
			}
		}
	})
}
