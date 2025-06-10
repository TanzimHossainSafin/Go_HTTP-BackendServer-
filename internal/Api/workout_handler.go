package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)
type WorkoutHandler struct {

}

func NewworkoutHandler() *WorkoutHandler{
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) GetWorkoutById(w http.ResponseWriter,r *http.Request) {
	paramsWorkoutId:=chi.URLParam(r,"workoutId")
	if paramsWorkoutId == "" {
		http.Error(w, "Workout ID is required", http.StatusBadRequest)
		return
	}
	_, err := strconv.ParseInt(paramsWorkoutId, 10, 64)
	if err!=nil{
		http.NotFound(w, r)
		return
	}
    w.Write([]byte("GetWorkoutById called"))
}

func (wh *WorkoutHandler) CreateWorkouts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateWorkouts called\n")
}