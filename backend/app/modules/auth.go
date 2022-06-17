package modules


import (
	// "app/middlewar
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"net/http"
	"log"
)

type Claims struct {
	jwt.StandardClaims
}



func CreateToken(userID int)(string,error){
	// JWT
	claims := jwt.StandardClaims{
		Issuer: strconv.Itoa(int(userID)),            // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "" ,nil
	}

	return token,nil
	

}

func SetCookies(w http.ResponseWriter, r *http.Request,token string) {
	log.Print("start set jwt")

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		Secure:   true,
		Domain:   "localhost",
		SameSite: http.SameSiteNoneMode,
		// HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	log.Print("complete set jwt")
}

func ReSetCookies(w http.ResponseWriter, r *http.Request,token string) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	log.Print("set cookie")
	http.SetCookie(w, cookie)

}

func JwtCheck(w http.ResponseWriter, r *http.Request)(string ,error){
	// h := r.Header
	// log.Println(w, h)

	// CookieからJWTを取得
	cookie, err := r.Cookie("jwt")
	if err != nil{
		log.Print("cookie jwt check err")
		return "" ,err
	}else{
		// token取得
		token, _ := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		claims := token.Claims.(*Claims)
		id := claims.Issuer
		
		return id,nil
	}
}