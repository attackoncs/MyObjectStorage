package main

import (
	"fmt"
	"log"
	"main/objects"
	"net/http"
	"os"
)

func main() {
	//设置存储目录和监听socket，不知为啥再终端中设置可以打印但是无法os.Getenv失败
	os.Setenv("STORAGE_ROOT", "/home/ubuntu")
	fmt.Println(os.Getenv("STORAGE_ROOT"))
	os.Setenv("LISTEN_ADDR", "127.0.0.1:1280")
	fmt.Println(os.Getenv("LISTEN_ADDR"))
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR"), nil))
}
