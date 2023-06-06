package routers

import (
	"github.com/HyperMutekiXXX/stackmanxxx_blog/handler"
	"net/http"
)

func InitIndex() {
	http.HandleFunc("/", handler.IndexHtml)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
