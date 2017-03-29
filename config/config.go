package config

import (
	"log"
	"os"

	"github.com/blevesearch/bleve"
	mgo "gopkg.in/mgo.v2"
)

//Conf nbfmjh
type Conf struct {
	MongoDB     string
	MongoServer string
	Database    *mgo.Database
	BleveFile   string
	BleveIndex  bleve.Index
}

var (
	config Conf
)

const (
	ADVERT  = "advert"
	LISTING = "listing"

	FACEBOOK = "facebook"
	GOOGLE   = "google"

	USERSCOLLECTION    = "Users"
	LISTINGSCOLLECTION = "Listings"
)

func CreateBleveIndex() (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	listingsMapping := bleve.NewDocumentMapping()

	nameFieldMapping := bleve.NewTextFieldMapping()
	nameFieldMapping.Analyzer = "en"
	listingsMapping.AddFieldMappingsAt("companyname", nameFieldMapping)

	descriptionFieldMapping := bleve.NewTextFieldMapping()
	descriptionFieldMapping.Analyzer = "en"
	listingsMapping.AddFieldMappingsAt("about", descriptionFieldMapping)

	specialisationFieldMapping := bleve.NewTextFieldMapping()
	listingsMapping.AddFieldMappingsAt("about", specialisationFieldMapping)

	mapping.AddDocumentMapping("listing", listingsMapping)

	bIndex, err := bleve.New(config.BleveFile, mapping)
	if err != nil {
		log.Println(err)
		return bIndex, err
	}

	return bIndex, nil
}

func Init() {
	MONGOSERVER := os.Getenv("MONGO_URL")
	MONGODB := os.Getenv("MONGODB")
	if MONGOSERVER == "" {
		log.Println("No mongo server address set, resulting to default address")
		MONGOSERVER = "127.0.0.1:27017"
		//MONGODB = "calabarpages"
		MONGODB = "yellowListings"
		//MONGODB = "y"
		//mongodb://localhost
	}

	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		log.Println(err)
	}

	log.Printf("mongoserver %s", MONGOSERVER)

	bleveFile := os.Getenv("BLEVE_PATH")
	if bleveFile == "" {
		log.Println("Blevefile not set, resulting to default address")
		bleveFile = "./yellowpages.bleve"
	}

	bleveIndex, err := bleve.Open(bleveFile)
	if err != nil {
		log.Println("create bleve index")
		log.Println(err)
		bleveIndex, err = CreateBleveIndex()
	}

	config = Conf{
		MongoDB:     MONGODB,
		MongoServer: MONGOSERVER,
		Database:    session.DB(MONGODB),
		BleveFile:   bleveFile,
		BleveIndex:  bleveIndex,
	}

}

func Get() *Conf {
	return &config
}
