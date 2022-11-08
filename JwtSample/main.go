package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ErrorServerBusy = "server is busy"
	ErrorRelogin    = "relogin"
)

type JwtClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"useId"`
	Password string `json:"password"`
	Username string `json:"username"`
}

var (
	Secret     = "test" // salt
	ExpireTime = 3600   //token expire time
)

func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.GET("/verify/:token", verify)
	r.GET("/refresh/:token", refresh)
	r.GET("/hello/:token", sayHello)
	r.Run(":8080")
}

func genToken(claims *JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	return signedToken, nil
}

func login(ctx *gin.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")
	claims := &JwtClaims{
		UserId:   1,
		Username: username,
		Password: password,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := genToken(claims)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func verifyAction(strToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New(ErrorRelogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorRelogin)
	}

	fmt.Println("verify")
	return claims, nil
}

func sayHello(ctx *gin.Context) {
	strToken := ctx.Param("token")
	cliam, err := verifyAction(strToken)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprint("hello,", cliam.Username)})
}

func verify(ctx *gin.Context) {
	strToken := ctx.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "verify: " + claim.Username})
}

func refresh(ctx *gin.Context) {
	strToken := ctx.Param("token")
	claims, err := verifyAction(strToken)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := genToken(claims)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": signedToken, "expiresAt": claims.ExpiresAt})
}
