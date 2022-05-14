package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) { //127.0.0.1/objects/test 会创建该文件，并将响应内容写入该文件
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
