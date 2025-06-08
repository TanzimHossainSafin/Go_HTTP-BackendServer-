package routes

import (
	"github.com/tanzim_safin/femProject/internal/app"
	"github.com/go-chi/chi"
	
)
func SetUpRoutes(app *app.Application) *chi.Mux {
	r:=chi.NewRouter()
	r.Get("/healthcheck", app.HealthCheck)
	r.Get("/workouts/{workoutId}", app.WorkoutHandler.GetWorkoutById)
	r.Post("/workouts", app.WorkoutHandler.CreateWorkouts)
	return r
}