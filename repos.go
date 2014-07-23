package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)

type Repo struct {
	Collection string
}

type Persister interface {
	All(entity interface{}) error
	FindById(id interface{}, entity interface{}) error
	Upsert(id interface{}, entity interface{}) error
	Delete(id interface{}) error
}


func (repo Repo) All(entity interface{}) error{
	session := clusterSession.Clone()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("").C(repo.Collection)
	err := c.Find(bson.M{}).All(entity)
	return err
}

func (repo Repo) FindById(id interface{}, entity interface{}) error{
	session := clusterSession.Clone()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C(repo.Collection)
	err := c.FindId(id).One(entity)
	return err
}

func (repo Repo) Upsert(id interface{}, entity interface{}) error {
	session := clusterSession.Clone()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C(repo.Collection)
	_, err := c.UpsertId(id,entity)
	return err
}

func (repo Repo) Delete(id interface{}) error {
	session := clusterSession.Clone()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C(repo.Collection)
	err:= c.RemoveId(id)
	return err
}