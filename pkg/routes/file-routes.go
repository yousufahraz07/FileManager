package routes

import (
	"github.com/gorilla/mux"
	"filemanager/pkg/controllers"
)

var RegisterFileManagerRoutes = func(router *mux.Router){
	router.HandleFunc("/file/", controllers.CreateFile).Methods("POST")
	router.HandleFunc("/file/", controllers.GetFile).Methods("GET")
	router.HandleFunc("/file/{fileId}", controllers.GetFileById).Methods("GET")
	router.HandleFunc("/file/{fileId}", controllers.UpdateFile).Methods("PUT")
	router.HandleFunc("/file/{fileId}", controllers.DeleteFile).Methods("DELETE")
}