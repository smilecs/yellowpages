package web

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/smilecs/yellowpages/config"
)

func GetAnalytics(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	listingsCollection := config.Get().Database.C(config.LISTINGSCOLLECTION).With(config.Get().Database.Session.Copy())

	listingsCount, err := listingsCollection.Count()
	if err != nil {
		log.Println(err)
	}

	unapprovedListingsCount, err := listingsCollection.Find(bson.M{
		"approved": false,
	}).Count()

	usersCollection := config.Get().Database.C(config.USERSCOLLECTION).With(config.Get().Database.Session.Copy())

	usersCount, err := usersCollection.Count()
	if err != nil {
		log.Println(err)
	}

	data["UnapprovedListingsCount"] = unapprovedListingsCount
	data["ListingsCount"] = listingsCount
	data["UsersCount"] = usersCount

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}

}
