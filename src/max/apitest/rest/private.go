package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MyKey = "test"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SetToken(res http.ResponseWriter, req *http.Request) {

	expireToken := time.Now().Add(time.Hour * 1).Unix()
	expireCookie := time.Now().Add(time.Hour * 1)

	claims := Claims{
		"admin",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    req.Host,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(res, &cookie)

	http.Redirect(res, req, "/private", 307)
}

func Validate(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			http.NotFound(res, req)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), MyKey, *claims)
			protectedPage(res, req.WithContext(ctx))
		} else {
			http.NotFound(res, req)
			return
		}
	})
}

func ProtectedProfile(res http.ResponseWriter, req *http.Request) {
	claims, ok := req.Context().Value(MyKey).(Claims)
	if !ok {
		http.NotFound(res, req)
		return
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<span>Hello "+claims.Username+"</span><br/><br/>")
	fmt.Fprintln(res, "<button onclick='window.location = "+"\"private/list\""+"' >lista utenti</button><br/><br/>"+
		"<button onclick='window.location = "+"\"private/select/1\""+"' >leggi il primo utente</button><br/><br/>")

	fmt.Fprintln(res, "<button onclick='window.location = "+"\"logout\""+"'>logout</button>")
}

func Logout(res http.ResponseWriter, req *http.Request) {
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(res, &deleteCookie)
	http.Redirect(res, req, "/", 307)
	return
}
