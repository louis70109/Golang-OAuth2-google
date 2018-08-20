package Routes

import (
	"net/http"
	"time"
	"log"
	"strings"
	"github.com/gorilla/mux"
	"oauth-github/Controller"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}



type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/*
this method is start listen Server
*/
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router

}

/*
	It is define router method
	2018/7/17 we must define CRUD function in author
*/
var routes = Routes{
	Route{
		"Callback",
		strings.ToUpper("POST"),
		"/callback",
		Controller.Callback,
	},
}
