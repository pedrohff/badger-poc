# Benchmark results

All tests were performed by adding a **150ms** delay for mocking each database operation.

## Comparing scenarios
The purpose of this test is to assert that by using the cache layer, there will always be performance gains, with three different groups. Each group will hold a specific amount of structures in it's cache layer.

The test is setup to 

### Running the test disabling the cache layer

```shell script
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_1_requests_-_caching_15_items-6         	       7	 160577686 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_2_requests_-_caching_15_items-6         	       7	 155770557 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_4_requests_-_caching_15_items-6         	       7	 155893271 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_8_requests_-_caching_15_items-6         	       7	 155864586 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_16_requests_-_caching_15_items-6        	       7	 155018200 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_32_requests_-_caching_15_items-6        	       7	 155722529 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_64_requests_-_caching_15_items-6        	       7	 154892686 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_128_requests_-_caching_15_items-6       	       7	 154797286 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_256_requests_-_caching_15_items-6       	       7	 156236157 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_512_requests_-_caching_15_items-6       	       7	 157309871 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[NOCACHE]_1024_requests_-_caching_15_items-6      	       7	 156345386 ns/op
```

### Caching 15 objects 

```shell script
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1_requests_-_caching_15_items-6         	      43	  26862181 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_2_requests_-_caching_15_items-6         	      51	  22108608 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_4_requests_-_caching_15_items-6         	      99	  24571610 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_8_requests_-_caching_15_items-6         	     100	  12910838 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_16_requests_-_caching_15_items-6        	     100	  12020022 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_32_requests_-_caching_15_items-6        	      63	  19165667 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_64_requests_-_caching_15_items-6        	      99	  16151790 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_128_requests_-_caching_15_items-6       	     100	  12239982 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_256_requests_-_caching_15_items-6       	      72	  17847194 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_512_requests_-_caching_15_items-6       	      70	  18356773 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1024_requests_-_caching_15_items-6      	      85	  29105631 ns/op
```

### Caching  100 objects


```shell script
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1_requests_-_caching_100_items-6        	     100	  22667712 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_2_requests_-_caching_100_items-6        	     100	  13959263 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_4_requests_-_caching_100_items-6        	     100	  11037310 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_8_requests_-_caching_100_items-6        	     100	  11270463 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_16_requests_-_caching_100_items-6       	     100	  16861488 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_32_requests_-_caching_100_items-6       	     100	  19395168 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_64_requests_-_caching_100_items-6       	     100	  18492352 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_128_requests_-_caching_100_items-6      	      99	  16898978 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_256_requests_-_caching_100_items-6      	      99	  16934031 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_512_requests_-_caching_100_items-6      	      93	  19785825 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1024_requests_-_caching_100_items-6     	      79	  32096644 ns/op
```
### Caching  1000 objects

```shell script
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1_requests_-_caching_1000_items-6       	     100	  14500265 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_2_requests_-_caching_1000_items-6       	     121	  11867466 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_4_requests_-_caching_1000_items-6       	      99	  16858590 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_8_requests_-_caching_1000_items-6       	     100	  14833841 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_16_requests_-_caching_1000_items-6      	     110	  10586738 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_32_requests_-_caching_1000_items-6      	     120	   8901569 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_64_requests_-_caching_1000_items-6      	     100	  10100271 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_128_requests_-_caching_1000_items-6     	     115	  10678866 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_256_requests_-_caching_1000_items-6     	     105	  15552615 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_512_requests_-_caching_1000_items-6     	      62	  21918826 ns/op
BenchmarkSeparatingTestsTogglingCacheLayer/[>>CACHE]_1024_requests_-_caching_1000_items-6    	      51	  28352718 ns/op
```

### Result