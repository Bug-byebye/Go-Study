package main
import (
	"net/http"
	"fmt"
	"html/template"
)

func helloHandleFanc (w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template-exm.tmpl")
	if err != nil {
		fmt.Println("t p failed, err:  ",err)
		return
	}
	name := "study gogogogogo!!!!!!!!!"
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/", helloHandleFanc)
	http.ListenAndServe(":8086", nil)
}
