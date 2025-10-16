package golevelui

import (
	"encoding/json"
	"net/http"

	"github.com/siddontang/go/hack"
	"github.com/syndtr/goleveldb/leveldb"
)

type updateRes struct {
	Success bool
}

type updateReq struct {
	DB    string
	Key   string
	Value string
}

func (l *LevelAdmin) apiKeyUpdate(writer http.ResponseWriter, request *http.Request) {
	reqData := &updateReq{}
	err := json.NewDecoder(request.Body).Decode(&reqData)
	if err != nil {
		l.writeError(writer, err)

		return
	}

	if reqData.DB == "" || reqData.Key == "" {
		http.NotFound(writer, request)
		return
	}

	if load, ok := l.dbs.Load(reqData.DB); ok {
		db := load.(*leveldb.DB)
		err := db.Put(hack.Slice(reqData.Key), []byte(reqData.Value), nil)
		if err != nil {
			l.writeError(writer, err)
			return
		}

		l.writeJson(writer, &updateRes{Success: true})
	} else {
		http.NotFound(writer, request)
	}
}
