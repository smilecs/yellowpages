package models

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/gosimple/slug"
	uuid "github.com/satori/go.uuid"
	"github.com/smilecs/yellowpages/config"
)

//Listing struct holds all submitted form data for listings
type Listing struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyName    string `bson:"companyname"`
	Address        string `bson:"address"`
	Hotline        string `bson:"hotline"`
	HotlinesList   []string
	Specialisation string `bson:"specialisation"`
	Category       string `bson:"category"`
	// Advert         string    `bson:"advert"`
	// Size           string    `bson:"size"`
	Image    string    `bson:"image"`
	Images   []string  `bson:"images"`
	Slug     string    `bson:"slug"`
	About    string    `bson:"about"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
	RC       string    `bson:"rc"` //
	Branch   string    `bson:"branch"`
	Product  string    `bson:"product"`
	Email    string    `bson:"email"`
	Website  string    `bson:"website"`
	DHr      string    `bson:"dhr"` //Working days and hours
	Verified string    `bson:"verified"`
	Approved bool      `bson:"approved"`
	Plus     string    `bson:"plus"`
	Expiry   time.Time `bson:"expiry"`
	Duration string    `bson:"duration"`
	Date     time.Time `bson:"date"`
	Pg       Page
	Reviews  string `bson:"reviews"`
}

type Listings struct {
	Data  []Listing
	Page  Page
	Query string
}

//Addlisting function adding listings data to db
func (r Listing) Add(conf *config.Conf) error {

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("yellowpagesng")

	p := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Date = time.Now()
	if r.Image != "" {
		tm, _ := strconv.Atoi(r.Duration)
		t := r.Date.AddDate(tm, 1, 0)
		r.Expiry = t
		if len(strings.Split(r.Image, "base64,")) > 1 {
			byt, er := base64.StdEncoding.DecodeString(strings.Split(r.Image, "base64,")[1])
			if er != nil {
				log.Println(er)
			}

			meta := strings.Split(r.Image, "base64,")[0]
			newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
			imagename := "listings/" + uuid.NewV1().String()

			err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite, s3.Options{})
			if err != nil {
				log.Println(err)
			}

			log.Println(bucket.URL(imagename))
			r.Image = bucket.URL(imagename)
		}

	}

	var images []string
	for _, v := range r.Images {
		if len(strings.Split(v, "base64,")) > 1 {
			var byt []byte

			byt, err = base64.StdEncoding.DecodeString(strings.Split(v, "base64,")[1])
			if err != nil {
				log.Println(err)
			}

			meta := strings.Split(v, "base64,")[0]
			newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
			imagename := "listings/" + uuid.NewV1().String()

			err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite, s3.Options{})
			if err != nil {
				log.Println(err)
			}
			images = append(images, bucket.URL(imagename))
		}
	}

	r.Images = images

	// r.Slug = strings.Replace(r.CompanyName, " ", "-", -1) + str
	// r.Slug = strings.Replace(r.Slug, "&", "-", -1) + str
	// r.Slug = strings.Replace(r.Slug, "/", "-", -1) + str
	// r.Slug = strings.Replace(r.Slug, ",", "-", -1) + str
	str := strconv.Itoa(p.Intn(10))
	r.Slug = slug.Make(r.CompanyName + " " + str)

	rJSONbyt, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))

		log.Print(string(rJSONbyt))
		log.Println(r.Slug)
		err := b.Put([]byte(r.Slug), rJSONbyt)
		return err
	})
	IndexSingleListingWithBleve(r)

	return err
}

func (r Listing) Edit(conf *config.Conf) error {
	log.Println(r)
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("yellowpagesng")

	if r.Image != "" {
		if len(strings.Split(r.Image, "base64,")) > 1 {
			meta := strings.Split(r.Image, "base64,")[0]
			tm, _ := strconv.Atoi(r.Duration)
			t := r.Date.AddDate(tm, 1, 0)
			r.Expiry = t
			byt, er := base64.StdEncoding.DecodeString(strings.Split(r.Image, "base64,")[1])
			if er != nil {
				log.Println(er)
			}

			newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
			imagename := "listings/" + uuid.NewV1().String()

			err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite, s3.Options{})
			if err != nil {
				log.Println(err)
			}

			log.Println(bucket.URL(imagename))
			r.Image = bucket.URL(imagename)
		}
	}

	var images []string
	for _, v := range r.Images {

		if len(strings.Split(v, "base64,")) > 1 {

			var byt []byte
			byt, err = base64.StdEncoding.DecodeString(strings.Split(v, "base64,")[1])
			if err != nil {
				log.Println(err)
			}

			meta := strings.Split(v, "base64,")[0]
			newmeta := strings.Replace(strings.Replace(meta, "data:", "", -1), ";", "", -1)
			imagename := "listings/" + uuid.NewV1().String()

			err = bucket.Put(imagename, byt, newmeta, s3.PublicReadWrite, s3.Options{})
			if err != nil {
				log.Println(err)
			}
			images = append(images, bucket.URL(imagename))
		} else {
			images = append(images, v)
		}
	}

	r.Images = images

	oldListing := Listing{}
	jsonByt := []byte{}

	// var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		jsonByt = b.Get([]byte(r.Slug))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &oldListing)
	if err != nil {
		log.Println(err)
	}

	oldListing.CompanyName = r.CompanyName
	oldListing.Address = r.Address
	oldListing.Hotline = r.Hotline
	oldListing.Specialisation = r.Specialisation
	oldListing.Category = r.Category
	oldListing.Image = r.Image
	oldListing.Images = r.Images
	oldListing.Slug = r.Slug
	oldListing.About = r.About
	oldListing.RC = r.RC
	oldListing.Branch = r.Branch
	oldListing.Product = r.Product
	oldListing.Email = r.Email
	oldListing.Website = r.Website
	oldListing.DHr = r.DHr

	rJSONbyt, err := json.Marshal(oldListing)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))

		err := b.Put([]byte(oldListing.Slug), rJSONbyt)
		return err
	})

	err = IndexSingleListingWithBleve(r)
	if err != nil {
		log.Println(err)
	}
	return err

}

func (r Listing) GetOne(conf *config.Conf, id string) (Listing, error) {
	result := Listing{}

	jsonByt := []byte{}
	log.Println(r)
	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		jsonByt = b.Get([]byte(id))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &result)
	if err != nil {
		log.Println(err)
	}

	return result, nil
}

func StringToPhoneNumbers(s string) []string {
	var response []string
	noCommaList := strings.Split(s, ",")
	for _, j := range noCommaList {
		noSpaceList := strings.Split(j, " ")
		for _, jj := range noSpaceList {
			if jj != " " && jj != "," && jj != "" {
				response = append(response, strings.TrimSpace(jj))
			}
		}
	}
	return response
}

//GetListings return listings
func (r Listing) GetAllApproved(conf *config.Conf) (Listings, error) {
	results := []Listing{}
	listings := Listings{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		b.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			listing := Listing{}
			json.Unmarshal(v, &listing)
			if listing.Approved {
				listing.HotlinesList = StringToPhoneNumbers(listing.Hotline)
				results = append(results, listing)
			}

			return nil
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		return listings, err
	}

	listings.Data = results
	return listings, nil
}

//Getunapproved function for Getunapproved handler
func (r Listing) GetAllUnapproved(conf *config.Conf) (Listings, error) {

	results := []Listing{}
	listings := Listings{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		b.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			listing := Listing{}
			json.Unmarshal(v, &listing)
			if !listing.Approved {
				listing.HotlinesList = StringToPhoneNumbers(listing.Hotline)
				results = append(results, listing)
			}

			// jsonBytArray = append(jsonBytArray, v)
			return nil
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		return listings, err
	}

	listings.Data = results
	return listings, nil

}

//GetcatListing function
func (r Listing) GetAllInCategory(conf *config.Conf, id string, page int) (Listings, error) {
	perPage := 10
	Results := Listings{}

	searchResult, err := SearchField(id, "Category", (page-1)*perPage)
	if err != nil {
		log.Println(err)
	}

	count := int(searchResult.Total)
	pg := SearchPagination(count, page, perPage)

	jsonBytArray := [][]byte{}
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		for _, v := range searchResult.Hits {
			jsonByt := b.Get([]byte(v.ID))
			jsonBytArray = append(jsonBytArray, jsonByt)
		}

		return err
	})
	if err != nil {
		log.Println(err)
	}

	listings := []Listing{}
	for _, v := range jsonBytArray {
		listing := Listing{}
		err = json.Unmarshal(v, &listing)
		if err != nil {
			log.Println(err)
		}
		listing.HotlinesList = StringToPhoneNumbers(listing.Hotline)
		listings = append(listings, listing)
	}
	Results.Data = listings

	Results.Page = pg
	if err != nil {
		return Results, err
	}
	return Results, nil

}

//GetcatListing function
func (r Listing) GetAllPlusListings(conf *config.Conf) (Listings, error) {

	results := []Listing{}
	listings := Listings{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		b.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			listing := Listing{}
			json.Unmarshal(v, &listing)
			if listing.Plus == "true" {
				listing.HotlinesList = StringToPhoneNumbers(listing.Hotline)
				results = append(results, listing)
			}

			// jsonBytArray = append(jsonBytArray, v)
			return nil
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		return listings, err
	}

	listings.Data = results
	return listings, nil
}

//Update to approve of a listing to show in the client side
func (r Listing) Approve(conf *config.Conf, slug string) error {

	oldListing := Listing{}
	jsonByt := []byte{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		jsonByt = b.Get([]byte(r.Slug))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &oldListing)
	if err != nil {
		log.Println(err)
	}

	oldListing.Approved = true

	rJSONbyt, err := json.Marshal(oldListing)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))

		err := b.Put([]byte(oldListing.Slug), rJSONbyt)
		return err
	})

	err = IndexSingleListingWithBleve(oldListing)
	if err != nil {
		log.Println(err)
	}
	return nil
}

//Update to approve of a listing to show in the client side
func (r Listing) Delete(conf *config.Conf, slug string) error {

	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))

		b.Delete([]byte(r.Slug))
		return nil
	})

	err := DeleteSingleListingInBleve(slug)
	if err != nil {
		log.Println(err)
	}
	return nil
}

//TimeUpdate to change time
func (r Listing) TimeUpdate(conf *config.Conf, id string, expiry string) error {
	x, _ := strconv.Atoi(expiry)

	oldListing := Listing{}
	jsonByt := []byte{}

	var err error
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		jsonByt = b.Get([]byte(r.Slug))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonByt, &oldListing)
	if err != nil {
		log.Println(err)
	}

	oldListing.Expiry = time.Now().AddDate(x, 1, 0)

	rJSONbyt, err := json.Marshal(oldListing)
	if err != nil {
		log.Println(err)
	}
	conf.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))

		err := b.Put([]byte(oldListing.Slug), rJSONbyt)
		return err
	})

	return nil
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

//Search searches
func (r Listing) Search(conf *config.Conf, query string, page int) (Listings, error) {
	perPage := 10
	Results := Listings{}
	log.Println("in search")

	searchResult, err := SearchWithIndex(query, (page-1)*perPage)
	if err != nil {
		log.Println(err)
	}

	count := int(searchResult.Total)
	pg := SearchPagination(count, page, perPage)

	jsonBytArray := [][]byte{}
	err = conf.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.LISTINGSCOLLECTION))
		for _, v := range searchResult.Hits {
			jsonByt := b.Get([]byte(v.ID))
			if len(jsonByt) > 0 {
				jsonBytArray = append(jsonBytArray, jsonByt)
			}

		}

		return err
	})
	if err != nil {
		log.Println(err)
	}
	listings := []Listing{}
	for _, v := range jsonBytArray {
		listing := Listing{}
		err = json.Unmarshal([]byte(v), &listing)
		if err != nil {
			log.Println(err)
		}
		listing.HotlinesList = StringToPhoneNumbers(listing.Hotline)
		listings = append(listings, listing)
	}
	Results.Data = listings

	Results.Page = pg
	if err != nil {
		return Results, err
	}
	return Results, nil
}
