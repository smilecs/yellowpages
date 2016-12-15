package web

import (
	"time"

	"github.com/tonyalaribe/yellowpages/models"
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
	Image          string    `bson:"image"`
	Images         []string  `bson:"images"`
	Slug           string    `bson:"slug"`
	About          string    `bson:"about"`
	Username       string    `bson:"username"`
	Password       string    `bson:"password"`
	RC             string    `bson:"rc"`
	Branch         string    `bson:"branch"`
	Product        string    `bson:"product"`
	Email          string    `bson:"email"`
	Website        string    `bson:"website"`
	Dhr            string    `bson:"dhr"`
	Verified       string    `bson:"verified"`
	Approved       bool      `bson:"approved"`
	Plus           string    `bson:"plus"`
	Expiry         time.Time `bson:"expiry"`
	Duration       string    `bson:"duration"`
}

type Post struct {
	Type    string
	Listing models.Listing
	Advert  models.Advert
}

type QuickTeller struct {
	PaymentReference string
	Amount           string
	TransactionDate  string
	ResponseCode     string
}

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}

type NewView struct {
	Data []*View
	Page models.Page
}

type Result struct {
	Data  []models.Listing
	Page  models.Page
	Query string
}

/*
//GetAdmins

func Get_Params(w http.ResponseWriter, r *http.Request) {
	log.Println(r.FormValue("resp_code"))
	log.Println("here")

}

func Post_Params(w http.ResponseWriter, r *http.Request) {
	var q string
	tmp := r.URL.Query().Get("q")
	if tmp != "" {
		if r.URL.Query().Get("q") == "3000" {
			q = "1"
		} else {
			q = "2"
		}
		http.Redirect(w, r, "/newapp?q="+q, http.StatusSeeOther)
	}

	//http.ServeFile(w, r, "cust/newapp.html?q=odpos")
}



func SliderHandler(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(tmp)
	data, _ := SliderView(page, 50)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

//GetNewView used to return new views
func GetNewView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("q")
	tmp := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(tmp)
	data, _ := RenderView(id, 50, page, 50)
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}


func CsvHandler(w http.ResponseWriter, r *http.Request) {
	MainSeal()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("done"))
}
*/
