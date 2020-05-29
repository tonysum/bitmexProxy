package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

func homePageHandler(tpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r)
	})
}

func main() {
	flag.Parse()
	tpl := template.Must(template.ParseFiles("index.html"))
	//h := newHub()
	router := http.NewServeMux()
	router.HandleFunc("/", homePageHandler(tpl))
	//router.HandleFunc("/ws", &wsHandler{h: h})
	log.Println("Go WebScoket Server is running on port 8080 ")
	log.Fatal(http.ListenAndServe(":8080", router))

}
