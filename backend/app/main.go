package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

var tasks = []Task{{
	ID:      1,
	Title:   "A",
	Content: "Aタスク",
	DueDate: time.Now(),
}, {
	ID:      2,
	Title:   "B",
	Content: "Bタスク",
	DueDate: time.Now(),
}, {
	ID:      3,
	Title:   "C",
	Content: "Cタスク",
	DueDate: time.Now(),
}}


func handler1(w http.ResponseWriter, r *http.Request) {
	//CORSの設定
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&tasks); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}


func main() {

	//http.ServeMux:URLにハンドラをバインド
	router := http.NewServeMux()

	//http.HandlerFunc
	// ・第１引数にTCPのアドレス
	// ・第２引数にhttp Handler
	//　ServeHTTPを実装
	router.Handle("/task1", http.HandlerFunc(handler1))

	//http.ServeMux.HandlerFunc
	//HandleFuncのショートカット的な
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())
	
		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	})


	
    //ListenAndServe:サーバの起動
	// ・第１引数にTCPのアドレス
	// ・第２引数にhttp Handler

	//log.Fatal:errorになるとexitする
	log.Fatal(http.ListenAndServe(":5050", router))	
}