package service

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.HandleFunc("/home/date", apitest(formatter)).Methods("GET")
	mx.HandleFunc("/", homeHandle).Methods("GET")
	mx.HandleFunc("/",Login).Methods("POST")
	mx.HandleFunc("/unknown", NotImplemented)
	/*实现对css/js/images的引用*/
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/static/"))))
}
