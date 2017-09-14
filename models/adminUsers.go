package models

import (
	"encoding/json"
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/smilecs/yellowpages/config"
	"golang.org/x/crypto/bcrypt"
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

	old := AdminUser{}
	jsonByt := []byte{}

	// var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.ADMINSCOLLECTION))
		jsonByt = b.Get([]byte(this.Username))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &old)
	if err != nil {
		log.Println(err)
	}

	old.Username = this.Username
	old.PasswordHash = this.PasswordHash
	old.Name = this.Name
	old.Email = this.Email
	old.Type = this.Type

	rJSONbyt, err := json.Marshal(old)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(rJSONbyt))
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.ADMINSCOLLECTION))

		err := b.Put([]byte(old.Username), rJSONbyt)
		return err
	})

	return nil
}

//Addlisting function adding listings data to db
func (this AdminUser) Get(conf *config.Conf) (AdminUser, error) {

	admin := AdminUser{}
	jsonByt := []byte{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.ADMINSCOLLECTION))
		jsonByt = b.Get([]byte(this.Username))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &admin)
	if err != nil {
		log.Println(err)
	}

	return admin, nil
}
