package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tonyalaribe/yellowpages/config"
	"github.com/tonyalaribe/yellowpages/models"
)

func FBLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	user.Type = config.FACEBOOK

	err = user.Add(config.Get())
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Message string
	}{"login success"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
