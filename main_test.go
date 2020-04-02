package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/pedrohff/badger-poc/pkg"
	"github.com/pedrohff/badger-poc/pkg/cars"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {

	type benchTest struct {
		requests       int
		cacheSize      int
		dbNetworkDelay int
	}

	dbNetworkDelay := 150
	tests := []benchTest{
		{1000, 15, dbNetworkDelay},
		{10000, 15, dbNetworkDelay},
		{100000, 15, dbNetworkDelay},
		{1000, 100, dbNetworkDelay},
		{10000, 100, dbNetworkDelay},
		{100000, 100, dbNetworkDelay},
		{1000, 1000, dbNetworkDelay},
		{10000, 1000, dbNetworkDelay},
		{100000, 1000, dbNetworkDelay},
	}

	for _, test := range tests {
		var totalGain int64
		t.Run(fmt.Sprintf("%d requests - caching %d items", test.requests, test.cacheSize), func(t *testing.T) {
			for i := 0; i < 100; i++ {

				timerWithCache := benchMarkReturningMillisSpent(test.requests, test.cacheSize, true, test.dbNetworkDelay)
				timerWithNoCache := benchMarkReturningMillisSpent(test.requests, test.cacheSize, false, test.dbNetworkDelay)

				//fmt.Printf("Cache timer: %dms\n", timerWithCache)
				//fmt.Printf("No cache timer: %dms\n", timerWithNoCache)

				totalGain += timerWithNoCache - timerWithCache

				// TODO swap this with counters and just measure the percentage of success
				assert.True(t, timerWithNoCache > timerWithCache)
			}
		})
		avgGain := totalGain / 100
		fmt.Printf("Requests: %d | CacheSize: %d | Avg gain: %dms\n", test.requests, test.cacheSize, avgGain)
	}

}

func benchMarkReturningMillisSpent(amountOfRequests int, amountToCache int, shouldCache bool, dbNetworkDelayMillis int) int64 {
	timeStarted := time.Now()
	badgerDb := pkg.SetupBadger()
	repository := cars.NewRepository(badgerDb, dbNetworkDelayMillis)
	group := sync.WaitGroup{}
	group.Add(amountOfRequests)
	for i := 0; i < amountOfRequests; i++ {
		id := fmt.Sprintf("test%d", i%amountToCache)
		if shouldCache {
			if i < amountToCache {
				repository.Save(cars.Car{
					Id:           id,
					Model:        "S" + id,
					Manufacturer: "Tesla" + id,
				})
			}
		}
		go func(id string, cache *badger.DB) {
			//start := time.Now()
			repository.FindById(id)
			//fmt.Printf("[%d] time: %dms\n", index, time.Since(start).Milliseconds())
			group.Done()

		}(id, badgerDb)
	}
	group.Wait()
	badgerDb.Close()
	return time.Since(timeStarted).Milliseconds()
}
