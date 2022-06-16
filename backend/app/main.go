package main

import (
	"log"
	"net/http"
	"app/handler"
	_ "github.com/go-sql-driver/mysql"	//これをimportしないとDB接続できない
	// "app/middleware"
)



func main() {

	http.HandleFunc("/", handler.Handle(handler.Login_handler))
	http.HandleFunc("/signup", handler.Handle(handler.Signup_handler))

	log.Fatal(http.ListenAndServe(":5050", nil))

}



