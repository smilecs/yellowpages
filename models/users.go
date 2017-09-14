package models

//
// import (
// 	"encoding/json"
// 	"log"
//
// 	bolt "github.com/coreos/bbolt"
// 	"github.com/smilecs/yellowpages/config"
// )
//
// //Listing struct holds all submitted form data for listings
// type User struct {
// 	Name  string
// 	Email string
// 	Type  string
// 	ID    string
// 	Link  string
// }
//
// //Addlisting function adding listings data to db
// func (r User) Add(conf *config.Conf) error {
// 	log.Println(r)
//
// 	rJSONByt, err := json.Marshal(r)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	err = conf.BoltDB.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(config.USERSCOLLECTION))
// 		err := b.Put([]byte(r.ID), rJSONByt)
// 		if err != nil {
// 			log.Println(err)
// 			return err
// 		}
// 		return nil
// 	})
// 	//
// 	// mgoSession := conf.Database.Session.Copy()
// 	// defer mgoSession.Close()
// 	//
// 	// collection := conf.Database.C(config.USERSCOLLECTION).With(mgoSession)
// 	// _, err := collection.Upsert(bson.M{"id": r.ID}, r)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return err
// 	// }
// 	return nil
// }
