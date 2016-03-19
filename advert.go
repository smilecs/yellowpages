package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	//"github.com/gorilla/context"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Advert struct for user ads
type Advert struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Image string        `bson:"image"`
	Name  string        `bson:"name"`
	Type  string        `bson:"type"`
}

//NewAd function for adding new adverts
func NewAd(r Advert) error {
	s, err := mgo.Dial(config.xx)

	defer s.Close()
	if err != nil {
		panic(err)
	}

	r.Type = "advert"

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("yellowpagesng")

	//p := rand.New(rand.NewSource(time.Now().UnixNano()))
	//str := strconv.Itoa(p.Intn(10))

	byt, err := base64.StdEncoding.DecodeString(strings.Split(r.Image, "base64,")[1])
	if err != nil {
		log.Println(err)
	}

	meta := strings.Split(r.Image, "base64,")[0]
	newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
	imagename := randSeq(10)

	err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite)
	if err != nil {
		log.Println(err)
	}

	log.Println(bucket.URL(imagename))

	r.Image = bucket.URL(imagename)

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
