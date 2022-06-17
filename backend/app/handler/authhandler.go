package handler
import (
	"encoding/json"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	//これをimportしないとDB接続できない
	"app/modules"
	"app/models"
	"golang.org/x/crypto/bcrypt"
	// "fmt"

	// "app/middlewar

)


func Login_handler(w http.ResponseWriter, r *http.Request) error{
	// Formデータを取得
	var user models.User
	user.Email=r.FormValue("email")
	user.Password=r.FormValue("password")
	var user_id int
	var db_userpassword string
	if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
		log.Print(err)
	}else{
	defer db.Close()
	// dbからデータ抜き出し
	err := db.QueryRow(`SELECT id,password FROM students WHERE email =?`,user.Email).Scan(&user_id,&db_userpassword)
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
		// JWT
		log.Print("create jwt")
		token,err:=modules.CreateToken(user_id)
		log.Print(user_id)
		if err !=nil{
			log.Print("create jwt error")
			w.WriteHeader(403)
			return nil
		}
		log.Print("createted jwt")
		response := models.Response {
			Status: "ok",
			Data: "success",
			Jwt:token,
			}
		//cookieの設定
		modules.SetCookies(w,r,token)
	
		json, _ := json.Marshal(response)
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		}	

	}
	return nil
}

func Signup_handler(w http.ResponseWriter, r *http.Request) error{
	// httpリクエストヘッダ確認
	// h := r.Header
	// log.Println(w, h)
	// h := r.Header.Get("Accept-Encoding")
	// fmt.Println(w, h)
	
	// Formデータを取得
	var user models.User
	// Content-typeがJsonの場合はこれ
	// json.NewDecoder(r.Body).Decode(&user)

	user.Name=r.FormValue("name")
	user.Email=r.FormValue("email")
	user.Password=r.FormValue("password")
	
	hash_password:= modules.Hash_password(user.Password)
	
	if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
		log.Print(err)
	}else{
	// dbに登録
	stmt,err := db.Prepare(`INSERT INTO students(password, email, name) VALUES (?,?,?)`)
	if err != nil {
		log.Print(err)
		return nil
	}
	_, err = stmt.Exec(hash_password,user.Email,user.Name)
	if err != nil {
		return nil
	}
	defer db.Close()
	}
		return nil
	}


func Logout_handler(w http.ResponseWriter, r *http.Request) error{
	token:=""
	//cookieの設定
	modules.ReSetCookies(w,r,token)
	response := models.Response {
		Status: "ok",
		Data: "log out",
		Jwt:token,
		}
	json, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.Write(json)

	return nil
}

func JwtCheck_handler(w http.ResponseWriter, r *http.Request) error{
	// CookieからJWTを参照し，Useridを取得
	user_id_jwt,err:=modules.JwtCheck(w,r)
	if err!= nil{
		w.WriteHeader(403)
		return nil
	}
	log.Print(user_id_jwt)
	var user models.User
	
	if db, err := sql.Open("mysql", "light:light@tcp(database:3306)/database"); err!=nil {
		log.Print(err)
	}else{
		defer db.Close()
		// dbからデータ抜き出し
		err := db.QueryRow(`SELECT name FROM students WHERE id =?`,user_id_jwt).Scan(&user.Name)
		if err != nil {
			log.Print(err)
			w.WriteHeader(400)
			return nil
		}else{
			response := models.Jwt_Authorized_Response{
				Status: "ok",
				Name: user.Name,
			  }
			
			json, _ := json.Marshal(response)
			w.WriteHeader(200)
			w.Header().Set("Content-Type","application/json; charset=utf-8")
			w.Write(json)
			
			return nil
		}

	}
		return nil
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