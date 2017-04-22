package models

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/smilecs/yellowpages/config"
	"gopkg.in/mgo.v2/bson"
)

//Category struct for use in registration
type Category struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Category string        `bson:"category"`
	Slug     string        `bson:"slug"`
	Show     string        `bson:"show"`
}

//Addcat function for adding category
func (r Category) Add(config *config.Conf) error {

	p := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := strconv.Itoa(p.Intn(10))
	r.Slug = strings.Replace(r.Category, " ", "-", -1) + str
	r.Slug = strings.Replace(r.Slug, "&", "-", -1) + str
	r.Show = "true"
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()

	collection := config.Database.C("Category").With(mgoSession)
	err := collection.Insert(r)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r Category) GetOne(config *config.Conf, id string) (Category, error) {
	result := Category{}
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()

	collection := config.Database.C("Category").With(mgoSession)
	err := collection.Find(bson.M{"slug": id}).One(&result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

//Getcat list
func (r Category) GetAll(config *config.Conf) ([]Category, error) {

	result := []Category{}
	//collection := config.Database.C("Category")
	mgoSession := config.Database.Session.Copy()
	defer mgoSession.Close()

	collection := config.Database.C("Category").With(mgoSession)
	err := collection.Find(bson.M{"show": "true"}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
