package rcv

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	m "github.com/senradev/app/models"
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
	if err := m.Dbm.Insert(ul); err != nil {
		log.Println(err)
	}
	log.Println("data received", ul)
	w.Write([]byte("ok"))
}

func getUplink(r *http.Request) (*m.Uplink, error) {
	tmp := new(m.Uplink)

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
