package models

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"

	bolt "github.com/coreos/bbolt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/satori/go.uuid"
	"github.com/smilecs/yellowpages/config"
)

//Advert struct for user ads
type Advert struct {
	ID    int    `json:"id,omitempty" bson:"_id,omitempty"`
	Image string `bson:"image"`
	Name  string `bson:"name"`
	Type  string `bson:"type"`
}

//NewAd function for adding new adverts
func (r Advert) New(conf *config.Conf) error {

	r.Type = "advert"

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("yellowpagesng")

	byt, err := base64.StdEncoding.DecodeString(strings.Split(r.Image, "base64,")[1])
	if err != nil {
		log.Println(err)
	}

	meta := strings.Split(r.Image, "base64,")[0]
	newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
	imagename := "adverts/" + uuid.NewV1().String()

	err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite)
	if err != nil {
		log.Println(err)
	}

	log.Println(bucket.URL(imagename))

	r.Image = bucket.URL(imagename)

	rJSONbyt, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.ADVERTSCOLLECTION))
		id, _ := b.NextSequence()
		r.ID = int(id)
		err := b.Put(itob(r.ID), rJSONbyt)
		return err
	})

	return nil
}

//Get function for getting all adverts in db
func (Advert) GetAll(conf *config.Conf) ([]Advert, error) {

	results := []Advert{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		b.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			ad := Advert{}
			json.Unmarshal(v, &ad)
			results = append(results, ad)

			return nil
		})
		return nil
	})

	return results, err
}
