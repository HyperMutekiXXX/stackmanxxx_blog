package model

import "github.com/HyperMutekiXXX/stackmanxxx_blog/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
