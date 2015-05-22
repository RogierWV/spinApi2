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
	Id int64 `json:"id"`
	Titel string `json:"titel"`
	Text string `json:"text"`
	Auteur string `json:"auteur"`
	Img_url string `json:"img_url"`
	Ctime time.Time `json:"ctime"`
	Image string `json:"image"`
}

type SpinData struct {
	Id int64 `json:"id"`
	Tijd time.Time `json:"tijd"`
	Mode string `json:"mode"`
	Hellingsgraad int64 `json:"hellingsgraad"`
	Snelheid int64 `json:"snelheid"`
	Batterij int64 `json:"batterij"`
	BallonCount int64 `json:"ballonCount"`
}

type ServoData struct {
	Id int64 `json:"id"`
	ServoId int64 `json:"servo_id"`
	Tijd time.Time `json:"tijd"`
	Voltage int64 `json:"voltage"`
	Positie int64 `json:"positie"`
	Load int64 `json:"load"`
	Temperatuur int64 `json:"temperatuur"`
}

func GetBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM blog")
	data := []BlogPost{}
	for rows.Next() {
		post := BlogPost{}
		rows.Scan(&post.Id, &post.Titel, &post.Text, &post.Auteur, &post.Img_url, &post.Ctime, &post.Image)
		data = append(data, post)
	}
	buf,_ := json.Marshal(data)
	w.Write(buf)
}

func GetLatestSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM spindata ORDER BY tijd DESC LIMIT 1")
	spin := SpinData{}
	rows.Scan(&spin.Id, &spin.Tijd, &spin.Mode, &spin.Hellingsgraad, &spin.Snelheid, &spin.Batterij, &spin.BallonCount)
	buf,_ := json.Marshal(spin)
	w.Write(buf)
}

func GetArchivedSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM spindata")
	data := []SpinData{}
	for rows.Next() {
		spin := SpinData{}
		rows.Scan(&spin.Id, &spin.Tijd, &spin.Mode, &spin.Hellingsgraad, &spin.Snelheid, &spin.Batterij, &spin.BallonCount)
		data = append(data, spin)
	}
	buf,_ := json.Marshal(data)
	w.Write(buf)
}

func GetLatestServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM servodata ORDER BY tijd DESC LIMIT 1")
	servo := ServoData{}
	rows.Scan(&servo.Id, &servo.ServoId, &servo.Tijd, &servo.Voltage, &servo.Positie, &servo.Load, &servo.Temperatuur)
	buf,_ := json.Marshal(servo)
	w.Write(buf)
}

func GetArchivedServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM servodata")
	data := []ServoData{}
	for rows.Next() {
		servo := ServoData{}
		rows.Scan(&servo.Id, &servo.ServoId, &servo.Tijd, &servo.Voltage, &servo.Positie, &servo.Load, &servo.Temperatuur)
		data = append(data, servo)
	}
	buf,_ := json.Marshal(data)
	w.Write(buf)
}

func PostBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	/*_,err := db.Query("INSERT INTO blog (titel, text, auteur, ctime, image) VALUES ($1, $2, $3, $4, $5)", )
	if err != nil {
		w.WriteHead(500)
		w.Write([]byte("error posting"))
		return
	}
	w.WriteHead(201)
	w.Write([]byte("successful"))*/
}

func PostSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func PostServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data,_ := json.Marshal("blah")
	w.Write(data)
}

func TeaPot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Error(w, http.StatusText(418), 418)
}

func main() {
	conn,_ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer conn.Close()
	router := httprouter.New()
	router.GET("/blog", GetBlog)
	router.GET("/spin/latest", GetLatestSpinData)
	router.GET("/spin/archive", GetArchivedSpinData)
	router.GET("/servo/latest", GetLatestServoData)
	router.GET("/servo/archive", GetArchivedServoData)
	router.GET("/teapot", TeaPot)
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