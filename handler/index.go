package handler

import (
	"encoding/json"
	"github.com/HyperMutekiXXX/stackmanxxx_blog/model"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Index(w http.ResponseWriter, rq *http.Request) {
	rq.Header.Set("Content-Type", "application/json")
	index := &model.Index{
		Title:   "你好",
		Message: "nb",
	}
	jsonStr, _ := json.Marshal(index)
	w.Write(jsonStr)
}

func IndexHtml(w http.ResponseWriter, rq *http.Request) {
	t := template.New("index.html")
	dir, _ := os.Getwd()
	t, _ = t.ParseFiles(dir + "\\template\\index.html")
	index := &model.Index{
		Title:   "你好",
		Message: "nb",
	}
	err := t.Execute(w, index)
	if err != nil {
		log.Println("template execute err : ", err)
	}

}
