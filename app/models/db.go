package models

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //for postgres
	"github.com/revel/revel"
)

// Schema is namespace for db
const Schema = "dev"

var (
	// Dbm is db handle to use
	Dbm *gorp.DbMap
)

func init() {
	initDB()
	updateModels()
}

func initDB() {
	var dbInfo string

	dbLocal := false
	if dbLocal {
		revel.INFO.Println("DATABASE LOCAL")
		dbInfo = "user=senhavells password=sc2devsenra dbname=senra sslmode=disable"
	} else {
		revel.INFO.Println("DATABASE RDS")
		dbInfo = `host=senra.cqwf1pvghoch.us-west-2.rds.amazonaws.com
		 user=senhavells password=sc2devsenra dbname=senra sslmode=disable`
	}

	Db, err := sql.Open("postgres", dbInfo)
	if Db == nil || err != nil {
		revel.ERROR.Println("could not connect to postgres", dbInfo)
		panic(err)
	}

	Dbm = &gorp.DbMap{Db: Db, Dialect: gorp.PostgresDialect{}}
}

func updateModels() {
	Dbm.AddTableWithNameAndSchema(Uplink{}, Schema, "uplink").SetKeys(false, "DevEui",
		"GatewayEui", "PDU", "TxTime")
	// t = Dbm.AddTableWithName(DownlinkPayloadStatus{}, "downlink_status").SetKeys(false,
	// 	"Deveui", "ID", "TransmissionStatus")
	// t = Dbm.AddTableWithName(NodeInfo{}, "node_info").SetKeys(true, "ID")
	// t = Dbm.AddTableWithName(JoinInfo{}, "join_info").SetKeys(true, "ID")
	// t = Dbm.AddTableWithName(NodeStatus{}, "node_status").SetKeys(false, "ID")

	Dbm.TraceOn("[gorp]", revel.INFO)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		log.Fatalln(err, "Error in creating tables")
	}
}
