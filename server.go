package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting go server ...")
	router := mux.NewRouter()
	
	router.HandleFunc("/echo",echo)
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w,r,"./static/index.html")
			}).Methods("GET")
	router.HandleFunc("/albums",listAlbums).Methods("GET")
	router.HandleFunc("/info",info).Methods("GET")
	port := os.Getenv("PORT")
	fmt.Println("Using port ", port) 
	
	http.Handle("/", router)
	http.Handle("/static/",http.FileServer(http.Dir("./")))
	http.ListenAndServe(":"+port, nil)
}

func echo(response http.ResponseWriter, request *http.Request) {

	fmt.Fprint(response, "{\"message\": \"Hello World!\"}")
}



