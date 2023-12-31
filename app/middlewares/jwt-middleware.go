package middlewares

import (
	"be_golang/klp3/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ClaimsToken struct {
	Emails string `json:"emails"`
	ID     string `json:"id"`
	Role   string `json:"role"`
	Iat    int    `json:"iat"`
	Exp    int    `json:"exp"`
	Token  string `json:"token"`
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId string, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["userRole"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
}

func ExtractToken(e echo.Context) (string,string,string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)
		role:=claims["role"].(string)
		email:=claims["emails"].(string)
		return userId,role,email
	}
	return "","",""
}
