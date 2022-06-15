package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"app/handler"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	// "app/middleware"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

type User struct {
	Token string `json:"token"`
	UserID int `json:"id"`
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	fmt.Print("checkpoint")
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


func login_handler(w http.ResponseWriter, r *http.Request) error{
	fmt.Print("checkpoint")
	user:=User{Token:"1234567890abcdef",UserID:1}
	jsonData,err:=json.Marshal(user)
	if err != nil {
		fmt.Print("checkpoint")
		return err
	}else{
		fmt.Print("checkpoint")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

		// Formデータを取得.
		form := r.PostForm
		
		fmt.Print(form)

		return nil
	}
}

func signup_handler(w http.ResponseWriter, r *http.Request) error{
	fmt.Print("checkpoint")

	db, err := sql.Open("mysql", "light:light@tcp(localhost:3306)/database")
	// Formデータを取得.
	email :=  r.FormValue("email")
	password :=  r.FormValue("password")

	hash:= Hash_password(password)
	fmt.Print("checkpoint")
	// dbに登録
	_, err = db.Exec(
		"INSERT INTO students(password, email) VALUES (?,?)",
		hash,
		email,
	)

	if err != nil {
		user:=User{Token:"1234567890abcdef",UserID:1}
		jsonData,_:=json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		defer db.Close()
		return nil
	}
	defer db.Close()

		return nil
	
}

func Hash_password(password string) string {
	new_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(new_password)
}
func main() {
	//handler.Handle 引数がnilじゃいけない理由
	//handler.Handleでmidllewareとhandler回してる

	http.HandleFunc("/", handler.Handle(login_handler))
	http.HandleFunc("/signup", handler.Handle(signup_handler))

	log.Fatal(http.ListenAndServe(":5050", nil))

}




/////////////////////////////////////
//////
//////GETとPOST書き分けるときに使える
//////
/////////////////////////////////////


// func handleOnlyPost(w http.ResponseWriter, r *http.Request) {

//     // HTTPメソッドをチェック（POSTのみ許可）
//     if r.Method != http.MethodPost {
//         w.WriteHeader(http.StatusMethodNotAllowed) // 405
//         w.Write([]byte("POSTだけだよー"))
//         return
//     }

//     w.Write([]byte("OK"))
// }