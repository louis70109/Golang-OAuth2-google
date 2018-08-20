package main


import (
	"oauth-github/Routes"
	"flag"
	"fmt"
	"log"
	"net/http"

)

var (
	hostname string
	port     int
)

func init() {
	flag.StringVar(&hostname, "hostname", "0.0.0.0", "The hostname or IP on which the REST server will listen")
	flag.IntVar(&port, "port", 3001, "The port on which the REST server will listen")
}

func main() {

	flag.Parse()
	var address = fmt.Sprintf("%s:%d", hostname, port) //組合字串
	log.Println("REST service listening on", address)

	// register router
	router := Routes.NewRouter()
	// start server listening
	web := http.ListenAndServe(address, router)

	if web != nil {
		log.Fatalln("ListenAndServe err:", web)
	}
	log.Println("Server end")
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}
