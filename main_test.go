package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pedrohff/badger-poc/pkg"
	"github.com/pedrohff/badger-poc/pkg/cars"
	"sync"
	"testing"
	"time"
)

func setupDb() {
	var err error
	cars.Database, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=cars_db password=admin sslmode=disable")
	cars.Database.DB().SetMaxIdleConns(10)
	cars.Database.DB().SetMaxOpenConns(30)
	if err != nil {
		panic(err)
	}
}

func BenchmarkSeparatingTestsTogglingCacheLayer(b *testing.B) {
	setupDb()
	defer cars.Database.Close()

	type benchTest struct {
		percentageOfCachedObjects int
		dbNetworkDelay            int
	}

	dbNetworkDelay := 0
	goroutineLimit := 100000
	goroutineStartAt := 100
	goroutineMultiplier := 10

	tests := []benchTest{
		{5, dbNetworkDelay},
		{15, dbNetworkDelay},
		{40, dbNetworkDelay},
		{70, dbNetworkDelay},
	}

	// Running tests for the scenario without cache layer
	for goroutineCount := goroutineStartAt; goroutineCount <= goroutineLimit; goroutineCount *= goroutineMultiplier {
		badgerDb := pkg.SetupBadger()
		repository := cars.NewRepository(badgerDb, dbNetworkDelay)
		b.Run(fmt.Sprintf("[NOCACHE] %d requests", goroutineCount), func(b *testing.B) {
			group := sync.WaitGroup{}
			group.Add(goroutineCount)

			// Each interaction of this loop will represent a data access
			for i := 0; i < goroutineCount; i++ {
				id := fmt.Sprintf("test%d", i)
				go func(id string) {
					repository.FindById(id)
					group.Done()
				}(id)
			}
			group.Wait()
		})
	}

	// Running tests for the scenario with cache layer
	for _, test := range tests {
		for goroutineCount := goroutineStartAt; goroutineCount <= goroutineLimit; goroutineCount *= goroutineMultiplier {

			// Logic to avoid unnecessary loops or unnecessary cached objects
			finalCacheSize := int((float64(test.percentageOfCachedObjects) / 100) * float64(goroutineCount))

			//for benchmarkRepeatedLoops := 0; benchmarkRepeatedLoops < b.N; benchmarkRepeatedLoops++ {

			badgerDb := pkg.SetupBadger()
			repository := cars.NewRepository(badgerDb, test.dbNetworkDelay)

			for i := 0; i < finalCacheSize; i++ {
				id := fmt.Sprintf("test%d", i)
				if i < finalCacheSize {
					repository.Save(cars.Car{
						Id:           id,
						Model:        "S" + id,
						Manufacturer: "Tesla" + id,
					})
				}
			}
			b.Run(fmt.Sprintf("[>>CACHE] %d requests - caching %d%%(%d items)", goroutineCount, test.percentageOfCachedObjects, finalCacheSize), func(b *testing.B) {
				group := sync.WaitGroup{}
				group.Add(goroutineCount)

				// Each interaction of this loop will represent a data access
				for i := 0; i < goroutineCount; i++ {
					id := fmt.Sprintf("test%d", i)
					go func(id string) {
						repository.FindById(id)
						group.Done()
					}(id)
				}
				group.Wait()
			})
			badgerDb.Close()
		}
	}
}

//1400673600
//1364517100
//306148ms - 306148400700
// 82530ms -  82530983900
func BenchmarkAbc(t *testing.B) {
	timestart := time.Now()
	dbNetworkDelayMillis := 150
	amountToCache := 1500
	amountOfRequests := 2048
	enableCache := true

	badgerDb := pkg.SetupBadger()
	repository := cars.NewRepository(badgerDb, dbNetworkDelayMillis)
	group := sync.WaitGroup{}

	if enableCache {
		for i := 0; i < amountToCache; i++ {
			id := fmt.Sprintf("test%d", i)
			if i < amountToCache {
				repository.Save(cars.Car{
					Id:           id,
					Model:        "S" + id,
					Manufacturer: "Tesla" + id,
				})
			}
		}
	}

	repeatLoops := 1
	group.Add(amountOfRequests * repeatLoops)
	// Each test will be repeated N times according to "repeatLoops"
	for range make([]int, repeatLoops) {
		// Each interaction of this loop will represent a data access
		for i := 0; i < amountOfRequests; i++ {
			id := fmt.Sprintf("test%d", i)
			if enableCache {
				if i < amountToCache {
					repository.Save(cars.Car{
						Id:           id,
						Model:        "S" + id,
						Manufacturer: "Tesla" + id,
					})
				}
			}
			//go func(id string, cache *badger.DB) {
			repository.FindById(id)
			group.Done()
			//}(id, badgerDb)
		}
	}
	group.Wait()
	badgerDb.Close()
	fmt.Printf("time: %dms\n", time.Since(timestart).Milliseconds())
}

func FuncaoSomaRetornoEstranho(x int, y int) (resultado int, retornoEhPositivo bool) {
	resultado = x + y
	retornoEhPositivo = resultado >= 0
	return
}
