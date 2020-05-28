package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
)

var upGrade = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

const message  = "Hello Gopher's World!"
const wsmessage  = "Hello webSocket is ok!"

func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w,message)
	w.Header().Set("content-type","text/plain; charset=uft-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func reader(conn *websocket.Conn)  {
	for{
		messageType, p ,err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType,p); err != nil {
			log.Println(err)
			return
		}
	}
}
/*
func writer(conn *websocket.Conn)  {
	var str []byte
	str =[]byte("Welcome my VIP!")
	err := conn.WriteMessage(1,str)
	if err != nil{
		log.Fatal("Write Message Error: ", err)
	}
}
*/

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upGrade.CheckOrigin = func(r *http.Request) bool {return true}
	ws, err := upGrade.Upgrade(w,r,nil)
	if err != nil {
		fmt.Println(err)
	
	}
	fmt.Println("Clinet Successfully connected ...")
	reader(ws)
	//writer(ws)


}



func main()  {
	fmt.Println("Go WebScoket Server is running ... ")
	mux := http.NewServeMux()
	mux.HandleFunc("/",homePage)
	mux.HandleFunc("/ws",wsEndpoint)
	log.Fatal(http.ListenAndServe(":80", mux))

}
