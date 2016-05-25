package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	//"github.com/gorilla/context"

	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	Slug           string        `bson:"slug"`
	About          string        `bson:"about"`
	Username       string        `bson:"username"`
	Password       string        `bson:"password"`
	RC             string        `bson:"rc"`
	Branch         string        `bson:"branch"`
	Product        string        `bson:"product"`
	Email          string        `bson:"email"`
	Website        string        `bson:"website"`
	DHr            string        `bson:"dhr"`
	Verified       string        `bson:"verified"`
	Approved       bool          `bson:"approved"`
	Plus           string        `bson:"plus"`
	Pg             Page
}

//Category struct for use in registration
type Category struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Category string        `bson:"category"`
	Slug     string        `bson:"slug"`
	Show     string        `bson:"show"`
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//Addlisting function adding listings data to db
func Addlisting(r Form) error {
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		panic(err)
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("yellowpagesng")

	p := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := strconv.Itoa(p.Intn(10))

	byt, err := base64.StdEncoding.DecodeString(strings.Split(r.Image, "base64,")[1])
	if err != nil {
		log.Println(err)
	}

	meta := strings.Split(r.Image, "base64,")[0]
	newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
	imagename := randSeq(30)

	err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite)
	if err != nil {
		log.Println(err)
	}

	log.Println(bucket.URL(imagename))

	r.Image = bucket.URL(imagename)

	var images []string
	for _, v := range r.Images {

		byt, err := base64.StdEncoding.DecodeString(strings.Split(v, "base64,")[1])
		if err != nil {
			log.Println(err)
		}

		meta := strings.Split(v, "base64,")[0]

		newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)

		imagename := randSeq(10)

		err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite)
		if err != nil {
			log.Println(err)
		}
		images = append(images, bucket.URL(imagename))
	}

	r.Images = images
	r.Slug = strings.Replace(r.CompanyName, " ", "-", -1) + str
	index := mgo.Index{
		Key:        []string{"$text:specialisation", "$text:companyname"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	collection := s.DB(config.xy).C("Listings")
	collection.EnsureIndex(index)
	collection.Insert(r)
	return err
}

//Addcat function for adding category
func Addcat(r Category) error {
	s, err := mgo.Dial(config.xx)

	defer s.Close()
	if err != nil {
		log.Println(err)
	}
	p := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := strconv.Itoa(p.Intn(10))
	r.Slug = strings.Replace(r.Category, " ", "-", -1) + str
	r.Show = "true"
	s.DB(config.xy).C("Category").Insert(r)
	return err
}

//UpdateListing to approve of a listing to show in the client side
func UpdateListing(id string) error {
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return err
	}

	defer session.Close()

	collection := session.DB(config.xy).C("Listings")
	query := bson.M{"slug": id}
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

	collection := session.DB(config.xy).C("Listings")
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

	collection := session.DB(config.xy).C("Listings")
	err = collection.Find(bson.M{"slug": r}).One(&result)
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

	collection := session.DB(config.xy).C("Listings")
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

	collection := session.DB(config.xy).C("Category")
	err = collection.Find(bson.M{"show": "true"}).All(&result)
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

	collection := session.DB(config.xy).C("Category")
	err = collection.Find(bson.M{"slug": r}).One(&result)
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

	collection := session.DB(config.xy).C("Listings")
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

func Fictionalcat(w http.ResponseWriter, r *http.Request) {
	cat := Category{}
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	cat.Slug = "PlusListings"
	cat.Category = "PlusListings"
	s.DB(config.xy).C("Category").Insert(cat)

	cat.Category = "Sponsored"
	cat.Slug = "Sponsored"
	s.DB(config.xy).C("Category").Insert(cat)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var form Form
	var result Form
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Println(err)
	}

	err = s.DB(config.xy).C("Listings").Find(bson.M{"username": form.Username, "password": form.Password}).One(&result)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte("logged"))

}
