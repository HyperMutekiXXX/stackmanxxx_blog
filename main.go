package main

import (
	"github.com/HyperMutekiXXX/stackmanxxx_blog/routers"
	"log"
	"net/http"
)

func main() {
	h := &http.Server{
		Addr: "127.0.0.1:8088",
	}
	routers.InitIndex()

	if err := h.ListenAndServe(); err != nil {
		log.Println("listen and serve err: ", err)
	}
}
