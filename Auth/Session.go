package auth

import (
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func VerifyJWTAndGetClaim(req *http.Request) (claim *Claims, err error) {
	bearerToken := req.Header.Get("Authorization")
	bearer := strings.Split(bearerToken, " ")
	if len(bearer) != 2 {
		return nil, http.ErrContentLength
	}
	tokenstring := bearer[1]
	fmt.Println(tokenstring)
	claim = &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, err
	})
	if err != nil {
		return claim, err
	}
	if token.Valid {
		return claim, nil
	}
	return claim, jwt.ErrSignatureInvalid
}

func getJWT(user MyDb.User) (string, error) {
	claims := Claims{
		User_ID:    user.Id,
		Email:      user.Email,
		Role:       user.Role,
		Company_id: user.Company_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
