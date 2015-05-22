package main

import (
	"time"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var conn *sql.DB

type BlogPost struct {
	id int64
	titel string
	text string
	auteur string
	img_url string
	ctime time.Time
	image string
}

func GetBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM blog")
	data := []BlogPost{}
	for rows.Next() {
		post := BlogPost{}
		rows.Scan(&post.id, &post.titel, &post.text, &post.auteur, &post.img_url, &post.ctime, &post.image)
		data = append(data, post)
	}
	buf,_ := json.Marshal(data)
	w.Write(buf)
}

func GetSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func GetServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func PostBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func PostSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func PostServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func main() {
	conn,_ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer conn.Close()
	router := httprouter.New()
	router.GET("/blog", GetBlog)
	router.GET("/spin", GetSpinData)
	router.GET("/servo", GetServoData)
	router.POST("/blog", PostBlog)
	router.POST("/spin", PostSpinData)
	router.POST("/servo", PostServoData)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Starting server at localhost:%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}