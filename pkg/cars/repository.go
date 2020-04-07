package cars

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

type Repository interface {
	FindById(id string) (*Car, error)
	Save(car Car) (*Car, error)
	Update(car Car) (*Car, error)
}

type repositoryOrchestrator struct {
	cache    cacheRepository
	database databaseRepository
}

func NewRepository(cache *badger.DB, dbNetworkDelay int) Repository {
	return repositoryOrchestrator{
		cache:    cacheRepository{cache: cache},
		database: databaseRepository{dbNetworkDelay: dbNetworkDelay},
	}
}

func (r repositoryOrchestrator) FindById(id string) (*Car, error) {
	byId, err := r.cache.FindById(id)
	switch err {
	case nil:
		return byId, nil
	case badger.ErrKeyNotFound:
		return r.database.FindById(id)
	default:
		return nil, err
	}
}

func (r repositoryOrchestrator) Save(car Car) (*Car, error) {
	savedCar, dberr := r.database.Save(car)
	if dberr != nil {
		return nil, dberr
	}
	_, err := r.cache.Save(*savedCar)
	if err != nil {
		return nil, err
	}
	return savedCar, nil
}

func (r repositoryOrchestrator) Update(car Car) (*Car, error) {
	panic("implement me")
}
