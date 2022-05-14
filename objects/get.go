package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//var StorageRoot string = "/home/ubuntu"

func get(w http.ResponseWriter, r *http.Request) { //类似put，但是是将文件拷贝到响应体
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f) //将文件内容拷贝到响应体w中
}
