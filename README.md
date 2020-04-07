# Badger POC

The purpose of this repository is to study and achieve an ideal strategy for accessing data in multiple repositories, for now, one will be for caching, and the other will the source-of-truth such as a database. The database access will be mocked, with just a parametrized delay. 

Another objective of the code implemented, is to test and compare how an application will perform by applying an in-memory cache layer over the database access, tweaking the amount of access, the cache size, and the mocked database access average delay.

Using [Badger](https://github.com/dgraph-io/badger) for in-memory cache, [GORM](https://gorm.io/) for database integration, and Go's standard lib for everything else.

Make sure to check the [project's journal](https://github.com/pedrohff/badger-poc/blob/master/JOURNAL.md) to see how it is being developed and how I organized my ideas for it.