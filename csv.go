package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Forms struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyName    string        `bson:"companyname"`
	Address        string        `bson:"address"`
	Hotline        string        `bson:"hotline"`
	Specialisation string        `bson:"specialisation"`
	Category       string        `bson:"category"`
	Advert         string        `bson:"advert"`
	Size           string        `bson:"size"`
	Image          string        `bson:"image"`
	Images         []string      `bson:"images"`
	Slug           string        `bson:"slug"`
	About          string        `bson:"about"`
	RC             string        `bson:"rc"`
	Branch         string        `bson:"branch"`
	Product        string        `bson:"product"`
	Email          string        `bson:"email"`
	Website        string        `bson:"website"`
	DHr            string        `bson:"dhr"`
	Verified       string        `bson:"verified"`
	Approved       bool          `bson:"approved"`
	Plus           string        `bson:"plus"`
}

type Cat struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Category string        `bson:"category"`
	Slug     string        `bson:"slug"`
}

func MainSeal() {
	s, err := mgo.Dial("mongodb://localhost")
	defer s.Close()
	if err != nil {
		panic(err)
	}
	collection := s.DB("yellowListings").C("Listings")
	col := s.DB("yellowListings").C("Category")
	files, err := ioutil.ReadDir("csvs/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		dat, err := ioutil.ReadFile("csvs/" + file.Name())
		//fmt.Println(file.Name())
		check(err)
		r := csv.NewReader(strings.NewReader(string(dat)))
		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		cat := new(Cat)
		cat.Category = records[0][0]
		cat.Slug = strings.Replace(records[0][0], " ", "-", -1)
		cat.Slug = strings.Replace(cat.Slug, "&", "AND", -1)
		col.Insert(cat)
		ru := 9 + 1
		for i := 4; i < len(records); i++ {
			data := records[i]
			fmt.Println(i)
			fmt.Println(data)
			form := new(Forms)
			for k := 0; k < len(data); k++ {
				switch k {
				case 1:
					form.CompanyName = data[k]
					fmt.Println("reached")
				case 2:
					form.Address = data[k]
				case 3:
					form.Hotline = data[k]
				case 4:
					form.Specialisation = data[k]
				case 5:
					form.RC = data[k]
				case 6:
					form.Branch = data[k]
				case 7:
					form.Product = data[k]
				case 8:
					form.Email = data[k]
				case 9:
					form.Website = data[k]
				case ru:
					form.DHr = data[k]
				}
				//tmp := strings.Split(data[k], "'\n'")
			}
			p := rand.New(rand.NewSource(time.Now().UnixNano()))
			str := strconv.Itoa(p.Intn(10))
			form.Slug = strings.Replace(form.CompanyName, " ", "-", -1) + str
			form.Category = cat.Slug
			form.Approved = true
			form.Plus = "false"
			collection.Insert(form)
		}
	}
}
