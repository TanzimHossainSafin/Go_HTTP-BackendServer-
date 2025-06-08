package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/tanzim_safin/femProject/internal/app"
	"github.com/tanzim_safin/femProject/internal/routes"
)
func main(){
  var port int
  flag.IntVar(&port, "port", 8080, "Port to run the application on")
  flag.Parse()
  app,err:=app.NewApplication()
  if err!=nil{
	panic(err)
  }
  r := routes.SetUpRoutes(app)
  server:=&http.Server{
    Addr: fmt.Sprintf(":%d", port),
    Handler: r,
    IdleTimeout: time.Minute,
    WriteTimeout: 30*time.Second,
    ReadTimeout: 30*time.Second,
  }
   app.Logger.Printf("We are running our app on port %d", port)
  err=server.ListenAndServe()
  if err!=nil{
    app.Logger.Fatal(err)
  }
}