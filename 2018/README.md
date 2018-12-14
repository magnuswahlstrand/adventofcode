# Advent of Code - 2018

Entries to the [advent of code 2018](https://adventofcode.com/2018).

## Time spent

| Day | Go     |        |          | Comment                                                                       |
| --- | ------ | ------ | -------- | ----------------------------------------------------------------------------- |
| Day | Part 1 | Part 2 | Total    |                                                                               |
| 1   |        |        | ?        |                                                                               |
| 2   | 13m    | 32m    | **45m**  | Part 2 quick in JS, since already solved in Go                                |
| 3   | 43m    | 35m    | **1h18** |                                                                               |
| 4   | 1h15m  | 4m     | **1h19** | Forgot to sort list, had to implement a few workarounds because of weird data |
| 5   | 29m    | 12m    | **41m**  | A few corner cases. New record for fastest Go.                                |
| 6   | 50m    | 14m    | **1h04** |                                                                               |
| 7   | 35m    | 55m    | **1h30** | Hard!                                                                         |
| 8   | 45m    | 1h15   | **2h**   | Clearly approaching the limit of my algorithm knowledge :-D. Fun though.      |
| 9   | 45m    | 2h     | **2h45** | First time profiling a Go program. First time using linked lists.             |
| 10  | 50m    | 5m     | **55m**  | Yeah! Pretty quick this time                                                  |
| 11  | 1h30   | 15m    | **1h45** | Wrong format + 30m. Part 2 very slow, but convereges in seconds.              |
| 12  | 45m    | 35m    | **1h20** | Not bad, messed up a bit with counting the generations, and the padding       |
| 13  | 1h30   | 45m    | **2h15** | Forgot to sort inside loop, +30 min on part 2                                 |
| 14  | 1h03m  | 42m    | **1h45** |                                                                               |
| 11  |        |        | ****     |                                                                               |
| 11  |        |        | ****     |                                                                               |
| 11  |        |        | ****     |                                                                               |
| 11  |        |        | ****     |                                                                               |

### Lessons learnt

1) `Linked lists` is a thing
2) `Pprof` is a thing
3) The [standard lib](https://golang.org/pkg/#stdlib) has implementations of  `ring`,``heap`,`list`.
4) Turn on GC tracing with GODEBUG=gctrace=1
5) Turn of GC with GOGC=-1
6) Printing negative and positive integers the same width umping `fmt.Printf("% 3d ->% 3d",...` -->" -11 ->123"
7) Use slices as map keys: `map[[0 2 7 0]:1 [0 2 6 0]:2]`  
8) VsCode: Set bookmark `cmd + shift + <N>`, go to bookmark `cmd + <N>`
9) VsCode: Previous cursor location `ctrl + -`, next cursor location `ctrl + shift + -`
10) `fmt.Sscanf` is very convenient for reading formatted input of various types
    1) Example: `fmt.Sscanf("string 10 -10", "%s %d %d",...)`
11) [Stringer](https://tour.golang.org/methods/17) interface
12) [Gonum](https://github.com/gonum/gonum) for Matrix manipulation (`matrix`,`sum`,`...`)
13) [sort.Slice](https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field) for slices of structs