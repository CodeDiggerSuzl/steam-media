package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
 	http.HandleFunc("/",firstpage)
	http.ListenAndServe(":8080",nil)
}
func firstpage(w http.ResponseWriter,r *http.Request){
	fmt.Println(r)
	io.WriteString(w,"<h1>My stream media</h1>")
}