package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pedrohff/badger-poc/pkg"
	"github.com/pedrohff/badger-poc/pkg/cars"
	"math/rand"
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
			// Clearing all database cache
			defer cars.Database.Exec("discard all")
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

			// Creating an unordered array so the database rows won't be read sequentially
			unorderedIntArray := createUnorderedArray(goroutineCount)

			badgerDb := pkg.SetupBadger()
			repository := cars.NewRepository(badgerDb, test.dbNetworkDelay)

			// Forcing data to be saved to cache
			// This is not the best approach for saving data to cache, its just the best way for me to control directly what is being cached
			// Also, the save method on the database repository is not doing anything
			for i := 0; i < finalCacheSize; i++ {
				id := fmt.Sprintf("test%d", unorderedIntArray[i])
				if i < finalCacheSize {
					repository.Save(cars.Car{
						Id:           id,
						Model:        "S" + id,
						Manufacturer: "Tesla" + id,
					})
				}
			}

			b.Run(fmt.Sprintf("[>>CACHE] %d requests - caching %d%%(%d items)", goroutineCount, test.percentageOfCachedObjects, finalCacheSize), func(b *testing.B) {
				// Clearing all database cache
				defer cars.Database.Exec("discard all")
				group := sync.WaitGroup{}
				group.Add(goroutineCount)

				// Each interaction of this loop will represent a data access
				for i := 0; i < goroutineCount; i++ {
					id := fmt.Sprintf("test%d", unorderedIntArray[i])
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

// This function can be improved, as it seems that rand.Shuffle already creates an unordered array
// so maybe the first loop is unnecessary
func createUnorderedArray(arrayLength int) []int {
	resultArray := make([]int, 0)
	for i := 0; i < arrayLength; i++ {
		resultArray = append(resultArray, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(resultArray), func(i, j int) { resultArray[i], resultArray[j] = resultArray[j], resultArray[i] })
	return resultArray
}
