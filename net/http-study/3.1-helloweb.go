package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"golang.org/x/text/message"
)

func helloWorld (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO GO WEB")
}


func myweb (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "HELLO GO WEB")
	t, err := template.ParseFiles("./shangguigu.tmpl")
	if err != nil {
		fmt.Println("t p failed, err:  ",err)
		return
	}
	name := "study gogogogogo!!!!!!!!!"
	t.Execute(w,name)

	len := r.ContentLength
	bodyContent := make([]byte, len)
	r.Body.Read(bodyContent)
	fmt.Fprintln(w, string(bodyContent))
}

/*func getBody (w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	bodyContent := make([]byte, len)
	r.Body.Read(bodyContent)
	fmt.Fprintln(w, string(bodyContent))
}*/


func main () {

	mux := http.NewServeMux()

	// 处理静态文件
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/shop/", myweb)
	mux.HandleFunc("/hello",helloWorld)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe();err != nil{
		log.Fatal(err)
	}
}

