package util
import (
	 "time"

	jwt "github.com/dgrijalva/jwt-go"

	"examination/pkg/setting"
)
//jwt工具包
var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Level int `json:"level"`
	jwt.StandardClaims
}

func GenerateToken(username string,Level int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		Level,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //hs256加密法
	tokenstring, err := token.SignedString(jwtSecret)   //这两句获取token

	return tokenstring, err
}

func ParseToken(tokenstring string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil //jwtsecret就是key,解析tokenstring生成token
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}