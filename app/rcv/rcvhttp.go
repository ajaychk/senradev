package rcv

import (
	"encoding/json"
	"log"
	"net/http"
)

var ChanUplink = make(chan interface{}, 10)

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
	tmp := struct {
		DevEui     string `json:"devEui"`
		GatewayEui string `json:"gwEui"`
		JoinID     int    `json:"joinId"`
		PDU        string `json:"pdu"`
		Port       int    `json:"port"`
		SeqNum     int    `json:"seqno"`
		TxTime     string `json:"txtime"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		log.Println("invalid data received")
		w.Write([]byte("invalid data received"))
		return
	}
	ChanUplink <- tmp
	log.Println("data received", tmp)
	w.Write([]byte("ok"))
}
