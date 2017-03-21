package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

func ReviewJSON(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	page := r.URL.Query().Get("p")
	pageInt := 1
	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			log.Println(err)
		}
	}

	listings, err := models.Reviews{}.GetAll(config.Get(), pageInt, query)
	if err != nil {
		log.Println(err)
	}
	//log.Println(listings)

	new_url := r.URL.Query()
	new_url.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(new_url)
	listings.Page.NextURL = r.URL.Path + "?" + new_url.Encode()

	data := struct {
		Posts          models.ReviewList
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:          listings,
		Page:           listings.Page,
		PageHeading:    "Reviews",
		PageSubheading: "",
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

func AddReviews(w http.ResponseWriter, r *http.Request) {
	var formdat models.Reviews
	err := json.NewDecoder(r.Body).Decode(&formdat)
	if err != nil {
		log.Println(err)
	}

	err = formdat.Add(config.Get())
	if err != nil {
		log.Println(err)
	}
	data, err := models.Reviews{}.GetAll(config.Get(), 1, formdat.Slug)
	if err != nil {
		log.Println(err)
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}
