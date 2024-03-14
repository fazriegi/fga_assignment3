package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"text/template"

	"github.com/fazriegi/fga_assignment3/helper"
	"github.com/fazriegi/fga_assignment3/model"
)

func StatusController(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("public", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var weather model.Weather
	data := helper.ReadFile("weather.json")

	if err = json.Unmarshal(data, &weather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, weather.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
