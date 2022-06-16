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
	_ "github.com/go-sql-driver/mysql"	//これをimportしないとDB接続できない
	// "app/middleware"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
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
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
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

		fmt.Print("checkpoint")
		// w.WriteHeader(http.StatusOK)

		// Formデータを取得.
		form := r.PostForm
		
		fmt.Print(form)

		return nil
	}


func signup_handler(w http.ResponseWriter, r *http.Request) error{
// httpリクエストヘッダ確認
	// h := r.Header
	// fmt.Println(w, h)
	// h := r.Header.Get("Accept-Encoding")
	// fmt.Println(w, h)

// Formデータを取得
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Print("checkpoint1")

	hash:= Hash_password(user.Password)
	
	if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
		log.Print(err)
	}else{
	// dbに登録
	stmt,err := db.Prepare(`INSERT INTO students(password, email) VALUES (?,?)`)
	if err != nil {
		log.Print(err)
		return nil
	}

	_, err = stmt.Exec(hash,user.Email)

	if err != nil {
		return nil
	}
	defer db.Close()

	}
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