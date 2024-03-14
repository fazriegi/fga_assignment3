package main

import (
	"fmt"
	"net/http"

	"github.com/fazriegi/fga_assignment3/controller"
	"github.com/fazriegi/fga_assignment3/helper"
	"github.com/fazriegi/fga_assignment3/model"
)

func main() {
	initData := model.Weather{
		Status: model.Status{
			Water: 0,
			Wind:  0,
		},
	}
	helper.WriteFile(initData, "weather.json")

	helper.StartCron()

	http.HandleFunc("/", controller.StatusController)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
