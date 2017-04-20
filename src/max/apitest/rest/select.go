package rest

import (
	"fmt"
	"max/apitest/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Select(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := service.GetUser(vars["userId"])
	if err == nil {
		fmt.Fprintln(w, "User --> Nome: ", user.Name, " Cognome: ", user.Cognome)
	} else {
		fmt.Fprintln(w, err)
	}
}
