package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"os"
	"fmt"
)

type Environment struct {
	services map[string][]ServiceDefinition
	profile  string
}

type ServiceDefinition struct {
	Name        string                 `json:"name"`
	Label       string                 `json:"label"`
	Tags        []string               `json:"tags"`
	Plan        string                 `json:"plan"`
	Credentials map[string]interface{} `json:"credentials"`
}

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

func NewEnvironment() *Environment {
	env := Environment{}
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &env.services)
	if err != nil {
		fmt.Errorf("error decoding string", err)
	}
	vcap := os.Getenv("VCAP_APPLICATION")
	if len(vcap) == 0 {
		env.profile = "local"
	} else {
		env.profile = "cloud"
	}
	return &env
}

func (this Environment) uri(serviceTag string) ServiceDefinition {
	if this.profile == "local" {
		service := ServiceDefinition{}
		service.Credentials["uri"] = "localhost"
		return service
	}
	for key, _ := range this.services {
		for _, service := range this.services[key] {
			return service
		}
	}
	return ServiceDefinition{}
}
