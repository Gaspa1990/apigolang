package rest

import (
	"fmt"
	"net/http"
)

func Dummy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Welcome!</h1>"+
		"<button onclick='window.location = "+"\"settoken\""+"' >test autentica</button>")
}
