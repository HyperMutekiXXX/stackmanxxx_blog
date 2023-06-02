package routers

import (
	"github.com/HyperMutekiXXX/stackmanxxx_blog/handler"
	"net/http"
)

func InitIndex() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/index.html", handler.IndexHtml)
}
