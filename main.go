package main

import (
	"fmt"
	_ "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jemblonganvalley/controller"
)

func main(){

	//membuat server variable
	app := mux.NewRouter().StrictSlash(true)

	//grouping route
	api := app.PathPrefix("/api").Subrouter()

	// router 
	api.HandleFunc("/user_create", User_create).Methods("POST")

	//connection and migration
	UserMigration()

	fmt.Println("server berjalan di port 9000")

	//kita jalankan server
	log.Fatal(http.ListenAndServe(":9000", app))
	

}