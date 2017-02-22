package models

import (
	"log"

	"github.com/tonyalaribe/yellowpages/config"
	"gopkg.in/mgo.v2/bson"
)

type Reviews struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Comment  string        `bson:"comment"`
	Rating   int           `bson:"rating"`
	SocialId string        `bson:"socialid"`
	ImageUrl string        `bson:"imageurl"`
}

type ReviewList struct {
	Data  []Reviews
	Page  Page
	Query string
}

func (r Reviews) Add(config *config.Conf) error {
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()
	collection := config.Database.C("Reviews").With(mgoSession)
	err := collection.Insert(r)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r Reviews) GetAll(config *config.Conf, page int) (ReviewList, error) {
	result := Reviews{}
	data := ReviewList{}
	perPage := 10
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()
	collection := config.Database.C("Reviews").With(mgoSession)
	q := collection.Find(bson.M{}).Sort("+")
	count, err := q.Count()
	if err != nil {
		return data, err
	}

	pg := SearchPagination(count, page, perPage)
	err = q.Limit(perPage).Skip(pg.Skip).All(&data.Data)
	data.Page = pg
	if err != nil {
		return data, err
	}
	return data, nil
}
