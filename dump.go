package main

import (
	"fmt"
	"time"

	"andiDB/storage"

	"gopkg.in/mgo.v2/bson"
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
	serializeToBson(dbs...)
}

func serializeToBson(databases ...interface{}) {
	data, err := bson.Marshal(databases)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q", data)
}
