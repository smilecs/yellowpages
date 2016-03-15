package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//View struct holds data of both adverts and listing view
type View struct {
	CompanyName    string `bson:"companyname"`
	Address        string `bson:"address"`
	Hotline        string `bson:"hotline"`
	Specialisation string `bson:"specialisation"`
	Category       string `bson:"category"`
	Type           string
	Image          string
}

//RenderView function to return listings mixed with adds
func RenderView(id string) ([]View, error) {
	tmp := []Form{}
	result := []View{}
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return result, err
	}
	defer session.Close()
	adCollection := session.DB("yellowListings").C("Listings")
	err = adCollection.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"category": id}).All(&tmp)
	if err != nil {
		return result, err
	}

	return result, nil
}
