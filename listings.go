package main

import (
	"encoding/json"
	//"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

//Form struct holds all submitted form data for listings
type Form struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyName    string        `bson:"companyname"`
	Address        string        `bson:"address"`
	Hotline        string        `bson:"hotline"`
	Specialisation string        `bson:"specialisation"`
	Category       string        `bson:"category"`
	Advert         string        `bson:"advert"`
	Size           string        `bson:"size"`
	Verified       string        `bson:"verified"`
	Approved       bool          `bson:"approved"`
}

//Addlisting function adding listings data to db
func Addlisting(r Form) error {
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		panic(err)
	}
	s.DB("yellowListings").C("Listings").Insert(r)
	return err
}

//UpdateListing to approve of a listing to show in the client side
func UpdateListing(id string) error {
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return err
	}

	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": bson.M{"approved": true}}

	err = collection.Update(query, change)

	if err != nil {
		return err
	}

	return nil
}

//Getunapproved function for Getunapproved handler
func Getunapproved() (Form, error) {
	result := Form{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"approved": false}).One(&result)
	log.Println(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//AddHandler for adding Listings
func AddHandler(w http.ResponseWriter, r *http.Request) {
	//formdata := context.Get(r, "body").(*Form)

	var formdata Form
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	formdata.Approved = false
	log.Println(formdata)
	Addlisting(formdata)

}

//Approvehandler to approve lsitings for view
func Approvehandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	log.Println(id)
	UpdateListing(id)
}

//GetunapprovedHandler gets unapproved listings
func GetunapprovedHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := Getunapproved()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
