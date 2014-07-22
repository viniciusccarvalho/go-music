package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	)

type Album struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear string `json:"releaseYear"`
	Genre       string `json:"genre"`
	TrackCount  int32  `json:"trackCount"`
	AlbumId     string `json:"albumId"`
}

type ApplicationInfo struct {
	Profiles []string `json:"profiles"`
	Services []string `json:"services"`
}


func listAlbums(response http.ResponseWriter, request *http.Request) {
	js, err := ioutil.ReadFile("albums.json")
	if err != nil {
		
	}
	response.Write(js)	
}

func info(response http.ResponseWriter, request *http.Request) {
	appInfo := ApplicationInfo{[]string{"cloud"},[]string{""}}
	js, _ := json.Marshal(appInfo)
	response.Write(js)
}