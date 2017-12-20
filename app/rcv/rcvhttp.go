package rcv

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	go initServer()
}

func initServer() {
	http.HandleFunc("/uplink", handleUplink)
	err := http.ListenAndServe(":9011", nil)
	if err != nil {
		panic(err)
	}
}

func handleUplink(w http.ResponseWriter, r *http.Request) {
	log.Println("uplink received")
	var tmp interface{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		log.Println("invalid data received")
		w.Write([]byte("invalid data received"))
		return
	}
	log.Println("data received", tmp)
	w.Write([]byte("ok"))
}
