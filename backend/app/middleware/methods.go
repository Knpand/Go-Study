package middleware
import (
	"net/http"
	"errors"
)
// preflightに対してstatus okを返す
// そうしないとPUTやDELETEのような非単純リクエストが実行されない
func AllowOptionsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func PostOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}

func GetOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}