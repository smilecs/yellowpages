package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//View struct holds data of both adverts and listing view
type View struct {
	ID             bson.ObjectId `bson:"id"`
	CompanyName    string        `bson:"companyname"`
	Address        string        `bson:"address"`
	Hotline        string        `bson:"hotline"`
	Specialisation string        `bson:"specialisation"`
	Category       string        `bson:"category"`
	Type           string
	Image          string
}
type Result struct {
	Data  []Form
	Pag   Page
	Query string
}

//RenderView function to return listings mixed with adds
func RenderView(id string) ([]*View, error) {
	tmp := []Form{}
	result := []*View{}
	//res := []Advert{}
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return result, err
	}
	defer session.Close()
	/*adCollection := session.DB("yellowListings").C("Adverts")
	err = adCollection.Find(bson.M{}).All(&res)
	if err != nil {
		return result, err
	}*/
	res, _ := GetAds()
	collection := session.DB("yellowListings").C("Listings")
	err = collection.Find(bson.M{"category": id}).All(&tmp)
	if err != nil {
		return result, err
	}
	ik := len(res)
	for i := 0; i < len(tmp); i++ {
		rs := tmp[i]
		if i%2 > 0 {
			if i < ik {
				views := new(View)
				rss := res[i]
				views.Image = rss.Image
				views.ID = rss.ID
				views.Type = rss.Type
				views.CompanyName = rss.Name
				result = append(result, views)
			}

		}
		view := new(View)
		view.Hotline = rs.Hotline
		view.ID = rs.ID
		view.Category = rs.Category
		view.Address = rs.Address
		view.CompanyName = rs.CompanyName
		view.Image = rs.Image
		view.Specialisation = rs.Specialisation
		view.Type = rs.Plus
		result = append(result, view)

	}

	return result, nil
}

func Search(query1 string, count int, page int, perpage int) (Result, error) {
	Results := Result{}
	var Page Page
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return Results, err
	}
	defer session.Close()
	col := session.DB("yellowListings").C("Listings")
	index := mgo.Index{
		Key: []string{"$text:specialisation", "$text:companyname"},
	}
	err = col.EnsureIndex(index)
	if err != nil {
		return Results, err
	}

	q := col.Find(bson.M{"$text": bson.M{"$search": query1}})
	Page = SearchPagination(count, page, perpage)
	err = q.Limit(perpage).Skip(Page.Skip).All(&Results.Data)
	Results.Pag = Page
	if err != nil {
		return Results, err
	}
	return Results, nil
}

//GetNewView used to return new views
func GetNewView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, _ := RenderView(id)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var data Result
	tmp := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(tmp)
	data.Query = r.URL.Query().Get("q")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	result, _ := Search(data.Query, 50, page, 50)
	newVal, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(newVal)
}
