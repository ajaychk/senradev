package rcv

import (
	"encoding/hex"
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
	ul, err := getUplink(r)

	if err != nil {
		w.Write([]byte("invalid data received"))
		return
	}

	ChanUplink <- ul
	log.Println("data received", ul)
	w.Write([]byte("ok"))
}

type uplinkType struct {
	DevEui     string `json:"devEui"`
	GatewayEui string `json:"gwEui"`
	JoinID     int    `json:"joinId"`
	PDU        string `json:"pdu"`
	Port       int    `json:"port"`
	SeqNum     int    `json:"seqno"`
	TxTime     string `json:"txtime"`
}

func getUplink(r *http.Request) (*uplinkType, error) {
	tmp := new(uplinkType)

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		log.Println("invalid data received", err)
		return nil, err
	}

	hPDU, err := hex.DecodeString(tmp.PDU)
	if err != nil {
		log.Println("invalid pdu received", err)
		return nil, err
	}
	tmp.PDU = string(hPDU)
	return tmp, nil
}
