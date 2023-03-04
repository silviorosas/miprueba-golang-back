package main

import (
	"api-go-4softwaredeveloper/routes"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
func main() {
	//commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonasRoutes(router)
	http.ListenAndServe(":3000", setHeaders(router))

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose desde puerto 3000")
	log.Println(server.ListenAndServe(), "ListenAndServe")

}
