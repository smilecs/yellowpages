package models

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/satori/go.uuid"
	"github.com/tonyalaribe/yellowpages/config"
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
func (r Advert) New(config *config.Conf) error {

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
	imagename := uuid.NewV1().String()

	err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite)
	if err != nil {
		log.Println(err)
	}

	log.Println(bucket.URL(imagename))

	r.Image = bucket.URL(imagename)
	collection := config.Database.C("Adverts").With(config.Database.Session.Copy())
	err = collection.Insert(r)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Get function for getting all adverts in db
func (Advert) GetAll(config *config.Conf) ([]Advert, error) {
	result := []Advert{}
	collection := config.Database.C("Adverts").With(config.Database.Session.Copy())
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
