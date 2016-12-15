package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tonyalaribe/yellowpages/config"
	"github.com/tonyalaribe/yellowpages/models"
)

func GetFromCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	page := r.URL.Query().Get("p")
	pageInt := 1
	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			log.Println(err)
		}
	}
	data, err := models.Listing{}.GetAllInCategory(config.Get(), id, pageInt)
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

//CategoryHandler for view
func CategoryListingsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)
	category := params.ByName("category")

	page := r.URL.Query().Get("p")
	pageInt := 1
	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			log.Println(err)
		}
	}

	listings, err := models.Listing{}.GetAllInCategory(config.Get(), category, pageInt)
	if err != nil {
		log.Println(err)
	}
	//log.Println(listings)
	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	k := 0
	for i := 0; i < len(listings.Data); i++ {
		if (i+1)%3 == 0 && k < len(ads) {
			post := Post{}
			post.Advert = ads[k]
			post.Type = "advert"
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = "listing"
		post.Listing = listings.Data[i]
		posts = append(posts, post)

	}

	new_url := r.URL.Query()
	new_url.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(new_url)
	listings.Page.NextURL = r.URL.Path + "?" + new_url.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:          posts,
		Page:           listings.Page,
		PageHeading:    "Category",
		PageSubheading: category,
	}
	tmp := GetTemplates().Lookup("list.html")
	tmp.Execute(w, data)
}

//CategoryHandler for view
func CategoryListingsJSON(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)
	category := params.ByName("category")

	page := r.URL.Query().Get("p")
	pageInt := 1
	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			log.Println(err)
		}
	}

	listings, err := models.Listing{}.GetAllInCategory(config.Get(), category, pageInt)
	if err != nil {
		log.Println(err)
	}
	//log.Println(listings)
	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	k := 0
	for i := 0; i < len(listings.Data); i++ {
		if (i+1)%3 == 0 && k < len(ads) {
			post := Post{}
			post.Advert = ads[k]
			post.Type = "advert"
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = "listing"
		post.Listing = listings.Data[i]
		posts = append(posts, post)

	}

	new_url := r.URL.Query()
	new_url.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(new_url)
	listings.Page.NextURL = r.URL.Path + "?" + new_url.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:          posts,
		Page:           listings.Page,
		PageHeading:    "Category",
		PageSubheading: category,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var formdat models.Category
	err := json.NewDecoder(r.Body).Decode(&formdat)
	if err != nil {
		log.Println(err)
	}

	err = formdat.Add(config.Get())
	if err != nil {
		log.Println(err)
	}
	data, err := models.Category{}.GetAll(config.Get())
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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	data, err := models.Category{}.GetAll(config.Get())
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
