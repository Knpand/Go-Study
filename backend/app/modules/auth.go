package modules


import (
	// "app/middlewar
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

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