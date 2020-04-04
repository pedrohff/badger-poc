# Development journal

## April 4th, 2020

`[2pm]`  
As I'm having lots of stuff to implement, structure and test, came to mind an idea of writing out the ideas passing through the development of this proof of concept.

For now, the structure of the code seems to be exactly what I would use in a production system, it lacks only a few tweaks on how and when the data should be cached, and some more control or exposition over the cache layer.

Since the repository's database layer is a simple mock, there's been some trouble on proving that the cache layer will show performance improvements. I've tried adding some timers and tweaking it, also tried adding some code-complexity to make this database access mock perform slower than using cache - understanding I could achieve false positive results.

Besides, I felt stuck on developing the tests, it has been overwritten a lot so far, and I feel I'm getting in a comfortable spot, not only for structuring but for reading the results as well.

Therefore, my next move will be to connect a real database to the project, check how I should structure its data, and rework my tests, adding some more random access, to reduce any kind of bias, as the results are based on linear loops, with sequenced arrays. 