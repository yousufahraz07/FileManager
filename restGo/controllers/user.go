package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yousufahraz07/restGo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type UserController struct{
	session *mgo.Session
}
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController)GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id = p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectHex(id)
	
	u:= models.User{}

	if err := uc.session.DB("mongo-golang").C("Users").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}

	uj,err := json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

w.Header().Set("Content-Type","application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w,"%s\n", uj)

}
func (uc UserController) CreateUser(w htttp.ResponseWriter,r *http.Request, _httprouter.Params){

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj,err := json.Marshal(u)

	if err!= nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w,"%s\n", uj)
}
func (uc UserController) DeleteUser(w httprouter, r *http.Request, p httprouter.Params){

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid:= bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid);{
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}