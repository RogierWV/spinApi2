package main

import (
	"time"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httputil"
	"log"
	"os"
	//"io"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/antage/eventsource.v1"
)

var conn *sql.DB
var es eventsource.EventSource

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

func GetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	row := conn.QueryRow("SELECT * FROM blog WHERE id = $1 LIMIT 1", ps.ByName("id"))
	data := BlogPost{}
	row.Scan(&data.Id, &data.Titel, &data.Text, &data.Auteur, &data.Img_url, &data.Ctime, &data.Image)
	buf,_ := json.Marshal(data)
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM spindata ORDER BY tijd DESC LIMIT 1")
	spin := SpinData{}
	rows.Next()
	rows.Scan(&spin.Id, &spin.Tijd, &spin.Mode, &spin.Hellingsgraad, &spin.Snelheid, &spin.Batterij, &spin.BallonCount)
	buf,_ := json.Marshal(spin)
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestSpinBatterij(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	row := conn.QueryRow("SELECT batterij FROM spindata ORDER BY tijd DESC LIMIT 1")
	var data int 
	row.Scan(&data)
	buf,_ := json.Marshal(data)
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestSpinMode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	row := conn.QueryRow("SELECT mode FROM spindata ORDER BY tijd DESC LIMIT 1")
	var data string 
	row.Scan(&data)
	buf,_ := json.Marshal(data)
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
	row := conn.QueryRow("SELECT * FROM servodata ORDER BY tijd DESC LIMIT 1")
	servo := ServoData{}
	row.Scan(&servo.Id, &servo.ServoId, &servo.Tijd, &servo.Voltage, &servo.Positie, &servo.Load, &servo.Temperatuur)
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

func GetLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows,_ := conn.Query("SELECT * FROM logs")
	data := []LogData{}
	for rows.Next() {
		log := LogData{}
		rows.Scan(&log.Id, &log.Log)
		data = append(data, log)
	}
	buf,_ := json.Marshal(data)
	SetHeaders(&w)
	w.Write(buf)
}

func GetLatestGyroData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	row := conn.QueryRow("SELECT hellingsgraad FROM spindata ORDER BY tijd DESC LIMIT 1")
	var helling int
	row.Scan(&helling)
	buf,_ := json.Marshal(helling)
	SetHeaders(&w)
	w.Write(buf)
}

func Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buf,_ := json.Marshal("test")
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
	_,err := conn.Query("INSERT INTO blog (titel, text, auteur, ctime) VALUES ($1, $2, $3, $4)", r.FormValue("onderwerp"), r.FormValue("bericht"), r.FormValue("naam"), time.Now())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(201)
	w.Write([]byte("<meta http-equiv=\"refresh\" content=\"1; url=http://knightspider.herokuapp.com/#/blog\">successful"))
}

func PostSpinData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// buf := make([]byte,100)
	// r.Body.Read(buf)
	// w.Write(buf)
	reqStr, _ := httputil.DumpRequest(r, true)
	w.Write(reqStr)
	r.ParseForm()
	_,err := conn.Query("INSERT INTO spindata (mode, hellingsgraad, snelheid, batterij, balloncount) VALUES ($1, $2, $3, $4, $5)", 
		//r.FormValue("mode"), r.FormValue("hellingsgraad"), r.FormValue("snelheid"), r.FormValue("batterij"), r.FormValue("ballonCount"))
		//"manueel", 0, 300, 50, 0)
		r.FormValue("mode"),
		r.FormValue("hellingsgraad"),
		300,50,
		//r.FormValue("batterij"),
		r.FormValue("ballonCount"))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}
	w.WriteHeader(201)
	//w.Write([]byte(fmt.Sprintf("mode = %s, hellingsgraad = %s, snelheid = %s, batterij = %s, balloncount = %s", mode, hellingsgraad, snelheid, batterij, balloncount)))
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

func PostLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_,err := conn.Query("INSERT INTO logs (log) VALUES ($1)", 
		r.FormValue("log"))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	es.SendEventMessage(r.FormValue("log"), "log", "")
	w.WriteHeader(201)
	w.Write([]byte(r.FormValue("log")))
}

func Head(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SetHeaders(&w)
	w.WriteHeader(204)
}

func main() {
	conn,_ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer conn.Close()

	es = eventsource.New(
		&eventsource.Settings{	
			Timeout: 5 * time.Second,
			CloseOnTimeout: false,
			IdleTimeout: 30 * time.Minute,
		},
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("X-Accel-Buffering: no"),
				[]byte("Access-Control-Allow-Origin: *"),
			}
		},
	)
	defer es.Close()

	router := httprouter.New()
	router.HEAD("/*path", Head)
	router.GET("/test", Test)
	router.GET("/blog", GetBlog)
	router.GET("/blog/:id", GetPost)
	router.GET("/spin/latest", GetLatestSpinData)
	router.GET("/spin/latest/batterij", GetLatestSpinBatterij)
	router.GET("/spin/latest/mode", GetLatestSpinMode)
	router.GET("/spin/latest/helling", GetLatestGyroData)
	router.GET("/spin/archive", GetArchivedSpinData)
	router.GET("/spin/archive/batterij", GetArchivedSpinBatterij)
	router.GET("/spin/archive/mode", GetArchivedSpinMode)
	router.GET("/servo/latest", GetLatestServoData)
	router.GET("/servo/archive", GetArchivedServoData)
	router.GET("/log", GetLogs)
	router.POST("/blog", PostBlog)
	router.POST("/spin", PostSpinData)
	router.POST("/servo", PostServoData)
	router.POST("/log", PostLog)

	http.Handle("/subscribe", es)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Starting server at localhost:%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}