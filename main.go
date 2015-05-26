package main

import (
	"time"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"os"
	//"io"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var conn *sql.DB

func SetHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
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
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM spindata ORDER BY tijd DESC LIMIT 1")
	spin := SpinData{}
	rows.Scan(&spin.Id, &spin.Tijd, &spin.Mode, &spin.Hellingsgraad, &spin.Snelheid, &spin.Batterij, &spin.BallonCount)
	buf,_ := json.Marshal(spin)
	SetHeaders(&w)
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
	SetHeaders(&w)
	w.Write(buf)
}

func GetArchivedSpinBatterij(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT batterij FROM spindata")
	data := make([]int, 0)
	var scanInt int
	for rows.Next() {
		rows.Scan(&scanInt)
		data = append(data, scanInt)
	}
	buf,_ := json.Marshal(data)
	fmt.Printf(string(buf))
	SetHeaders(&w)
	w.Write(buf)
}

func GetArchivedSpinMode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT mode FROM spindata")
	data := make([]string, 0)
	var scanStr string
	for rows.Next() {
		rows.Scan(&scanStr)
		data = append(data, scanStr)
	}
	buf,_ := json.Marshal(data)
	fmt.Printf(string(buf))
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM servodata ORDER BY tijd DESC LIMIT 1")
	servo := ServoData{}
	rows.Scan(&servo.Id, &servo.ServoId, &servo.Tijd, &servo.Voltage, &servo.Positie, &servo.Load, &servo.Temperatuur)
	buf,_ := json.Marshal(servo)
	SetHeaders(&w)
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
	SetHeaders(&w)
	w.Write(buf)
}

func PostBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//r.ParseMultipartForm(32 << 20)
	/*file, handler, err := r.FormFile("uploadfile")
	defer file.Close()
	if err == nil {
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

	err = nil*/

	//_,err := conn.Query("INSERT INTO blog (titel, text, auteur, ctime, image) VALUES ($1, $2, $3, $4, $5)", r.FormValue("titel"), r.FormValue("text"), r.FormValue("auteur"), time.Now(), "http://idp-api.herokuapp.com/img/"+handler.Filename)
	_,err := conn.Query("INSERT INTO blog (titel, text, auteur, ctime) VALUES ($1, $2, $3, $4)", r.FormValue("titel"), r.FormValue("text"), r.FormValue("auteur"), time.Now())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("successful"))
}

func PostSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_,err := conn.Query("INSERT INTO spindata (tijd, mode, hellingsgraad, snelheid, batterij, balloncount) VALUES ($1, $2, $3, $4, $5, $6)", time.Now(), 
		r.FormValue("mode"), r.FormValue("hellingsgraad"), r.FormValue("snelheid"), r.FormValue("batterij"), r.FormValue("ballonCount"))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("successful"))
}

func PostServoData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_,err := conn.Query("INSERT INTO servodata (servo_id, tijd, voltage, positie, load, temperatuur) VALUES ($1, $2, $3, $4, $5, $6)", 
		r.FormValue("servo_id"), time.Now(), r.FormValue("voltage"), r.FormValue("positie"), r.FormValue("load"), r.FormValue("Temperatuur"))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("successful"))
}

/*func GetDoc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w,r,"./static/doc.html")
}*/

func Static(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w,r,"./static/"+ps.ByName("file"))
}

func main() {
	conn,_ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer conn.Close()
	router := httprouter.New()
	//router.GET("/", GetDoc)
	router.GET("/blog", GetBlog)
	router.GET("/spin/latest", GetLatestSpinData)
	router.GET("/spin/archive", GetArchivedSpinData)
	router.GET("/spin/archive/batterij", GetArchivedSpinBatterij)
	router.GET("/spin/archive/mode", GetArchivedSpinMode)
	router.GET("/servo/latest", GetLatestServoData)
	router.GET("/servo/archive", GetArchivedServoData)
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