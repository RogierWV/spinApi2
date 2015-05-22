package main

import (
	"time"
)

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

type Batterij struct {
	data []int64 `json:"data"`
}