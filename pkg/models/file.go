package models

import (
	"filemanager/pkg/config"
	"github.com/jinzhu/gorm"
)

var db  *gorm.DB

type File struct{
	gorm.Model
	Name string `gorm: ""json :"name"`
}

func init(){
	config.Connect()
	db= config.GetDB()
	db.AutoMigrate(&File{})
}

func( f *File) CreateFile() *File{
	db.NewRecord(f)
	db.Create(&f)
	return f
}

func GetAllFiles() []File{
	var Files []File
	db.Find(&Files)
	return Files
}

func GetFileById(Id int64)(*File, *gorm.DB){
	var getFile File
	db := db.Where("ID=?", Id).Find(&getFile)
	return &getFile,db
}

func DeleteFile(ID int64) File{
	var file  File
	db.Where("ID=?",ID).Delete(file)
	return file
}