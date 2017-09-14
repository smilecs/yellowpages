package models

import (
	"encoding/json"
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/gosimple/slug"
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
func (r Category) Add(conf *config.Conf) error {

	// p := rand.New(rand.NewSource(time.Now().UnixNano()))
	// str := strconv.Itoa(p.Intn(10))
	// r.Slug = strings.Replace(r.Category, " ", "-", -1) + str
	// r.Slug = strings.Replace(r.Slug, "&", "-", -1) + str

	if r.Slug == "" {
		r.Slug = slug.Make(r.Slug)
	}

	r.Show = "true"

	rJSONbyt, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.CATEGORIESCOLLECTION))

		err := b.Put([]byte(r.Slug), rJSONbyt)
		return err
	})
	//
	// mgoSession := conf.Database.Session.Copy()
	// defer mgoSession.Close()
	//
	// collection := conf.Database.C("Category").With(mgoSession)
	// err := collection.Insert(r)
	// if err != nil {
	// 	log.Println(err)
	// }
	return err
}

func (r Category) GetOne(conf *config.Conf, id string) (Category, error) {
	result := Category{}
	jsonByt := []byte{}
	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.CATEGORIESCOLLECTION))
		jsonByt = b.Get([]byte(id))
		return err
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	err = json.Unmarshal(jsonByt, &result)
	if err != nil {
		log.Println(err)
		return result, err
	}

	return result, nil
	//
	// mgoSession := conf.Database.Session.Copy()
	// defer mgoSession.Close()
	//
	// collection := conf.Database.C("Category").With(mgoSession)
	// err := collection.Find(bson.M{"slug": id}).One(&result)
	// if err != nil {
	// 	log.Println(err)
	// 	return result, err
	// }
	// return result, nil
}

//Getcat list
func (r Category) GetAll(conf *config.Conf) ([]Category, error) {

	results := []Category{}

	jsonBytArray := [][]byte{}
	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.CATEGORIESCOLLECTION))
		b.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			jsonBytArray = append(jsonBytArray, v)
			return nil
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		return results, err
	}

	for _, v := range jsonBytArray {
		category := Category{}
		err = json.Unmarshal(v, &category)
		if err != nil {
			log.Println(err)
		}
		results = append(results, category)
	}
	return results, nil
	//
	// //collection := conf.Database.C("Category")
	// mgoSession := conf.Database.Session.Copy()
	// defer mgoSession.Close()
	//
	// collection := conf.Database.C("Category").With(mgoSession)
	// err := collection.Find(bson.M{"show": "true"}).All(&result)
	// if err != nil {
	// 	return result, err
	// }
	// return result, nil
}
