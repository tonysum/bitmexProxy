package main

import (
	"net/http"

	"fmt"
)

const msg string = "Hello Gopher's World!"
func MainProc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(msg))
}
func main()  {
	proxyServer := http.NewServeMux()
	proxyServer.HandleFunc("/",MainProc)
	//http.HandleFunc("/",MainProc)
	err := http.ListenAndServe(":8090", proxyServer)
	if err != nil {
		fmt.Println("ok")
	}
}
