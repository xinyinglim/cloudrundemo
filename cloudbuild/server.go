package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

var a App
var helloworldMessage = "Hello World"

//if checkmessage does not equal to message, build will fail
var helloworldCheckMessage = "Hello World"

func main() {
	a := App{}
	a.Initialize()
	addr := os.Getenv("PORT")
	fmt.Println(addr)
	if addr == "" {
		addr = "8081"
	}
	a.Run(fmt.Sprintf(":%s", addr))
}

type App struct {
	Router *mux.Router
	// DB     *sql.DB //if you require database
}

//exposes references to the router and database

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/helloworld", a.helloWorld).Methods("GET")
}

//Run where addr is ":8081"
func (a *App) Run(addr string) {
	fmt.Printf("Run at %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

//handler
func (a *App) helloWorld(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r) //returns route variables for the current variable
	// id, err := strconv.Atoi(vars["id"])
	respondWithString(w, http.StatusOK, helloworldMessage)
}

func respondWithString(w http.ResponseWriter, code int, message string) {
	response := []byte(message)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write(response)
}
