package cars

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v2"
)

type cacheRepository struct {
	cache *badger.DB
}

func (c cacheRepository) FindById(id string) (*Car, error) {
	car := &Car{}
	err := c.cache.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			valCopy := append([]byte{}, val...)
			return json.Unmarshal(valCopy, car)
		})
	})

	if err != nil {
		return nil, err
	}

	//fmt.Print(c.cache.DataCacheMetrics() == nil)
	//fmt.Println("cache hit")
	return car, nil
}

func (c cacheRepository) Save(car Car) (*Car, error) {
	return c.saveUpdate(car)
}

func (c cacheRepository) saveUpdate(car Car) (*Car, error) {
	marshal, err := json.Marshal(car)
	if err != nil {
		return nil, err
	}
	err = c.cache.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(car.Id), marshal)
	})

	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (c cacheRepository) Update(car Car) (*Car, error) {
	return c.saveUpdate(car)
}
