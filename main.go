package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
	"github.com/smilecs/yellowpages/web"
)

// Router struct would carry the httprouter instance, so its methods could be verwritten and replaced with methds with wraphandler
type Router struct {
	*httprouter.Router
}

// Get is an endpoint to only accept requests of method GET
func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

// Post is an endpoint to only accept requests of method POST
func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

// Put is an endpoint to only accept requests of method PUT
func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

// Delete is an endpoint to only accept requests of method DELETE
func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

// NewRouter is a wrapper that makes the httprouter struct a child of the router struct
func NewRouter() *Router {
	return &Router{httprouter.New()}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	web.TemplateInit()
	config.Init()
	// defer config.Get().Database.Session.Close()
	defer config.Get().BoltDB.Close()
	defer config.Get().BleveIndex.Close()

	middlewares := alice.New(web.LoggingHandler)
	//web.RecoverHandler, context.ClearHandler,
	router := NewRouter()

	router.ServeFiles("/zohoverify/*filepath", http.Dir("assets"))
	router.Get("/google28373290a86b6ef4.html", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/google28373290a86b6ef4.html")
	}))

	fileServer := http.FileServer(http.Dir("./ui/assets"))
	router.GET("/assets/*filepath", func(w http.ResponseWriter, r *http.Request, httpParams httprouter.Params) {
		// w.Header().Set("Vary", "Accept-Encoding")
		// w.Header().Set("Cache-Control", "public, max-age=7776000")
		r.URL.Path = httpParams.ByName("filepath")
		fileServer.ServeHTTP(w, r)
	})

	router.Get("/api/setup/xyzabc123", middlewares.ThenFunc(web.Setup))

	router.Get("/", middlewares.ThenFunc(web.HomeHandler))
	router.Get("/search", middlewares.ThenFunc(web.SearchResultHandler))
	router.Get("/pluslistings", middlewares.ThenFunc(web.GetPlusListingsHandler))
	router.Get("/adverts", middlewares.ThenFunc(web.GetAdvertsHandler))
	router.Get("/categories/:category", middlewares.ThenFunc(web.CategoryListingsHandler))
	router.Get("/listings/:listing", middlewares.ThenFunc(web.SingleListingHandler))

	router.Get("/register_business", middlewares.ThenFunc(web.RegisterListing))
	router.Get("/privacy_policy", middlewares.ThenFunc(web.PrivacyPolicy))

	router.Get("/register_plus_business", middlewares.ThenFunc(web.RegisterPlusListing))

	//router.Post("/login", middlewares.ThenFunc(web.Login))
	router.Post("/api/admin/login", middlewares.ThenFunc(web.AdminLogin))
	router.Get("/api/categories/:category", middlewares.ThenFunc(web.CategoryListingsJSON))
	router.Get("/api/search", middlewares.ThenFunc(web.SearchResultJSON))
	router.Get("/api/pluslistings", middlewares.ThenFunc(web.GetPlusListingsJSON))

	router.Get("/api/analytics", middlewares.ThenFunc(web.GetAnalytics))

	router.Get("/api/categories", middlewares.ThenFunc(web.GetCategories))
	router.Post("/api/categories", middlewares.ThenFunc(web.AddCategory))
	router.Post("/api/listings/add", middlewares.ThenFunc(web.AddListing))
	router.Get("/api/listings/unapproved", middlewares.ThenFunc(web.Getunapproved))
	router.Get("/api/listings/approve", middlewares.ThenFunc(web.Approvehandler))
	router.Get("/api/listings/delete", middlewares.ThenFunc(web.Deletehandler))
	router.Get("/api/listings/slug/:slug", middlewares.ThenFunc(web.SingleListingHandlerJSON))
	router.Post("/api/listings/edit/:slug", middlewares.ThenFunc(web.EditListing))
	router.Post("/api/adverts/new", middlewares.ThenFunc(web.NewAdHandler))
	router.Get("/api/adverts/all", middlewares.ThenFunc(web.GetAdvertsJSON))

	// router.Get("/api/adminList", middlewares.ThenFunc(web.GetAdminsHandler))
	// router.Post("/api/newuser", middlewares.ThenFunc(web.NewUserHandler))

	// router.Post("/api/social_login", middlewares.ThenFunc(web.SocialLogin))
	// router.Post("/api/add_review", middlewares.ThenFunc(web.AddReviews))
	// router.Get("/api/get_reviews", middlewares.ThenFunc(web.ReviewJSON))

	//router.Get("/Upload", middlewares.ThenFunc(web.CsvHandler))

	router.Get("/api/index_data", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		models.IndexData()
	}))
	router.Get("/api/xxx/all_approved", middlewares.ThenFunc(web.Getapproved))

	router.Get("/admin/*filepath", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/admin.html")
	}))

	// SPECIFY WHETHER TO MIGRATE THE FILES
	var migrate = flag.Bool("migrate", false, "Migrate should migrate the listings")
	flag.Parse()
	if *migrate {
		log.Println("I should print out all the listings here...")
		filename := "calabarpages_dump.json"
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("Error reading %s: %v", filename, err)
		}
		var listings models.Listings
		err = json.Unmarshal(fileContent, &listings)
		log.Printf("Lenght of data: %d\n", len(listings.Data))
		listings.Add() // this function adds all the listings.Data to the db
	} else {
		log.Println("No migration flag")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")
		PORT = "8080"
	}

	handler := cors.New(cors.Options{
		//		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedOrigins: []string{"*"},

		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-Auth-Token", "*"},
		Debug:            false,
	}).Handler(router)
	log.Printf("serving on port: %s\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
