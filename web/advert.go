package web

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

//NewAdHandler handler for adding new Adverts
func NewAdHandler(w http.ResponseWriter, r *http.Request) {
	var formdata models.Advert
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	formdata.New(config.Get())
}

//ResultHandler for result view
func GetAdvertsHandler(w http.ResponseWriter, r *http.Request) {

	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	for i := 0; i < len(ads); i++ {
		post := Post{}
		post.Advert = ads[i]
		post.Type = "advert"
		posts = append(posts, post)

	}

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:       posts,
		Page:        models.Page{},
		PageHeading: "Adverts",
	}

	tmp := GetTemplates().Lookup("list.html")

	tmp.Execute(w, data)
}

//ResultHandler for result view
func GetAdvertsJSON(w http.ResponseWriter, r *http.Request) {

	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	for i := 0; i < len(ads); i++ {
		post := Post{}
		post.Advert = ads[i]
		post.Type = "advert"
		posts = append(posts, post)

	}

	data := struct {
		Posts          []Post
		PageHeading    string
		PageSubheading string
	}{
		Posts:       posts,
		PageHeading: "Adverts",
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}
