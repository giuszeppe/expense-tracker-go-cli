package single

import (
	"database/sql"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	Database *sql.DB
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			db, err := sql.Open("sqlite3", "file.db")
			if err != nil {
				log.Fatal(err)
			}
			singleInstance = &Single{Database: db}
		}
	}

	return singleInstance
}
