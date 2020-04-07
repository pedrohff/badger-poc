# Development journal

## April 4th, 2020

`[2pm]`  
As I'm having lots of stuff to implement, structure and test, came to mind an idea of writing out the ideas passing through the development of this proof of concept.

For now, the structure of the code seems to be exactly what I would use in a production system, it lacks only a few tweaks on how and when the data should be cached, and some more control or exposition over the cache layer.

Since the repository's database layer is a simple mock, there's been some trouble on proving that the cache layer will show performance improvements. I've tried adding some timers and tweaking it, also tried adding some code-complexity to make this database access mock perform slower than using cache - understanding I could achieve false positive results.

Besides, I felt stuck on developing the tests, it has been overwritten a lot so far, and I feel I'm getting in a comfortable spot, not only for structuring but for reading the results as well.

Therefore, my next move will be to connect a real database to the project, check how I should structure its data, and rework my tests, adding some more random access, to reduce any kind of bias, as the results are based on linear loops, with sequenced arrays. 

## April 7th, 2020

Starting describing the day before, I did as planned, connected the project to a PostgreSQL database, created its structure, installed [GORM](https://gorm.io/) (the ORM I use at work) then set everything up. The structure of the database variable is not ideal, please do not take this as a good pattern, at least try to protect the access to the database, I'm leaving it like this for tests purpose only.

This initial database setup took me some time, then I got back to testing, integrating them with the database. For starters, I still left a 150ms delay before each database access to simulate a scenario where there's a considerable physical distance between the application and the database, as this used to be a production scenario for some time.  

The first thing I noticed during the first tests, was that there was a gap in the average time of operations. As the tests were set up to simulate scenarios with 100, 1000, 10000 and 100000 goroutines or requests, and with the growth of parallelism, the overall performance got worse, and for controlling that, it was better to separate and compare the results in two groups, one will group the results of the tests with 100 and 1000 requests, the other grouping tests with 10000 and 100000 requests.

As expected, I finally got the results I wanted, the cache layer always showed performance improvement as the cache size increased. The final touch ups before publishing and documenting the results, I pretend to randomize the access of each database row, and try to clean up the database's cache after each test, again, to remove any bias.
       