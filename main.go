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
	"github.com/tonyalaribe/yellowpages/config"
	"github.com/tonyalaribe/yellowpages/web"
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
	commonHandlers := alice.New(web.LoggingHandler)
	//web.RecoverHandler, context.ClearHandler,
	router := NewRouter()

	router.ServeFiles("/zohoverify/*filepath", http.Dir("assets"))

	router.Get("/google28373290a86b6ef4.html", commonHandlers.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/google28373290a86b6ef4.html")
	}))

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	router.Get("/", commonHandlers.ThenFunc(web.HomeHandler))
	router.Get("/search", commonHandlers.ThenFunc(web.SearchResultHandler))
	router.Get("/pluslistings", commonHandlers.ThenFunc(web.GetPlusListingsHandler))
	router.Get("/adverts", commonHandlers.ThenFunc(web.GetAdvertsHandler))
	router.Get("/categories/:category", commonHandlers.ThenFunc(web.CategoryListingsHandler))
	router.Get("/listings/:listing", commonHandlers.ThenFunc(web.SingleListingHandler))

	router.Get("/register_business", commonHandlers.ThenFunc(web.RegisterListing))
	router.Get("/privacy_policy", commonHandlers.ThenFunc(web.PrivacyPolicy))

	router.Get("/register_plus_business", commonHandlers.ThenFunc(web.RegisterPlusListing))

	router.Get("/admin", commonHandlers.ThenFunc(web.FrontAdminHandler))
	router.Get("/login", commonHandlers.ThenFunc(web.LoginAdmin))
	router.Get("/Newlisting", commonHandlers.ThenFunc(web.ClientIndex))
	router.Get("/addlistingtemp", commonHandlers.ThenFunc(web.AddListingViewHandler))
	router.Get("/addlisting", commonHandlers.ThenFunc(web.AddListingView))

	//router.Post("/login", commonHandlers.ThenFunc(web.Login))
	router.Post("/adminlogin", commonHandlers.ThenFunc(web.AdminLogin))
	router.Get("/viewlistingtemp", commonHandlers.ThenFunc(web.UnapprovedViewHandler))

	router.Get("/api/categories/:category", commonHandlers.ThenFunc(web.CategoryListingsJSON))
	router.Get("/api/search", commonHandlers.ThenFunc(web.SearchResultJSON))
	router.Get("/api/pluslistings", commonHandlers.ThenFunc(web.GetPlusListingsJSON))

	router.Get("/api/unapproved", commonHandlers.ThenFunc(web.Getunapproved))
	router.Post("/api/addcat", commonHandlers.ThenFunc(web.AddCategory))
	router.Get("/api/getcat", commonHandlers.ThenFunc(web.GetCategories))
	router.Get("/api/adminList", commonHandlers.ThenFunc(web.GetAdminsHandler))
	router.Post("/api/newuser", commonHandlers.ThenFunc(web.NewUserHandler))
	router.Post("/api/newAd", commonHandlers.ThenFunc(web.NewAdHandler))
	router.Post("/api/addlisting", commonHandlers.ThenFunc(web.AddListing))
	router.Post("/api/approve", commonHandlers.ThenFunc(web.Approvehandler))

	router.Post("/api/social_login", commonHandlers.ThenFunc(web.SocialLogin))
	router.Post("/api/add_review", commonHandlers.ThenFunc(web.AddReviews))
	router.Get("/api/get_reviews", commonHandlers.ThenFunc(web.reviewJSON))

	router.Get("/newad", commonHandlers.ThenFunc(web.NewAdvertHandler))

	router.Get("/addcattemp", commonHandlers.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "admin/partials/addcat.html")
	}))
	router.ServeFiles("/temp/*filepath", http.Dir("admin/partials"))

	router.ServeFiles("/cust/partials/*filepath", http.Dir("cust/partials"))
	router.ServeFiles("/cust/js/*filepath", http.Dir("cust/js"))

	router.Get("/api/adverts", commonHandlers.ThenFunc(web.GetAdvertsJSON))

	/*
		router.Get("/client", commonHandlers.ThenFunc(ClientAdmin))
		router.Get("/Newlisting", commonHandlers.ThenFunc(ClientIndex))

		router.Get("/client/update", commonHandlers.ThenFunc(TimeUpdatehandler))

		//router.Get("/api/plus", commonHandlers.ThenFunc(web.GetPlusPayHandler))

		router.Get("/category/*routes", commonHandlers.ThenFunc(ClientViewHandler))
		router.Get("/plus/*routes", commonHandlers.ThenFunc(ClientViewHandler))
		router.Get("/advert/*routes", commonHandlers.ThenFunc(ClientViewHandler))
		router.Get("/listing/*routes", commonHandlers.ThenFunc(ClientViewHandler))
		router.Get("/result/*routes", commonHandlers.ThenFunc(ClientViewHandler))

		router.Get("/cust", commonHandlers.ThenFunc(CustHandler))

		//api requests below


		router.Post("/login", commonHandlers.ThenFunc(web.Login))

		router.Post("/api/result", commonHandlers.ThenFunc(web.SearchHandler))
		router.Get("/api/listings", commonHandlers.ThenFunc(web.GetListHandler))
		router.Get("/api/getcatList", commonHandlers.ThenFunc(web.GetHandler))
		router.Get("/api/getsingle", commonHandlers.ThenFunc(web.getCatHandler))
		router.Get("/api/getsinglelist", commonHandlers.ThenFunc(web.getlistHandler))
		router.Get("/api/newview", commonHandlers.ThenFunc(web.GetNewView))
		router.Get("/api/falseview", commonHandlers.ThenFunc(web.FalseH))


		router.Get("/false", commonHandlers.ThenFunc(web.Fictionalcat))
		router.Get("/slider", commonHandlers.ThenFunc(web.SliderHandler))


		router.Get("/Upload", commonHandlers.ThenFunc(web.CsvHandler))
		router.Get("/fix", commonHandlers.ThenFunc(web.Fix))
		//forpayment
		router.Get("/newapp", commonHandlers.ThenFunc(web.PaymentAfter))
		router.Get("/napp", commonHandlers.ThenFunc(web.Post_Params))
		router.Get("/error", commonHandlers.ThenFunc(web.NoPaymentAfter))
	*/

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
