package main

import (
	"fmt"
	"image/color"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500,500))

	//api

	res, err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q=noida&appid=6b4711303141575cb98b2daddbcf0a21");
	if err!=nil{
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err)
	}

	//  use "https://quicktype.io/" to get structure of json data for go

	weather,err:=UnmarshalWelcome(body)
	if err!=nil{
		fmt.Println(err)
	}

	img:= canvas.NewImageFromFile("weather.jpg")
	img.FillMode = canvas.ImageFillOriginal

	label1:=canvas.NewText("weather Details",color.White)
	label1.TextStyle =fyne.TextStyle{Bold: true}

	label2:= canvas.NewText(fmt.Sprintf("Country %s",weather.Sys.Country),color.White)
	label3:= canvas.NewText(fmt.Sprintf("Wind Speed %.2f",weather.Wind.Speed),color.White)
	label4:= canvas.NewText(fmt.Sprintf("Temperature %2f",weather.Main.Temp),color.White)

	w.SetContent(
		container.NewVBox(
			label1,
			img,
			label2,
			label3,
			label4,
		),
	)
	w.ShowAndRun()
}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}


