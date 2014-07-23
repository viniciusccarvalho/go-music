package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Album struct {
	Id          bson.ObjectId `json:"id"           bson:"_id"`
	Title       string        `json:"title"`
	Artist      string        `json:"artist"`
	ReleaseYear string        `json:"releaseYear"`
	Genre       string        `json:"genre"`
	TrackCount  int32         `json:"trackCount"`
	AlbumId     string        `json:"albumId"`
}

type ApplicationInfo struct {
	Profiles []string `json:"profiles"`
	Services []string `json:"services"`
}

func listAlbums(response http.ResponseWriter, request *http.Request) {
	session, err := openSession()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("").C("album")
	results := make([]Album, 10)
	c.Find(bson.M{}).All(&results)
	js, err := json.Marshal(results)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Write(js)
}

func addAlbum(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	album := Album{}
	err := decoder.Decode(&album)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if !album.Id.Valid() {
		album.Id = bson.NewObjectId()
	}
	session, err := openSession()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C("album")
	c.UpsertId(album.Id,&album)
	response.Write([]byte("{}"))
}

func deleteAlbum(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	session, err := openSession()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C("album")
	bid := bson.ObjectIdHex(id)
	c.RemoveId(bid)
}


func info(response http.ResponseWriter, request *http.Request) {
	serviceList := []string{}

	for service, _ := range env.services {
		serviceList = append(serviceList, service)
	}

	appInfo := ApplicationInfo{[]string{env.profile}, serviceList}
	js, _ := json.Marshal(appInfo)
	response.Write(js)
}

func openSession() (*mgo.Session, error) {
	service := env.uri("document")
	session, err := mgo.Dial(fmt.Sprintf("%v", service.Credentials["uri"]))

	return session, err
}
