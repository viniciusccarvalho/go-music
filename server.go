package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	env            = NewEnvironment()
	clusterSession *mgo.Session
)

func init() {
	service := env.uri("document")
	clusterSession, err := mgo.Dial(fmt.Sprintf("%v", service.Credentials["uri"]))

	session := clusterSession.Copy()

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("").C("album")

	total, _ := c.Count()
	fmt.Printf("Found %v documents\n", total)
	if total == 0 {
		fmt.Println("Starting DB ...")
		data, err := ioutil.ReadFile("albums.json")
		if err != nil {
			fmt.Print("Could not open albums.json")
		}
		albums := make([]Album, 10)
		err = json.Unmarshal([]byte(data), &albums)
		for _, album := range albums {
			album.Id = bson.NewObjectId()
			err = c.Insert(&album)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	}).Methods("GET")
	router.HandleFunc("/albums", listAlbums).Methods("GET")
	router.HandleFunc("/albums", addAlbum).Methods("POST", "PUT")
	router.HandleFunc("/info", info).Methods("GET")
	router.HandleFunc("/albums/{id}", deleteAlbum).Methods("DELETE")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}

	fmt.Println("Starting Go server on port ", port)

	http.Handle("/", router)
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":"+port, nil)
}



