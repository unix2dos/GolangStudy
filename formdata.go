package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "shoudao")
	}
}
func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/login?username=aaa
//username	liuwei
//output: username: [liuwei aaa]
