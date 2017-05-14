package models

import (
	"log"

	"strconv"

	"github.com/smilecs/yellowpages/config"
	"gopkg.in/mgo.v2/bson"
)

type Reviews struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Comment  string        `bson:"comment"`
	Rating   int           `bson:"rating"`
	SocialId string        `bson:"socialid"`
	ImageUrl string        `bson:"imageurl"`
	Slug     string        `bson:"slug"`
	Name     string        `bson:"name"`
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
	collection2 := config.Database.C("Listings").With(mgoSession)
	_, err := collection.Upsert(bson.M{"socialid": r.SocialId}, r)
	if err != nil {
		log.Println(err)
	}
	count, _ := r.GetSize(config, r.Slug)
	val := strconv.Itoa(count)
	log.Println("next" + val)
	updat := bson.M{"reviews": val}
	query := bson.M{"slug": r.Slug}
	collection2.Update(query, bson.M{"$set": updat})
	return err
}

func (r Reviews) GetAll(config *config.Conf, page int, query string) (ReviewList, error) {
	data := ReviewList{}
	perPage := 10
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()
	collection := config.Database.C("Reviews").With(mgoSession)
	q := collection.Find(bson.M{"slug": query})
	count, err := q.Count()
	log.Println(count)
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

func (r Reviews) GetSize(config *config.Conf, query string) (int, error) {
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()
	collection := config.Database.C("Reviews").With(mgoSession)
	q := collection.Find(bson.M{"slug": query})
	count, err := q.Count()
	log.Println(count)
	if err != nil {
		return count, err
	}
	return count, nil
}
