package main

import (
	"fmt"
	"time"

	"andiDB/storage"
)

var (
	delay = 59
)

// AsyncDumpDb periodically dump db to bson file
func AsyncDumpDb() {
	for t := range time.Tick(time.Duration(delay) * time.Second) {
		_ = t
		fmt.Println("will dump..")
		go dump()
	}
}

func dump() {
	dbs := storage.AllDb()
	fmt.Println(dbs)
}
