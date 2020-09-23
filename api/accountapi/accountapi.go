package accountapi

import (
	"LearningProgramming/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecretKey"

func CreateToken(res http.ResponseWriter, req *http.Request) {
	var account entities.Account

	err := json.NewDecoder(req.Body).Decode(&account)

	if err != nil {
		ResponseWithJson(res, http.StatusBadRequest, err.Error())
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": account.UserName,
			"password": account.Password,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err2 := token.SignedString([]byte(secretKey))

		if err2 != nil {
			ResponseWithJson(res, http.StatusBadRequest, err2.Error())
		} else {
			ResponseWithJson(res, http.StatusOK, entities.Token{Token: tokenString})
		}
	}

}

func CheckToken(res http.ResponseWriter, req *http.Request) {
	tokenString := req.Header.Get("key")

	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err == nil && result.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("InValid")
	}
}

func ResponseWithError(res http.ResponseWriter, code int, message string) {
	ResponseWithJson(res, code, map[string]string{"error": message})
}

func ResponseWithJson(res http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(response)
}
