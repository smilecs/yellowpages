package main

import (
	"net/http"
)

//FrontAdminHandler for serving admin page
func FrontAdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/index.html")
}

//AddListingViewHandler to render veiw page
func AddListingViewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/addlisting.html")
}

//UnapprovedViewHandler to render veiw page
func UnapprovedViewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/view_listings.html")
}

func addcatViewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/partials/addcat.html")
}
