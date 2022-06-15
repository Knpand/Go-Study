package middleware

import (
	"net/http"
)

func CorsMiddleware(w http.ResponseWriter) error {
	// protocol := "http://"
	// host := "localhost:3000"

	//CORSの設定
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	return nil
}


