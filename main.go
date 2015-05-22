package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var conn DB

func GetBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
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
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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