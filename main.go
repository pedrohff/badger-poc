package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/y"
	"github.com/pedrohff/badger-poc/pkg"
	"github.com/pedrohff/badger-poc/pkg/cars"
	"sync"
	"time"
)

func main() {
	badgerDb := pkg.SetupBadger()
	repository := cars.NewRepository(badgerDb, 100)

	startOuter := time.Now()

	group := sync.WaitGroup{}
	amountOfTests := 1000
	group.Add(amountOfTests)

	for i := 0; i < amountOfTests; i++ {
		id := fmt.Sprintf("test%d", i%15)
		if i < 15 {
			repository.Save(cars.Car{
				Id:           id,
				Model:        "S" + id,
				Manufacturer: "Tesla" + id,
			})
		}
		go func(id string, cache *badger.DB) {
			//start := time.Now()
			repository.FindById(id)
			//fmt.Printf("[%d] time: %dms\n", index, time.Since(start).Milliseconds())
			group.Done()

		}(id, badgerDb)
	}
	group.Wait()

	fmt.Printf("NumReads: %v\n", y.NumReads)
	fmt.Printf("NumGets: %v\n", y.NumGets)
	fmt.Printf("NumPuts: %v\n", y.NumPuts)
	fmt.Printf("BytesRead: %v\n", y.NumBytesRead)
	fmt.Printf("NumMemtableGets: %v\n", y.NumMemtableGets)

	//badgerDb.DataCacheMetrics()
	//fmt.Printf("hits: %d\n", badgerDb.DataCacheMetrics().Hits())
	//fmt.Printf("misses: %v\n", badgerDb.DataCacheMetrics().Misses())
	//fmt.Printf("metrics: %s\n", badgerDb.DataCacheMetrics().String())

	badgerDb.Close()
	fmt.Printf("\n\ntest time: %dms\n", time.Since(startOuter).Milliseconds())
}
