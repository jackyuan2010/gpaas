package gin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId   uuid.UUID
	UserName string
}

var (
	secret     = []byte("123456789")
	noVerify   = []interface{}{"/login", "/ping"}
	effectTime = 2 * time.Hour
)

func GernerateToken(claims *JWTClaims) string {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		fmt.Println(err)
	}
	return sign
}

func JwtVerify(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		panic("token not exist!")
	}
	c.Set("user", ParseToken(token))
}

func ParseToken(tokenString string) *JWTClaims {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		fmt.Println(err)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}

func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	claims := ParseToken(tokenString)

	jwt.TimeFunc = time.Now

	claims.StandardClaims.ExpiresAt = time.Now().Add(effectTime).Unix()
	return GernerateToken(claims)
}
