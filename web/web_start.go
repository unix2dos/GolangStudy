package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
)

func luyou(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world")
}


func main(){
	http.HandleFunc("/", luyou)
	err := http.ListenAndServe(":9090", nil) //最重要的步骤, 第二个参数是路由处理
	if err != nil {
		log.Fatal("err: ", err)
	}

	//	ln, err := net.Listen("tcp", addr)
	//  srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
	//  go c.serve() 每个请求建立一个goroutine
	//  c.readRequest()分析请求
	//  映射url与路由
}