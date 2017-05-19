package models

import (
	"log"

	"github.com/smilecs/yellowpages/config"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//Listing struct holds all submitted form data for listings
type AdminUser struct {
	Name         string
	Username     string
	Password     string `bson:"-"`
	PasswordHash []byte
	Email        string
	Type         string
	ID           string
}

//Addlisting function adding listings data to db
func (this AdminUser) Add(conf *config.Conf) error {
	log.Println(this)
	var err error
	this.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}

	mgoSession := conf.Database.Session.Copy()
	defer mgoSession.Close()

	collection := conf.Database.C(config.ADMINSCOLLECTION).With(mgoSession)
	_, err = collection.Upsert(bson.M{"username": this.Username}, this)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
