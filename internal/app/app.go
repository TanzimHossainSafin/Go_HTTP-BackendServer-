package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

    "github.com/tanzim_safin/femProject/internal/Api"
)
type Application struct{
    Logger *log.Logger
    WorkoutHandler *api.WorkoutHandler
}

func NewApplication()(*Application,error){
    logger:=log.New(os.Stdout,"",log.Ldate|log.Ltime)
    workoutHandler:=api.NewworkoutHandler()
    app:=&Application{
        Logger:logger,
        WorkoutHandler:workoutHandler,
    }
    return app,nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w,"the Health Check is OK")
}