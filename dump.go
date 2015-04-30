package main

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"os"
	"time"

	"andiDB/storage"

	"gopkg.in/mgo.v2/bson"
)

var (
	delay = 59

	fileNamePrefix = "dump.db"
)

// AsyncDumpDb periodically dump db to bson file
func AsyncDumpDb() {
	for t := range time.Tick(time.Duration(delay) * time.Second) {
		_ = t
		fmt.Println("will dump..")
		go dump()
	}
}

// dump databases to fileNamePrefix
func dump() {
	dbs := storage.AllDb()
	fmt.Println(dbs)
	d, err := serializeToBson(dbs...)
	if err != nil {
		fmt.Println(err)
		return
	}
	writeToFile(d...)
	LoadFile()
}

// serializeToBson convert json/struct data to BSON
func serializeToBson(databases ...interface{}) ([]byte, error) {
	data, err := bson.Marshal(databases)
	if err != nil {
		fmt.Println(err)
		var b []byte
		return b, err
	}
	fmt.Printf("%q", data)
	return data, nil
}

// writeToFile attempt to write byteArr to file
func writeToFile(p ...byte) {

	fmt.Println("writing to: ", fileNamePrefix)

	f, err := os.Create(fileNamePrefix)
	if err != nil {

		// do better error handling here
		// if cant save file
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	//var i io.Writer
	n, err := f.Write(p)
	if err != nil {
		fmt.Println(n, err)
	}
	fmt.Printf("wrote %d bytes: ", n)
}

// LoadFile load database file
func LoadFile() {
	b, err := ioutil.ReadFile(fileNamePrefix)
	if err != nil {
		panic(err)
	}

	fmt.Println("files read..")
	var s interface{}
	err = bson.Unmarshal(b, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for the meantime print the result
	fmt.Println(s)
}
