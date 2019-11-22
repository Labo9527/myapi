package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct{
	Id int
	Username string
	Password string
	Mail string
}

type Presentuser struct{
	Id int
	Username string
	Mail string
}

type Article struct{
	Id int
	Title string
	Date string
	ReadnNum int
	Context string
	Author User
}

var users []User
var articles []Article

func QueryUser(w http.ResponseWriter, r *http.Request) {
	for _, user := range users{
		var mypre Presentuser
		mypre.Username=user.Username
		mypre.Mail=user.Mail
		mypre.Id=user.Id
		json.NewEncoder(w).Encode(mypre)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var myuser User
	_ = json.NewDecoder(r.Body).Decode(&myuser)
	if myuser.Username==""||myuser.Password==""||myuser.Mail=="" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Please Enter username, password and mail")
		return
	}
	myuser.Id=len(users)+1
	users = append(users, myuser)
	json.NewEncoder(w).Encode(myuser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	inputs := mux.Vars(r)
	for i, user := range users {
		if strconv.Itoa(user.Id) == inputs["id"] && user.Password==inputs["password"] {
			copy(users[i:], users[i+1:])
			users = users[:len(users)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func CreateArticle(w http.ResponseWriter, r *http.Request){
	inputs := mux.Vars(r)
	var nowUser User
	var find bool
	for _,user := range users{
		if strconv.Itoa(user.Id) == inputs["id"] && user.Password== inputs["password"] {
			nowUser=user
			find=true
			break
		}
	}
	if find==false{
		fmt.Fprintf(w, "No such user")
		return
	}
	var myarticle Article
	_ = json.NewDecoder(r.Body).Decode(&myarticle)
	if myarticle.Title=="" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Please Enter title")
		return
	}
	myarticle.Date=time.Now().Format("2006-01-02 15:04:05")
	myarticle.ReadnNum=0
	myarticle.Author=nowUser
	myarticle.Id=len(articles)+1
	articles = append(articles, myarticle)
	json.NewEncoder(w).Encode(myarticle)
}

func QueryArticle(w http.ResponseWriter, r *http.Request){
	for _, article := range articles{
		article.Id=article.Id+1
		json.NewEncoder(w).Encode(article)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	inputs := mux.Vars(r)
	var myid int
	myid, _ = strconv.Atoi(inputs["id"])
	for i,article := range articles{
		if i == myid - 1{
			if inputs["password"] != article.Author.Password{
				w.WriteHeader(403)
				fmt.Fprintf(w, "Error Password!")
			}
			copy(articles[i:], articles[i+1:])
			articles = articles[:len(articles)-1]
			json.NewEncoder(w).Encode(articles)
			return
		}
	}
}

func QueryAuser(w http.ResponseWriter, r *http.Request){
	inputs := mux.Vars(r)
	fmt.Println(inputs["id"],inputs["xx"])
	for _, user := range users{
		if strconv.Itoa(user.Id)!=inputs["id"] {
			continue
		}
		var mypre Presentuser
		mypre.Username=user.Username
		mypre.Mail=user.Mail
		mypre.Id=user.Id
		json.NewEncoder(w).Encode(mypre)
		break
	}
}

func QueryAuserarticle(w http.ResponseWriter, r *http.Request){
	inputs := mux.Vars(r)
	for _, user := range users{
		if strconv.Itoa(user.Id)!=inputs["id"] {
			continue
		}
		for _, article := range articles{
			if article.Author!=user{
				continue
			}
			article.Id=article.Id+1
			json.NewEncoder(w).Encode(article)
		}
		break
	}
}

func main(){

	mymux:=mux.NewRouter()
	mymux.HandleFunc("/users", QueryUser).Methods("GET")
	mymux.HandleFunc("/users/{id}", QueryAuser).Methods("GET")
	mymux.HandleFunc("/users/{id}/articles", QueryAuserarticle).Methods("GET")
	mymux.HandleFunc("/users", CreateUser).Methods("POST")
	mymux.HandleFunc("/users/{id}-{password}", DeleteUser).Methods("DELETE")
	mymux.HandleFunc("/articles", QueryArticle).Methods("GET")
	mymux.HandleFunc("/articles/{id}-{password}", CreateArticle).Methods("POST")
	mymux.HandleFunc("/articles/{id}-{password}", DeleteArticle).Methods("DELETE")


	fmt.Println("Listning to port 9090")
	err := http.ListenAndServe(":9090", mymux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}