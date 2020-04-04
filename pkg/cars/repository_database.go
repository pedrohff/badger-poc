package cars

import (
	"encoding/json"
	"fmt"
	"time"
)

// Represents a Repository that will fetch/input data on a remote database
type databaseRepository struct {
	dbNetworkDelay int
}

func (d databaseRepository) FindById(id string) (*Car, error) {
	//fmt.Printf("\t> Finding id: %s\n", id)

	// Adding a timer to simulate db connection and external network access
	// 150
	time.Sleep(time.Millisecond * time.Duration(d.dbNetworkDelay))

	// Adding some complexity to it
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 10; b++ {
			for c := 1; c <= 10; c++ {
				_ = a * (b * c)
			}
		}
	}
	jsonStr := fmt.Sprintf(`{"Id": "%s", "Model":"Uno", "Manifacturer": "Fiat"}`, id)
	var car Car
	e := json.Unmarshal([]byte(jsonStr), &car)
	if e != nil {
		return nil, e
	}

	return &car, nil
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
