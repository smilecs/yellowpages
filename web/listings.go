package web

import (
	"encoding/json"
	"strconv"

	//"github.com/gorilla/context"

	"log"
	"net/http"

	"github.com/tonyalaribe/yellowpages/config"
	"github.com/tonyalaribe/yellowpages/models"
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
	formdata.Add(config.Get())

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

func GetListing(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, err := models.Listing{}.GetOne(config.Get(), id)
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

//Approvehandler to approve lsitings for view
func Approvehandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	err := models.Listing{}.Approve(config.Get(), id)
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

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	var result Admin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	log.Println(user)
	collection := config.Get().Database.C("admin").With(config.Get().Database.Session.Copy())

	log.Println(user.Username)
	err = collection.Find(bson.M{"username": user.Username, "password": user.Password}).One(&result)
	if err != nil {
		log.Println(err)
	}
	data, _ := json.Marshal(result)
	w.Write(data)
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

/*
func GetPlusPayHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetPlusForPay(r.URL.Query().Get("id"))
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
*/
/*
func GetPlusPayHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetPlusForPay(r.URL.Query().Get("id"))
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

//GetListHandler for getting listings
func GetListHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetListings()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}


func Fictionalcat(w http.ResponseWriter, r *http.Request) {
	cat := Category{}
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	cat.Slug = "PlusListings"
	cat.Category = "PlusListings"
	s.DB(config.xy).C("Category").Insert(cat)

	cat.Category = "Sponsored"
	cat.Slug = "Sponsored"
	s.DB(config.xy).C("Category").Insert(cat)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var form Form
	var result Form
	s, err := mgo.Dial(config.xx)
	defer s.Close()
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Println(err)
	}

	err = s.DB(config.xy).C("Listings").Find(bson.M{"username": form.Username, "password": form.Password}).One(&result)
	if err != nil {
		log.Println(err)
	}
	data, _ := json.Marshal(result)
	w.Write(data)

}

func GetTr() error {
	session, err := mgo.Dial(config.xx)

	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB(config.xy).C("Listings")
	change := bson.M{"$set": bson.M{"category": "TRANSPORT-TRAVELS8"}}
	query := bson.M{"category": bson.RegEx{"TRAN.*", ""}}
	_, err = collection.UpdateAll(query, change)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Fix(w http.ResponseWriter, r *http.Request) {
	GetTr()
}*/
