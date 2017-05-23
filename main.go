package main

import (
	"context"
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
		//r.Context().Set(r, "params", ps)
		//r.Context().
		//context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	web.TemplateInit()
	config.Init()
	defer config.Get().Database.Session.Close()
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

	router.Get("/admin", middlewares.ThenFunc(web.FrontAdminHandler))
	router.Get("/login", middlewares.ThenFunc(web.LoginAdmin))
	router.Get("/Newlisting", middlewares.ThenFunc(web.ClientIndex))
	router.Get("/addlistingtemp", middlewares.ThenFunc(web.AddListingViewHandler))
	router.Get("/addlisting", middlewares.ThenFunc(web.AddListingView))

	//router.Post("/login", middlewares.ThenFunc(web.Login))
	router.Post("/adminlogin", middlewares.ThenFunc(web.AdminLoginOld))
	router.Post("/api/admin/login", middlewares.ThenFunc(web.AdminLogin))
	router.Get("/viewlistingtemp", middlewares.ThenFunc(web.UnapprovedViewHandler))

	router.Get("/api/categories/:category", middlewares.ThenFunc(web.CategoryListingsJSON))
	router.Get("/api/search", middlewares.ThenFunc(web.SearchResultJSON))
	router.Get("/api/pluslistings", middlewares.ThenFunc(web.GetPlusListingsJSON))

	//Remove
	router.Post("/api/addcat", middlewares.ThenFunc(web.AddCategory))
	router.Get("/api/getcat", middlewares.ThenFunc(web.GetCategories))
	router.Post("/api/addlisting", middlewares.ThenFunc(web.AddListing))
	router.Get("/api/unapproved", middlewares.ThenFunc(web.Getunapproved))

	router.Get("/api/categories", middlewares.ThenFunc(web.GetCategories))
	router.Post("/api/categories", middlewares.ThenFunc(web.AddCategory))
	router.Post("/api/listings/add", middlewares.ThenFunc(web.AddListing))
	router.Get("/api/listings/unapproved", middlewares.ThenFunc(web.Getunapproved))

	router.Get("/api/adminList", middlewares.ThenFunc(web.GetAdminsHandler))
	router.Post("/api/newuser", middlewares.ThenFunc(web.NewUserHandler))
	router.Post("/api/newAd", middlewares.ThenFunc(web.NewAdHandler))

	router.Post("/api/approve", middlewares.ThenFunc(web.Approvehandler))

	router.Post("/api/social_login", middlewares.ThenFunc(web.SocialLogin))
	router.Post("/api/add_review", middlewares.ThenFunc(web.AddReviews))
	router.Get("/api/get_reviews", middlewares.ThenFunc(web.ReviewJSON))

	router.Get("/newad", middlewares.ThenFunc(web.NewAdvertHandler))

	router.Get("/addcattemp", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "admin/partials/addcat.html")
	}))
	router.ServeFiles("/temp/*filepath", http.Dir("admin/partials"))

	router.ServeFiles("/cust/partials/*filepath", http.Dir("cust/partials"))
	router.ServeFiles("/cust/js/*filepath", http.Dir("cust/js"))

	router.Get("/api/adverts", middlewares.ThenFunc(web.GetAdvertsJSON))
	//router.Get("/Upload", middlewares.ThenFunc(web.CsvHandler))

	router.Get("/api/index_data", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		models.IndexMongoDBListingsCollectionWithBleve()
	}))

	router.Get("/dashboard/*filepath", middlewares.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/admin.html")
	}))

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
	log.Println("serving ")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
