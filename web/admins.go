package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/smilecs/yellowpages/config"

	"gopkg.in/mgo.v2/bson"
)

func GetAdmins() ([]User, error) {
	result := []User{}

	collection := config.Get().Database.C("admin").With(config.Get().Database.Session.Copy())
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAdminsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := GetAdmins()
	if err != nil {
		log.Println(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	collection := config.Get().Database.C("admin").With(config.Get().Database.Session.Copy())

	err = collection.Insert(user)
	if err != nil {
		log.Println(err)
	}
}
