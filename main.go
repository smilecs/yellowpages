package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
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
}

var (
	config Conf
)

func init() {
	config = Conf{
		xx: "mongodb://localhost",
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	router := NewRouter()
	router.Get("/admin", commonHandlers.ThenFunc(FrontAdminHandler))
	//fs := http.FileServer(http.Dir("admin/assets/"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.Get("/addlistingtemp", commonHandlers.ThenFunc(AddListingViewHandler))
	router.ServeFiles("/assets/*filepath", http.Dir("admin/assets"))
	router.Post("/api/addlisting", commonHandlers.ThenFunc(AddHandler))
	router.Post("/api/approve", commonHandlers.ThenFunc(Approvehandler))
	router.Get("/api/unapproved", commonHandlers.ThenFunc(GetunapprovedHandler))
	log.Println(config.xx)
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
