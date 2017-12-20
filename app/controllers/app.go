package controllers

import (
	r "github.com/revel/revel"
	"golang.org/x/net/websocket"

	"github.com/senradev/app/rcv"
)

type App struct {
	*r.Controller
}

func (c App) Index() r.Result {
	return c.Render()
}

// Uplink receives and sends uplink
func (c App) Uplink(ws *websocket.Conn) r.Result {
	defer ws.Close()

	var rcvd string
	websocket.Message.Receive(ws, &rcvd)
	r.INFO.Println("message received", rcvd)

	for {
		ul := <-rcv.ChanUplink

		if err := websocket.JSON.Send(ws, ul); err != nil {
			r.ERROR.Println("error in sending data")
			break
		}
		r.INFO.Println("uplink sent:", ul)
	}
	return c.Render()
}
