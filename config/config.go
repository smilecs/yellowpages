package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve"
	bolt "github.com/coreos/bbolt"
	mgo "gopkg.in/mgo.v2"
)

//Conf nbfmjh
type Conf struct {
	MongoDB     string
	MongoServer string
	Database    *mgo.Database
	BoltFile    string
	BoltDB      *bolt.DB
	BleveFile   string
	BleveIndex  bleve.Index
	Encryption  struct {
		Private []byte
		Public  []byte
	}
}

var (
	config Conf
)

const (
	ADVERT  = "advert"
	LISTING = "listing"

	FACEBOOK = "facebook"
	GOOGLE   = "google"

	USERSCOLLECTION      = "Users"
	ADMINSCOLLECTION     = "Admins"
	LISTINGSCOLLECTION   = "Listings"
	CATEGORIESCOLLECTION = "Categories"
	ADVERTSCOLLECTION    = "Adverts"
	REVIEWSCOLLECTION    = "Adverts"
)

func CreateBleveIndex() (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	// listingsMapping := bleve.NewDocumentMapping()
	//
	// nameFieldMapping := bleve.NewTextFieldMapping()
	// nameFieldMapping.Analyzer = "en"
	// listingsMapping.AddFieldMappingsAt("companyname", nameFieldMapping)
	//
	// categoryFieldMapping := bleve.NewTextFieldMapping()
	// listingsMapping.AddFieldMappingsAt("category", categoryFieldMapping)
	//
	// descriptionFieldMapping := bleve.NewTextFieldMapping()
	// listingsMapping.AddFieldMappingsAt("about", descriptionFieldMapping)
	//
	// specialisationFieldMapping := bleve.NewTextFieldMapping()
	// listingsMapping.AddFieldMappingsAt("specialisation", specialisationFieldMapping)
	//
	// mapping.AddDocumentMapping("listing", listingsMapping)

	log.Println(config.BleveFile)
	bIndex, err := bleve.New(config.BleveFile, mapping)
	if err != nil {
		log.Println(err)
		return bIndex, err
	}

	return bIndex, nil
}

func Init() {
	// MONGOSERVER := os.Getenv("MONGO_URL")
	// MONGODB := os.Getenv("MONGODB")
	// if MONGOSERVER == "" {
	// 	log.Println("No mongo server address set, resulting to default address")
	// 	MONGOSERVER = "127.0.0.1:27017"
	// 	MONGODB = "yellowListings"
	// }
	//
	// session, err := mgo.Dial(MONGOSERVER)
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// log.Printf("mongoserver %s", MONGOSERVER)
	//
	config = Conf{
		// MongoDB:     MONGODB,
		// MongoServer: MONGOSERVER,
		// Database:    session.DB(MONGODB),
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "."
		log.Printf("db path not set, resulting to default address: %s\n", dbPath)
	}
	config.BoltFile = filepath.Join(dbPath, "yellowpages.bolt")

	log.Printf("bolt file: %s\n", config.BoltFile)
	db, err := bolt.Open(config.BoltFile, 0600, &bolt.Options{Timeout: 5 * time.Minute})
	if err != nil {
		log.Printf("create bleve index err: %v\n", err)
	}
	config.BoltDB = db

	config.BoltDB.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(USERSCOLLECTION))
		tx.CreateBucketIfNotExists([]byte(ADMINSCOLLECTION))
		tx.CreateBucketIfNotExists([]byte(LISTINGSCOLLECTION))
		tx.CreateBucketIfNotExists([]byte(CATEGORIESCOLLECTION))
		tx.CreateBucketIfNotExists([]byte(ADVERTSCOLLECTION))
		tx.CreateBucketIfNotExists([]byte(REVIEWSCOLLECTION))
		return err
	})

	// bleveFile := os.Getenv("BLEVE_PATH")
	// if bleveFile == "" {
	// 	log.Println("Blevefile not set, resulting to default address")
	// 	bleveFile = "./yellowpages.bleve"
	// }
	config.BleveFile = filepath.Join(dbPath, "yellowpages.bleve")

	log.Printf("bleve file: %s", config.BleveFile)
	bleveIndex, err := bleve.Open(config.BleveFile)
	if err != nil {
		log.Println("create bleve index")
		log.Println(err)
		bleveIndex, err = CreateBleveIndex()
	}
	config.BleveIndex = bleveIndex

	config.Encryption.Public, err = ioutil.ReadFile("./config/encryption_keys/public.pem")
	if err != nil {
		log.Println("Error reading public key")
		log.Println(err)
		return
	}

	config.Encryption.Private, err = ioutil.ReadFile("./config/encryption_keys/private.pem")
	if err != nil {
		log.Println("Error reading private key")
		log.Println(err)
		return
	}

}

func Get() *Conf {
	return &config
}
