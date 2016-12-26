package config

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

//Conf nbfmjh
type Conf struct {
	MongoDB     string
	MongoServer string
	Database    *mgo.Database
}

var (
	config Conf
)

const (
	ADVERT  = "advert"
	LISTING = "listing"

	FACEBOOK = "facebook"
	GOOGLE   = "google"

	USERSCOLLECTION = "Users"
)

func Init() {
	MONGOSERVER := os.Getenv("MONGO_URL")
	MONGODB := os.Getenv("MONGODB")
	if MONGOSERVER == "" {
		log.Println("No mongo server address set, resulting to default address")
		MONGOSERVER = "127.0.0.1:27017"
		MONGODB = "calabarpages"
		//MONGODB = "yellowListings"
		//MONGODB = "y"
		//mongodb://localhost
	}

	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		log.Println(err)
	}

	config = Conf{
		MongoDB:     MONGODB,
		MongoServer: MONGOSERVER,
		Database:    session.DB(MONGODB),
	}

	log.Printf("mongoserver %s", MONGOSERVER)
}

func Get() *Conf {
	return &config
}
