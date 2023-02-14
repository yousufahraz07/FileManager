package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"filemanager/pkg/utils"
	"filemanager/pkg/models"
)

var NewFile models.File

func GetFile(w http.ResponseWriter, r *http.Request){
	newFile:= models.GetAllFiles()
	res,_:=json.Marshal(newFile)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFileById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fileID:= vars["fileId"]
	ID,err:= strconv.ParseInt(fileID,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	fileDetails, _ := models.GetFileById(ID)
	res,  _ := json.Marshal(fileDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) 
}

func CreateFile(w http.ResponseWriter, r *http.Request){
	CreateFile := &models.File{}
	utils.ParseBody(r,CreateFile)
	b:= CreateFile.CreateFile()
	res, _ :=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteFile(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fileId := vars["fileId"]
	ID, err := strconv.ParseInt(fileId,0,0)
	if err!= nil{
		fmt.Println("error while parsing")
	}	
	file := models.DeleteFile(ID)
	res, _ := json.Marshal(file)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFile(w http.ResponseWriter, r *http.Request){
	var UpdateFile = &models.File{}
	utils.ParseBody(r, UpdateFile)
	vars := mux.Vars(r)
	fileID := vars["fileID"]
	ID,err := strconv.ParseInt(fileID,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	fileDetails, db:= models.GetFileById(ID)
	if UpdateFile.Name!= ""{
		fileDetails.Name = UpdateFile.Name
	}
	db.Save(&fileDetails)
	res, _ := json.Marshal(fileDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}