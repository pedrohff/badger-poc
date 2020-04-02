package cars

import (
	"time"
)

// Represents a repository that will fetch/input data on a remote database
type databaseRepository struct {
	dbNetworkDelay int
}

func (d databaseRepository) FindById(id string) (*Car, error) {
	//fmt.Printf("\t> Finding id: %s\n", id)

	// Adding a timer to simulate db connection and external network access
	time.Sleep(time.Millisecond * time.Duration(d.dbNetworkDelay))
	return &Car{
		Id:           id,
		Model:        "Uno",
		Manufacturer: "Fiat",
	}, nil
}

func (d databaseRepository) Save(car Car) (*Car, error) {
	//fmt.Printf("\t> Save %s\n", car.Id)

	// Adding a timer to simulate db connection and external network access
	//time.Sleep(time.Millisecond * 150)
	return &car, nil
}

func (d databaseRepository) Update(car Car) (*Car, error) {
	//fmt.Printf("\t> Update %s\n", car.Id)

	// Adding a timer to simulate db connection and external network access
	time.Sleep(time.Millisecond * 150)
	return &car, nil
}
