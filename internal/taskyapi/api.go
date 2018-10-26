package taskyapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/database"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
)

var DB *database.TaskyStorage

type task struct {
	Index       int    `json:"index" bson:"index"`
	Title       string `json:"title" bson:"title"`
	Time        string `json:"time" bson:"time"`
	Description string `json:"description" bson:"description"`
	Username    string `json:"username" bson:"username"`

	// Only used for frontend
	Action string `json:"action" bson:"action"`
}

type taskIndices struct {
	Username           string `json:"username" bson:"username"`
	MissingTaskIndices []int  `json:"missingIndices" bson:"missingIndices"`
	SortedTaskIndices  []int  `json:"sortedIndices" bson:"sortedIndices"`
}

func Setup() {
	appName := viper.GetString(config.AppNameKey)
	DB = database.NewTaskyStorage(appName)
	decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
}

func RegisterRoutes(mux chi.Router) {
	mux.Get("/tasky/get_tasks", getTasks)
	mux.Get("/tasky/get_missing_task_indices", getMissingTaskIndices)

	mux.Post("/tasky/create_task", createTask)
	mux.Post("/tasky/remove_task", removeTask)
	mux.Post("/tasky/update_task", updateTaskContent)
	mux.Post("/tasky/update_task_pos", updateTaskPosition)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
}

func getMissingTaskIndices(w http.ResponseWriter, r *http.Request) {
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if badRequest(w, err) {
		return
	}

	var t task
	err = decoder.Decode(&t, r.Form)
	if badRequest(w, err) {
		return
	}
	err = DB.Tasks.Insert(t)
	if badRequest(w, err) {
		return
	}

	var ti taskIndices
	err = DB.TaskIndices.Find(bson.M{"username": t.Username}).One(&ti)
	if badRequest(w, err) {
		return
	}
	removeTaskIndexIfExists(&ti.MissingTaskIndices, t.Index)
	insertTaskIndex(&ti.SortedTaskIndices, 0, t.Index)
	if len(ti.MissingTaskIndices) == 0 {
		insertTaskIndexSorted(&ti.MissingTaskIndices, len(ti.SortedTaskIndices))
	}
	err = DB.TaskIndices.Update(
		bson.M{"username": t.Username},
		bson.M{"$set": bson.M{"missingIndices": ti.MissingTaskIndices, "sortedIndices": ti.SortedTaskIndices}},
	)
	if badRequest(w, err) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateTaskContent(w http.ResponseWriter, r *http.Request) {
	// Update database with new content of task
}

func updateTaskPosition(w http.ResponseWriter, r *http.Request) {
	// Update database with new position of task
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if badRequest(w, err) {
		return
	}

	var t task
	err = decoder.Decode(&t, r.Form)
	if badRequest(w, err) {
		return
	}
	err = DB.Tasks.Remove(t)
	if badRequest(w, err) {
		return
	}

	var ti taskIndices
	err = DB.TaskIndices.Find(bson.M{"username": t.Username}).One(&ti)
	if badRequest(w, err) {
		return
	}
	insertTaskIndexSorted(&ti.MissingTaskIndices, t.Index)
	removeTaskIndexIfExists(&ti.SortedTaskIndices, t.Index)
	err = DB.TaskIndices.Update(
		bson.M{"username": t.Username},
		bson.M{"$set": bson.M{"missingIndices": ti.MissingTaskIndices, "sortedIndices": ti.SortedTaskIndices}},
	)
	if badRequest(w, err) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func removeTaskIndexIfExists(ti *[]int, indexToRemove int) {
	for index, value := range *ti {
		if value == indexToRemove {
			*ti = append((*ti)[:index], (*ti)[index+1:]...)
		}
	}
}

func insertTaskIndex(ti *[]int, indexToInsertAt int, indexToInsert int) {
	*ti = append((*ti)[:indexToInsertAt], append([]int{indexToInsert}, (*ti)[indexToInsertAt:]...)...)
}

func insertTaskIndexSorted(ti *[]int, indexToInsert int) {
	for index, value := range *ti {
		if value > indexToInsert {
			insertTaskIndex(ti, index, indexToInsert)
			return
		}
	}
	insertTaskIndex(ti, len(*ti), indexToInsert)
}

func badRequest(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return true
	}
	return false
}
