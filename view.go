package main

import (
	"encoding/json"
	"net/http"

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

//RenderView function to return listings mixed with adds
func RenderView(id string) ([]*View, error) {
	tmp := []Form{}
	result := []*View{}
	res := []Advert{}
	session, err := mgo.Dial(config.xx)
	if err != nil {
		return result, err
	}
	defer session.Close()
	adCollection := session.DB("yellowListings").C("Adverts")
	err = adCollection.Find(bson.M{}).All(&res)
	if err != nil {
		return result, err
	}
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

//GetNewView used to return new views
func GetNewView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	data, _ := RenderView(id)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
