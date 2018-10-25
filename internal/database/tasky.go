package database

import (
	"github.com/globalsign/mgo"
)

type TaskyStorage struct {
	backendDB   *mgo.Database
	TaskIndices *mgo.Collection
	Tasks       *mgo.Collection
}

func NewTaskyStorage(dbName string) *TaskyStorage {
	backendDB := Manager.Create(dbName)

	return &TaskyStorage{
		backendDB:   backendDB,
		TaskIndices: backendDB.C("taskIndices"),
		Tasks:       backendDB.C("tasks"),
	}
}
