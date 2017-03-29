package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/blevesearch/bleve"
	"github.com/tonyalaribe/yellowpages/config"
)

func IndexSingleListingWithBleve(listing Listing) error {

	bleveIndex := config.Get().BleveIndex

	// index some data
	err := bleveIndex.Index(listing.Slug, listing)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("indexed %s successfully", listing.Slug)
	return nil
}

func IndexMongoDBListingsCollectionWithBleve() {
	conf := config.Get()
	mgoSession := conf.Database.Session.Copy()
	defer mgoSession.Close()
	collection := conf.Database.C(config.LISTINGSCOLLECTION).With(mgoSession)

	bleveIndex := config.Get().BleveIndex

	q := collection.Find(bson.M{"plus": "true"})
	count, err := q.Count()
	if err != nil {
		log.Println(err)
	}
	cycles := count / 100
	// log.Printf("cycles: %+v", cycles)
	for i := 0; i <= cycles; i++ {
		// log.Printf("cycle: %d", i)
		listings := []Listing{}
		err = q.Limit(100).Skip(i * 100).All(&listings)
		if err != nil {
			log.Println(err)
		}

		for _, listing := range listings {
			err = bleveIndex.Index(listing.Slug, listing)
			if err != nil {
				log.Println(err)
				//return err
			}
			log.Printf("indexed %s successfully", listing.Slug)
		}
	}

}

func SearchWithIndex(queryString string) (*bleve.SearchResult, error) {
	query := bleve.NewFuzzyQuery(queryString)
	query.Fuzziness = 2
	search := bleve.NewSearchRequest(query)
	bleveIndex := config.Get().BleveIndex

	searchResults, err := bleveIndex.Search(search)
	if err != nil {
		log.Println(err)
		return searchResults, err
	}

	return searchResults, nil
}
