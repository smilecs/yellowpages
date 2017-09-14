package models

//
// import (
// 	"encoding/json"
// 	"log"
//
// 	"strconv"
//
// 	bolt "github.com/coreos/bbolt"
// 	"github.com/smilecs/yellowpages/config"
// 	"gopkg.in/mgo.v2/bson"
// )
//
// type Reviews struct {
// 	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
// 	Comment  string        `bson:"comment"`
// 	Rating   int           `bson:"rating"`
// 	SocialId string        `bson:"socialid"`
// 	ImageUrl string        `bson:"imageurl"`
// 	Slug     string        `bson:"slug"`
// 	Name     string        `bson:"name"`
// }
//
// type ReviewList struct {
// 	Data  []Reviews
// 	Page  Page
// 	Query string
// }
//
// func (r Reviews) Add(conf *config.Conf) error {
// 	// mgoSession := config.Database.Session.Copy()
// 	// defer mgoSession.Close()
// 	// collection := config.Database.C("Reviews").With(mgoSession)
// 	// collection2 := config.Database.C("Listings").With(mgoSession)
// 	// _, err := collection.Upsert(bson.M{"socialid": r.SocialId}, r)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
//
// 	jsonByt := []byte{}
//
// 	rJSONbyt, err := json.Marshal(r)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	conf.BoltDB.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(config.REVIEWSCOLLECTION))
//
// 		err = b.Put([]byte(r.SocialId), rJSONbyt)
// 		return err
// 	})
//
// 	count, _ := r.GetSize(config, r.Slug)
// 	val := strconv.Itoa(count)
// 	log.Println("next" + val)
// 	// updat := bson.M{"reviews": val}
// 	// query := bson.M{"slug": r.Slug}
// 	// collection2.Update(query, bson.M{"$set": updat})
// 	//
//
// 	old := Listing{}
// 	jsonByt := []byte{}
//
// 	// var err error
// 	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
// 		jsonByt = b.Get([]byte(r.Slug))
// 		return err
// 	})
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = json.Unmarshal(jsonByt, &old)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	old.Reviews = val
//
// 	rJSONbyt, err := json.Marshal(old)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	conf.BoltDB.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
//
// 		err := b.Put([]byte(old.Username), rJSONbyt)
// 		return err
// 	})
//
// 	return err
//
// }
//
// func (r Reviews) GetAll(config *config.Conf, page int, query string) (ReviewList, error) {
// 	data := ReviewList{}
// 	perPage := 10
// 	mgoSession := config.Database.Session.Copy()
// 	defer mgoSession.Close()
// 	collection := config.Database.C("Reviews").With(mgoSession)
// 	q := collection.Find(bson.M{"slug": query})
// 	count, err := q.Count()
// 	log.Println(count)
// 	if err != nil {
// 		return data, err
// 	}
//
// 	pg := SearchPagination(count, page, perPage)
// 	err = q.Limit(perPage).Skip(pg.Skip).All(&data.Data)
// 	data.Page = pg
//
// 	if err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }
//
// func (r Reviews) GetSize(config *config.Conf, query string) (int, error) {
// 	mgoSession := config.Database.Session.Copy()
// 	defer mgoSession.Close()
// 	collection := config.Database.C("Reviews").With(mgoSession)
// 	q := collection.Find(bson.M{"slug": query})
// 	count, err := q.Count()
// 	log.Println(count)
// 	if err != nil {
// 		return count, err
// 	}
// 	return count, nil
// }
