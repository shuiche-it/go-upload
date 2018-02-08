package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
)

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request)  {


	fmt.Println("method:", r.Method) //POST

	//因为上传文件的类型是multipart/form-data 所以不能使用 r.ParseForm(), 这个只能获得普通post
	//r.ParseMultipartForm(32 << 20) //上传最大文件限制32M

	file, handler, err := r.FormFile("id")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("/Users/zhangzexuan/gopath/src/study/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
func main() {

	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe("0.0.0.0:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
