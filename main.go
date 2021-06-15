package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Welcome, Faizan..!!"
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  426,
		Height: 240,
		Title:  "Go-React App",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
