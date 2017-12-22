package models

import "time"

// Uplink represents uplink received from light node
type Uplink struct {
	DevEui     string    `db:"deveui" json:"devEui"`
	GatewayEui string    `db:"gweui" json:"gwEui"`
	JoinID     int       `db:"joinid" json:"joinId"`
	PDU        string    `db:"pdu" json:"pdu"`
	Port       int       `db:"port" json:"port"`
	SeqNum     int       `db:"seqnum" json:"seqno"`
	TxTime     time.Time `db:"txtime" json:"txtime"`
}
