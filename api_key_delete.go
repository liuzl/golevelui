package golevelui

import (
	"encoding/json"
	"net/http"

	"github.com/siddontang/go/hack"
	"github.com/syndtr/goleveldb/leveldb"
)

type deleteRes struct {
	Success bool
}

type deleteReq struct {
	DB  string
	Key string
}

func (l *LevelAdmin) apiKeyDelete(writer http.ResponseWriter, request *http.Request) {
	reqData := &deleteReq{}
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
		if has, err := db.Has(hack.Slice(reqData.Key), nil); has && err == nil {
			if err := db.Delete(hack.Slice(reqData.Key), nil); err != nil {
				l.writeError(writer, err)
				return
			}
		} else {
			http.NotFound(writer, request)
			return
		}

		l.writeJson(writer, &deleteRes{Success: true})
	} else {
		http.NotFound(writer, request)
	}
}
