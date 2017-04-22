package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

//FrontAdminHandler for serving admin page
func FrontAdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/index.html")
}

//AddListingViewHandler to render veiw page
func AddListingViewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/addlisting.html")
}

func AddListingView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/newapp.html")
}

//UnapprovedViewHandler to render veiw page
func UnapprovedViewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/view_listings.html")
}

//AdvertHandler for adding adverts
func AdvertHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/newAd.html")
}

func CustHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/newapp.html")
}

//NewAdvertHandler for new Adverts
func NewAdvertHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/newAd.html")
}

func NoPaymentAfter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cust/assests/partials/warning.html")
}

func PaymentAfter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cust/newapp.html")
}

func ClientAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cust/index.html")
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/login.html")
}

func ClientIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cust/addList.html")
}

//HomeHandler for listing view
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.Category{}.GetAll(config.Get())
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Categories []models.Category
	}{
		Categories: categories,
	}
	tmp := GetTemplates().Lookup("index.html")
	tmp.Execute(w, data)
}

//CategoryHandler for view
func SingleListingHandler(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)
	slug := params.ByName("listing")
	log.Printf("params %+v", params)

	listing, err := models.Listing{}.GetOne(config.Get(), slug)
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Listing models.Listing
	}{
		Listing: listing,
	}
	tmp := GetTemplates().Lookup("single.html")
	log.Printf("%+v", tmp)
	tmp.Execute(w, data)
}

//HomeHandler for listing view
func RegisterListing(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Plus bool
	}{
		Plus: false,
	}
	tmp := GetTemplates().Lookup("register_business.html")
	tmp.Execute(w, data)
}

func PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	tmp := GetTemplates().Lookup("privacy_policy.html")
	tmp.Execute(w, "")
}

//HomeHandler for listing view
func RegisterPlusListing(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Plus bool
	}{
		Plus: true,
	}
	tmp := GetTemplates().Lookup("register_business.html")
	tmp.Execute(w, data)
}
