
package main
import(
	"net/http"
)

func Sayhello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello"))
}

/*func main() {
	http.HandleFunc("/sayhello",Sayhello)
	http.ListenAndServe(":8080", nil)
}
*/	