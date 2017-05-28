package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

func PaymentAfter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cust/newapp.html")
}

//HomeHandler for listing view
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.Category{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Categories []models.Category
	}{
		Categories: categories,
	}
	tmp := GetTemplates().Lookup("index.html")
	tmp.Execute(w, data)
}

//CategoryHandler for view
func SingleListingHandler(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)
	slug := params.ByName("listing")
	log.Printf("params %+v", params)

	listing, err := models.Listing{}.GetOne(config.Get(), slug)
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Listing models.Listing
	}{
		Listing: listing,
	}
	tmp := GetTemplates().Lookup("single.html")
	log.Printf("%+v", tmp)
	tmp.Execute(w, data)
}

//HomeHandler for listing view
func RegisterListing(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Plus bool
	}{
		Plus: false,
	}
	tmp := GetTemplates().Lookup("register_business.html")
	tmp.Execute(w, data)
}

func PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	tmp := GetTemplates().Lookup("privacy_policy.html")
	tmp.Execute(w, "")
}

//HomeHandler for listing view
func RegisterPlusListing(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Plus bool
	}{
		Plus: true,
	}
	tmp := GetTemplates().Lookup("register_business.html")
	tmp.Execute(w, data)
}
