package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	 _"github.com/jinzhu/gorm/dialects/mysql"
 "filemanager/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterFileManagerRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}
