package main

import (
	"LearningProgramming/api/accountapi"
	"LearningProgramming/api/demo_api"
	"LearningProgramming/middlewares/jwtAuth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")

	router := mux.NewRouter()

	router.HandleFunc("/api/account/generatekey", accountapi.CreateToken).Methods("POST")
	router.HandleFunc("/api/account/checkkey", accountapi.CheckToken).Methods("GET")

	router.Handle("/api/demo/demo1", jwtAuth.JWTAuth(http.HandlerFunc(demo_api.Demo1))).Methods("GET")
	router.Handle("/api/demo/demo2", jwtAuth.JWTAuth(http.HandlerFunc(demo_api.Demo2))).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
	}
}
