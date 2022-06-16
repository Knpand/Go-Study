package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	//これをimportしないとDB接続できない
	"app/modules"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	// "app/middlewar
)


type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Response struct{
	Status string `json:"status"`
	Data string `json:"data"`
}
func Login_handler(w http.ResponseWriter, r *http.Request) error{
	// Formデータを取得
	var user User
	var db_userpassword string
	json.NewDecoder(r.Body).Decode(&user)
	
	if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
		log.Print(err)
	}else{
	defer db.Close()
	// dbからデータ抜き出し
	err := db.QueryRow(`SELECT password FROM students WHERE email =?`,user.Email).Scan(&db_userpassword)
	if err != nil {
		log.Print(err)
		  w.WriteHeader(403)
		return nil
	}
	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(db_userpassword), []byte(user.Password)); err != nil {
		  w.WriteHeader(403)
		  return nil
	}else{
		fmt.Print("passcheck OK")
		response := Response {
			Status: "ok",
			Data: "success",
		  }
		
		json, _ := json.Marshal(response)
		log.Print(json)
		w.WriteHeader(200)
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		w.Write(json)
	}

	}
		return nil
}
func Signup_handler(w http.ResponseWriter, r *http.Request) error{
	// httpリクエストヘッダ確認
		// h := r.Header
		// fmt.Println(w, h)
		// h := r.Header.Get("Accept-Encoding")
		// fmt.Println(w, h)
	
	// Formデータを取得
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		hash_password:= modules.Hash_password(user.Password)
		
		if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
			log.Print(err)
		}else{
		// dbに登録
		stmt,err := db.Prepare(`INSERT INTO students(password, email) VALUES (?,?)`)
		if err != nil {
			log.Print(err)
			return nil
		}
	
		_, err = stmt.Exec(hash_password,user.Email)
	
		if err != nil {
			return nil
		}
		defer db.Close()
	
		}
			return nil
		
	}

	