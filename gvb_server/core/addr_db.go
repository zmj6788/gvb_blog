package core

import (
	"gvb_server/global"
	"log"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatal(err)
	}
	global.AddrDB = db
}
