package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/tonyalaribe/yellowpages/config"
)

//Listing struct holds all submitted form data for listings
type User struct {
	Name  string
	Email string
	Type  string
	ID    string
	Link  string
}

//Addlisting function adding listings data to db
func (r User) Add(conf *config.Conf) error {
	collection := conf.Database.C(config.USERSCOLLECTION).With(conf.Database.Session.Copy())
	_, err := collection.Upsert(bson.M{"id": r.ID}, r)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
