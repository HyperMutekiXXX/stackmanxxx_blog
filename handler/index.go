package handler

import (
	"github.com/HyperMutekiXXX/stackmanxxx_blog/config"
	"github.com/HyperMutekiXXX/stackmanxxx_blog/consts"
	"github.com/HyperMutekiXXX/stackmanxxx_blog/model"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}

/**
加载
*/
func IndexHtml(w http.ResponseWriter, rq *http.Request) {
	t := template.New("index.html")
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	filenames := getFilenames()
	t, err := t.ParseFiles(filenames[consts.Index], filenames[consts.Header], filenames[consts.Home], filenames[consts.Footer], filenames[consts.Personal], filenames[consts.Post], filenames[consts.Pagination])
	if err != nil {
		log.Println(err)
	}
	var categorys = []model.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []model.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &model.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}

/**
获取文件名字map
*/
func getFilenames() map[int64]string {
	dir, _ := os.Getwd()
	index := dir + "\\template\\index.html"
	home := dir + "\\template\\home.html"
	header := dir + "\\template\\layout\\header.html"
	footer := dir + "\\template\\layout\\footer.html"
	personal := dir + "\\template\\layout\\personal.html"
	post := dir + "\\template\\layout\\post-list.html"
	pagination := dir + "\\template\\layout\\pagination.html"
	filenames := make(map[int64]string)
	filenames[consts.Index] = index
	filenames[consts.Home] = home
	filenames[consts.Header] = header
	filenames[consts.Footer] = footer
	filenames[consts.Personal] = personal
	filenames[consts.Post] = post
	filenames[consts.Pagination] = pagination

	return filenames
}
