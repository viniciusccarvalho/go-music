package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func listAlbums(response http.ResponseWriter, request *http.Request) {
	repo := Repo{"album"}
	results := make([]Album, 10)
	repo.All(&results)
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
	repo := Repo{"album"}
	repo.Upsert(album.Id,&album)
	
	response.Write([]byte("{}"))
}

func deleteAlbum(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	repo := Repo{"album"}
	bid := bson.ObjectIdHex(id)
	repo.Delete(bid)
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

