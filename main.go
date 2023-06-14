package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"
)

// TemperatureDataElem 気温データの一つのデータセット
type TemperatureDataElem struct {
	Label string
	Data  []float64
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/html/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var temperatureData []TemperatureDataElem
	temperatureData = append(temperatureData, TemperatureDataElem{
		Label: "沖縄県",
		Data:  []float64{17.0, 17.1, 18.9, 21.4, 24.0, 26.8, 28.9, 28.7, 27.6, 25.2, 22.1, 18.7},
	})
	temperatureData = append(temperatureData, TemperatureDataElem{
		Label: "東京都",
		Data:  []float64{5.2, 5.7, 8.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
	})
	t, err := template.ParseFiles("view/html/data.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, temperatureData); err != nil {
		panic(err.Error())
	}
}

func typingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/html/typing.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

var views embed.FS

func main() {
	dir, _ := os.Getwd()
	log.Println(dir)
	log.Print(http.Dir(dir + "/view/scripts/"))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/view/html/data/", dataHandler)
	http.HandleFunc("/view/html/typing/", typingHandler)
	http.Handle("/view/", http.FileServer(http.FS(views)))

	http.Handle("/view/scripts/", http.StripPrefix("/view/scripts/", http.FileServer(http.Dir(dir+"/view/scripts/"))))
	http.Handle("/view/style/", http.StripPrefix("/view/style/", http.FileServer(http.Dir(dir+"/view/style/"))))
	http.ListenAndServe(":60180", nil)
}
