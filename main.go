package main

import "net/http"
// import "./handlers"
// import "./storage"
import handlers "github.com/gaurav-225/BuildingAndLeveraging-REST-APIs-withGo/handlers"
import storage "github.com/gaurav-225/BuildingAndLeveraging-REST-APIs-withGo/storage"

func main() {
	println("Hello, World!")

	db := storage.NewInMemoryDB()

	mux := http.NewServeMux()

	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))



	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println("Error starting server", err)
		return
	}
}