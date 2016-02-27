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
	Image          string        `bson:"image"`
	Images         []string      `bson:"images"`
	Verified       string        `bson:"verified"`
	Approved       bool          `bson:"approved"`
	Plus           string        `bson:"plus"`
}

//Category struct for use in registration
type Category struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Category string        `bson:"category"`
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

//Addcat function for adding category
func Addcat(r Category) error {
	s, err := mgo.Dial(config.xx)

	r.ID = bson.NewObjectId()
	defer s.Close()
	if err != nil {
		log.Println(err)
	}
	s.DB("yellowListings").C("Category").Insert(r)
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
func Getunapproved() ([]Form, error) {
	result := []Form{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"approved": false}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getSingleList(r string) (Form, error) {
	result := Form{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(r)}).One(&result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

//GetListings return listings
func GetListings() ([]Form, error) {
	result := []Form{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"approved": true}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//Getcat list
func Getcat() ([]Category, error) {
	result := []Category{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Category")
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getSinglecat(r string) (Category, error) {
	result := Category{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Category")
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(r)}).One(&result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

//GetcatListing function
func GetcatListing(id string) ([]Form, error) {
	result := []Form{}
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return result, err
	}
	defer session.Close()

	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"category": id}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//AddHandler for adding Listings
func AddHandler(w http.ResponseWriter, r *http.Request) {
	var formdata Form
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	formdata.Approved = false
	Addlisting(formdata)

}

//GetHandler function
func GetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, _ := GetcatListing(id)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func getlistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, _ := getSingleList(id)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func getCatHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, _ := getSinglecat(id)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func addCatHandler(w http.ResponseWriter, r *http.Request) {
	var formdat Category
	err := json.NewDecoder(r.Body).Decode(&formdat)
	if err != nil {
		log.Println(err)
	}
	Addcat(formdat)
	data, _ := Getcat()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}

//Approvehandler to approve lsitings for view
func Approvehandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	UpdateListing(id)
	data, _ := Getunapproved()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

//GetunapprovedHandler gets unapproved listings
func GetunapprovedHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := Getunapproved()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

//GetListHandler for getting listings
func GetListHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetListings()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func getcatHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := Getcat()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
