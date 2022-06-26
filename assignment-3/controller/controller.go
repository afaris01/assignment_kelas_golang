package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"assignment-3/model"
)

var water int = 10
var wind int = 5

func DataBase() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100

	water = rand.Intn(max-min+1) + min
	wind = rand.Intn(max-min+1) + min

	fmt.Println("Water : ", water)
	fmt.Println("Wind  : ", wind)

	statusJSON := &model.Status{Water: water, Wind: wind}
	JSONOutput := &model.Output{Status: *statusJSON}
	b, _ := json.MarshalIndent(JSONOutput, "", "   ")
	_ = ioutil.WriteFile("./static/file.json", b, 0644)

}

func LoadHTML() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.Handle("/", http.FileServer(http.Dir("./view")))
	http.HandleFunc("/newvalue", func(rw http.ResponseWriter, r *http.Request) {
		DataBase()
	})
	http.ListenAndServe(":5050", nil)
}
