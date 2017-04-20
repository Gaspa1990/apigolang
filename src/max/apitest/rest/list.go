package rest

import (
	"encoding/json"
	"max/apitest/service"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	result := service.ListUser()
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
