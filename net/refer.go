package main
import(
	"net/http"
)

type Refer struct{
	handler http.Handler
	refer string
}

//func (this *Refer) ServeHTTP (w http.ResponseWriter, r *http.Request) {
func (this *Refer) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.Referer()==this.refer {
		this.handler.ServeHTTP(w,r)	
	}else{
		w.WriteHeader(403)
	}
}

func myHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is my handler"))
}

func hello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello........"))
}
/*
func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer: "www.shirdon.com",
	}
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":9090", referer)
}*/