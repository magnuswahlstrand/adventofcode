# Advent of Code - 2018

Entries to the [advent of code 2018](https://adventofcode.com/2018).

## Time spent

| Day | Go     |        |           | Node   |        |         | Comment                                                                       |
|-----|--------|--------|-----------|--------|--------|---------|-------------------------------------------------------------------------------|
| Day | Part 1 | Part 2 | Total     | Part 1 | Part 2 | Total   |                                                                               |
| 1   |        |        | ?         |        |        | ?       |                                                                               |
| 2   | 13m    | 32m    | **45m**   | 32m    | 15m    | **47m** | Part 2 quick in JS, since already solved in Go                                |
| 3   | 43m    | 35m    | **1h18m** |        |        |         |                                                                               |
| 4   | 1h15m  | 4m     | **1h19m** |        |        |         | Forgot to sort list, had to implement a few workarounds because of weird data |
| 5   | 29m    | 12m    | **41m**   |        |        |         | A few corner cases. New record for fastest Go.                                |
| 6   | 50m    | 14m    | **1h04**  |        |        |         |                                                                               |
| 7   | 35m    | 55m    | **1h30**  |        |        |         | Hard!                                                                         |
| 8   | 45m    | 1h15   | **2h**    |        |        |         | Clearly approaching the limit of my algorithm knowledge :-D. Fun though.      |
| 9   | 45m    | 2h     | **2h45**  |        |        |         | First time profiling a Go program. First time using linked lists.             |
| 10  | 50m    | 5m     | **55m**   |        |        |         | Yeah! Pretty quick this time                                                  |
| 11  |        |        | ****      |        |        |         |                                                                               |
| 11  |        |        | ****      |        |        |         |                                                                               |
| 11  |        |        | ****      |        |        |         |                                                                               |
| 11  |        |        | ****      |        |        |         |                                                                               |
| 11  |        |        | ****      |        |        |         |                                                                               |


### Lessons leartn
1) `Linked lists` is a thing
2) `Pprof` is a thing
3) The [standard lib](https://golang.org/pkg/#stdlib) has implementations of  `ring`,``heap`,`list`.
4) Turn on GC tracing with GODEBUG=gctrace=1
5) Turn of GC with GOGC=-1
6) Printing negative and positive integers the same width umping `fmt.Printf("% 3d ->% 3d",...` -->" -11 ->123" 
7) Use slices as map keys: `map[[0 2 7 0]:1 [0 2 6 0]:2]` 
