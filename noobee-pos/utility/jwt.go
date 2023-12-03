package utility

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Id      int       `json:"id"`
	Expires time.Time `json:"expires"`
}

var secretKey string
var expired int

func InitToken(secret string, expiredToken int) {
	secretKey = secret
	expired = expiredToken
}

func NewJWT(id int) JWT {
	return JWT{
		Id: id,

		Expires: time.Now().Add(time.Duration(expired) * time.Minute),
	}
}

func (j JWT) GenerateToken() (tokString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      j.Id,
		"expires": j.Expires,
	})

	tokString, err = token.SignedString([]byte(secretKey))
	return
}

func VerifyToken(tokString string) (token JWT, err error) {
	jwtToken, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		err = fmt.Errorf("invalid token")
		return
	}

	id := claims["id"]
	expires := fmt.Sprintf("%v", claims["expires"])
	expiresTime, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return
	}

	if time.Now().After(expiresTime) {
		err = fmt.Errorf("token expired")
		return
	}

	idInt, err := strconv.Atoi(fmt.Sprintf("%v", id))
	if err != nil {
		return
	}

	token = NewJWT(idInt)

	return
}
