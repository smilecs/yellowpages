package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/blevesearch/bleve"
	"github.com/smilecs/yellowpages/config"
)

func IndexSingleListingWithBleve(listing Listing) error {

	bleveIndex := config.Get().BleveIndex

	plus := false
	if listing.Plus == "true" {
		plus = true
	}
	indexableListing := struct {
		CompanyName    string
		Specialisation string
		Category       string
		Plus           bool
	}{
		CompanyName:    listing.CompanyName,
		Specialisation: listing.Specialisation,
		Category:       listing.Category,
		Plus:           plus,
	}
	// index some data
	err := bleveIndex.Index(listing.Slug, indexableListing)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("indexed %s successfully", listing.Slug)
	return nil
}

func DeleteSingleListingInBleve(slug string) error {
	bleveIndex := config.Get().BleveIndex

	// index some data
	err := bleveIndex.Delete(slug)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("deleted %s successfully", slug)
	return nil
}

func IndexMongoDBListingsCollectionWithBleve() {
	conf := config.Get()
	mgoSession := conf.Database.Session.Copy()
	defer mgoSession.Close()
	collection := conf.Database.C(config.LISTINGSCOLLECTION).With(mgoSession)

	// bleveIndex := config.Get().BleveIndex

	q := collection.Find(bson.M{"approved": true})
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
			//
			// plus := false
			// if listing.Plus == "true" {
			// 	plus = true
			// }
			// indexableListing := struct {
			// 	CompanyName    string
			// 	Specialisation string
			// 	Category       string
			// 	Plus           bool
			// }{
			// 	CompanyName:    listing.CompanyName,
			// 	Specialisation: listing.Specialisation,
			// 	Category:       listing.Category,
			// 	Plus:           plus,
			// }
			// log.Println(listing)
			// log.Println(listing.Slug)

			pListing := listing
			err = pListing.Add(conf)
			if err != nil {
				log.Println(err)
				//return err
			}
			// err = bleveIndex.Index(pListing.Slug, indexableListing)
			// if err != nil {
			// 	log.Println(err)
			// 	//return err
			// }
			log.Printf("indexed %s successfully", pListing.Slug)
		}
	}

}

func IndexData() {
	IndexCategories()
	IndexMongoDBListingsCollectionWithBleve()

}
func IndexCategories() {
	conf := config.Get()
	mgoSession := conf.Database.Session.Copy()
	defer mgoSession.Close()
	collection := conf.Database.C("Category").With(mgoSession)

	// bleveIndex := config.Get().BleveIndex

	q := collection.Find(bson.M{})
	count, err := q.Count()
	if err != nil {
		log.Println(err)
	}
	log.Printf("count: %d", count)
	cycles := count / 100
	// log.Printf("cycles: %+v", cycles)
	for i := 0; i <= cycles; i++ {
		// log.Printf("cycle: %d", i)
		items := []Category{}
		err = q.Limit(100).Skip(i * 100).All(&items)
		if err != nil {
			log.Println(err)
		}

		for _, item := range items {
			item.Add(conf)
			log.Printf("indexed %s successfully", item.Slug)
		}
	}

}

func SearchWithIndex(queryString string, skip int) (*bleve.SearchResult, error) {
	log.Println(skip)
	query := bleve.NewFuzzyQuery(queryString)
	query.Fuzziness = 3
	search := bleve.NewSearchRequest(query)
	search.From = skip
	search.Size = 10
	search.SortBy([]string{"-Plus", "-CompanyName", "-_score"})
	// search.SortBy([]string{"-plus", "-_score", "_id"})
	bleveIndex := config.Get().BleveIndex
	log.Println(bleveIndex.Mapping().DefaultSearchField())
	searchResults, err := bleveIndex.Search(search)
	log.Println(searchResults.String())
	if err != nil {
		log.Println(err)
		return searchResults, err
	}

	return searchResults, nil
}

func SearchField(queryString string, field string, skip int) (*bleve.SearchResult, error) {
	log.Println(skip)

	query := bleve.NewMatchQuery(queryString)
	// query.SetField(field)
	log.Println(queryString)
	search := bleve.NewSearchRequest(query)
	search.From = skip
	search.Size = 10
	search.SortBy([]string{"-Plus", "-Category", "-_score"})
	// search.SortBy([]string{"-plus", "-_score", "_id"})
	bleveIndex := config.Get().BleveIndex
	// log.Println(bleveIndex.Mapping().AnalyzerNamed("Category"))
	searchResults, err := bleveIndex.Search(search)
	log.Println(searchResults.String())
	if err != nil {
		log.Println(err)
		return searchResults, err
	}

	return searchResults, nil
}
