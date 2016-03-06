package main

import (
	"encoding/json"
	//"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

//Advert struct for user ads
type Advert struct {
	ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Ad   string        `bson:"ad"`
	Name string        `bson:"name"`
}

//NewAd function for adding new adverts
func NewAd(r Advert) error {
	s, err := mgo.Dial(config.xx)

	defer s.Close()
	if err != nil {
		panic(err)
	}
	s.DB("yellowListings").C("Adverts").Insert(r)
	return err
}

//GetAds function for getting all adverts in db
func GetAds() ([]Advert, error) {
	result := []Advert{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Adverts")
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//NewAdHandler handler for adding new Adverts
func NewAdHandler(w http.ResponseWriter, r *http.Request) {
	var formdata Advert
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	NewAd(formdata)
}
