package web

import (
	"encoding/json"
	"strconv"

	//"github.com/gorilla/context"

	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
	"gopkg.in/mgo.v2/bson"
)

type Admin struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	username string        `bson:"username"`
}

const ADVERT = "advert"

//AddHandler for adding Listings
func AddListing(w http.ResponseWriter, r *http.Request) {
	var formdata models.Listing
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	// log.Printf("%+v", formdata)
	formdata.Add(config.Get())

	data := struct {
		Message string
	}{"Successs adding listing"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//AddHandler for adding Listings
func EditListing(w http.ResponseWriter, r *http.Request) {
	var formdata models.Listing
	err := json.NewDecoder(r.Body).Decode(&formdata)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v", formdata)
	formdata.Edit(config.Get())

	data := struct {
		Message string
	}{"Successs adding listing"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//GetHandler function
func GetListings(w http.ResponseWriter, r *http.Request) {
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

func SingleListingHandlerJSON(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)
	slug := params.ByName("slug")

	data, err := models.Listing{}.GetOne(config.Get(), slug)
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

//ResultHandler for result view
func GetPlusListingsHandler(w http.ResponseWriter, r *http.Request) {

	listings, err := models.Listing{}.GetAllPlusListings(config.Get())
	if err != nil {
		log.Println(err)
	}

	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	k := 0
	for i := 0; i < len(listings.Data); i++ {
		if (i+1)%3 == 0 && k > len(ads) {
			post := Post{}
			post.Advert = ads[k]
			post.Type = config.ADVERT
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = config.LISTING
		post.Listing = listings.Data[i]
		posts = append(posts, post)

	}

	newURL := r.URL.Query()
	newURL.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(newURL)
	listings.Page.NextURL = r.URL.Path + "?" + newURL.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:       posts,
		Page:        listings.Page,
		PageHeading: "Pluslistings",
	}

	tmp := GetTemplates().Lookup("list.html")

	tmp.Execute(w, data)
}

//ResultHandler for result view
func GetPlusListingsJSON(w http.ResponseWriter, r *http.Request) {

	listings, err := models.Listing{}.GetAllPlusListings(config.Get())
	if err != nil {
		log.Println(err)
	}

	ads, err := models.Advert{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}

	posts := []Post{}
	k := 0
	for i := 0; i < len(listings.Data); i++ {
		if (i+1)%3 == 0 && k > len(ads) {
			post := Post{}
			post.Advert = ads[k]
			post.Type = config.ADVERT
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = config.LISTING
		post.Listing = listings.Data[i]
		posts = append(posts, post)

	}

	newURL := r.URL.Query()
	newURL.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(newURL)
	listings.Page.NextURL = r.URL.Path + "?" + newURL.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:       posts,
		Page:        listings.Page,
		PageHeading: "Pluslistings",
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

//Approvehandler to approve lsitings for view
func Approvehandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("q")
	err := models.Listing{}.Approve(config.Get(), slug)
	message := make(map[string]interface{})
	if err != nil {
		log.Println(err)
		message["message"] = err.Error()
		message["code"] = http.StatusConflict
		w.WriteHeader(http.StatusConflict)
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			log.Println(err)
		}
		return
	}
	message["message"] = "Approved Successfully"
	message["code"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println(err)
	}
}

func Deletehandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("q")
	err := models.Listing{}.Delete(config.Get(), slug)

	message := make(map[string]interface{})
	if err != nil {
		log.Println(err)
		message["message"] = err.Error()
		message["code"] = http.StatusConflict
		w.WriteHeader(http.StatusConflict)
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			log.Println(err)
		}
		return
	}
	message["message"] = "Deleted Successfully"
	message["code"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println(err)
	}
}

func TimeUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	expiry := r.URL.Query().Get("expiry")
	if expiry == "3000" {
		expiry = "1"
	} else {
		expiry = "2"
	}
	err := models.Listing{}.TimeUpdate(config.Get(), id, expiry)
	if err != nil {
		log.Println(err)
	}
	http.ServeFile(w, r, "cust/partials/success.html")
}

//GetunapprovedHandler gets unapproved listings
func Getunapproved(w http.ResponseWriter, r *http.Request) {
	data, err := models.Listing{}.GetAllUnapproved(config.Get())
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

//ResultHandler for result view
func SearchResultHandler(w http.ResponseWriter, r *http.Request) {
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
	listings, err := models.Listing{}.Search(config.Get(), query, pageInt)
	if err != nil {
		log.Println(err)
	}

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
			post.Type = config.ADVERT
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = config.LISTING
		post.Listing = listings.Data[i]
		posts = append(posts, post)
	}

	newURL := r.URL.Query()
	newURL.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(newURL)
	listings.Page.NextURL = r.URL.Path + "?" + newURL.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:          posts,
		Page:           listings.Page,
		PageHeading:    "Search",
		PageSubheading: query,
	}

	tmp := GetTemplates().Lookup("list.html")

	tmp.Execute(w, data)
}

//ResultHandler for result view
func SearchResultJSON(w http.ResponseWriter, r *http.Request) {
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
	listings, err := models.Listing{}.Search(config.Get(), query, pageInt)
	if err != nil {
		log.Println(err)
	}

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
			post.Type = config.ADVERT
			k++
			posts = append(posts, post)
		}
		post := Post{}
		post.Type = config.LISTING
		post.Listing = listings.Data[i]
		posts = append(posts, post)
	}

	newURL := r.URL.Query()
	newURL.Set("p", strconv.Itoa(listings.Page.NextVal))
	log.Println(newURL)
	listings.Page.NextURL = r.URL.Path + "?" + newURL.Encode()

	data := struct {
		Posts          []Post
		Page           models.Page
		PageHeading    string
		PageSubheading string
	}{
		Posts:          posts,
		Page:           listings.Page,
		PageHeading:    "Search",
		PageSubheading: query,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}
