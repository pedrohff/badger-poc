package pkg

import (
	"github.com/dgraph-io/badger/v2"
)

func SetupBadger() *badger.DB {
	opts := badger.DefaultOptions("").WithInMemory(true).WithEventLogging(false).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db
}
