package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
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
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

//Conf nbfmjh
type Conf struct {
	xx string
	xy string
}

var (
	config Conf

	//IMAGE_DIR = "./"
)

func init() {

	MONGOSERVER := os.Getenv("MONGOLAB_URI")
	MONGODB := os.Getenv("MONGODB")
	if MONGOSERVER == "" {
		log.Println("No mongo server address set, resulting to default address")
		MONGOSERVER = "mongodb://localhost"
		MONGODB = "yellowListings"
		//mongodb://localhost
	}
	config = Conf{
		xy: MONGODB,
		xx: MONGOSERVER,
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	router := NewRouter()

	router.ServeFiles("/zohoverify/*filepath", http.Dir("assets"))
	router.Get("/admin", commonHandlers.ThenFunc(FrontAdminHandler))
	router.Get("/login", commonHandlers.ThenFunc(LoginAdmin))
	router.Get("/", commonHandlers.ThenFunc(ClientViewHandler))
	router.Get("/client", commonHandlers.ThenFunc(ClientAdmin))
	router.Get("/Newlisting", commonHandlers.ThenFunc(ClientIndex))
	//fs := http.FileServer(http.Dir("admin/assets/"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.Get("/addlistingtemp", commonHandlers.ThenFunc(AddListingViewHandler))
	router.Get("/addlisting", commonHandlers.ThenFunc(AddListingView))
	router.Get("/addcattemp", commonHandlers.ThenFunc(addcatViewHandler))
	router.Get("/viewlistingtemp", commonHandlers.ThenFunc(UnapprovedViewHandler))
	router.Get("/category", commonHandlers.ThenFunc(CategoryHandler))
	router.Get("/result", commonHandlers.ThenFunc(ResultHandler))
	router.Get("/listing", commonHandlers.ThenFunc(ListingHandler))
	router.Get("/home", commonHandlers.ThenFunc(HomeHandler))
	router.Get("/newad", commonHandlers.ThenFunc(NewAdvertHandler))
	router.Get("/cust", commonHandlers.ThenFunc(CustHandler))
	router.Get("/client/update", commonHandlers.ThenFunc(TimeUpdatehandler))
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.ServeFiles("/temp/*filepath", http.Dir("admin/partials"))
	router.ServeFiles("/cust/partials/*filepath", http.Dir("cust/partials"))
	router.ServeFiles("/cust/js/*filepath", http.Dir("cust/js"))
	//api requests below
	router.Post("/api/addcat", commonHandlers.ThenFunc(addCatHandler))
	router.Get("/api/getcat", commonHandlers.ThenFunc(getcatHandler))
	router.Post("/api/newuser", commonHandlers.ThenFunc(NewUserHandler))
	router.Post("/api/addlisting", commonHandlers.ThenFunc(AddHandler))
	router.Post("/api/approve", commonHandlers.ThenFunc(Approvehandler))
	router.Post("/login", commonHandlers.ThenFunc(Login))
	router.Post("/adminlogin", commonHandlers.ThenFunc(AdminLogin))
	router.Get("/api/unapproved", commonHandlers.ThenFunc(GetunapprovedHandler))
	router.Post("/api/newAd", commonHandlers.ThenFunc(NewAdHandler))
	router.Post("/api/result", commonHandlers.ThenFunc(SearchHandler))
	router.Get("/api/listings", commonHandlers.ThenFunc(GetListHandler))
	router.Get("/api/getcatList", commonHandlers.ThenFunc(GetHandler))
	router.Get("/api/getsingle", commonHandlers.ThenFunc(getCatHandler))
	router.Get("/api/getsinglelist", commonHandlers.ThenFunc(getlistHandler))
	router.Get("/api/newview", commonHandlers.ThenFunc(GetNewView))
	router.Get("/api/falseview", commonHandlers.ThenFunc(FalseH))
	router.Get("/api/plus", commonHandlers.ThenFunc(GetPlusPayHandler))
	router.Get("/api/adminList", commonHandlers.ThenFunc(GetAdminsHandler))
	router.Get("/false", commonHandlers.ThenFunc(Fictionalcat))

	router.Get("/advert", commonHandlers.ThenFunc(FalseA))
	router.Get("/Upload", commonHandlers.ThenFunc(CsvHandler))
	router.Get("/fix", commonHandlers.ThenFunc(Fix))
	//forpayment
	router.Get("/newapp", commonHandlers.ThenFunc(PaymentAfter))
	router.Post("/newapp", commonHandlers.ThenFunc(Post_Params))
	router.Get("/error", commonHandlers.ThenFunc(NoPaymentAfter))

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
