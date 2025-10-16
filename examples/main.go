package main

import (
	"github.com/liuzl/golevelui"
	"github.com/syndtr/goleveldb/leveldb"
	"time"
)

func main() {
	file, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}
	golevelui.GetLevelAdmin().Register(file, "db")
	err = golevelui.GetLevelAdmin().Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}
