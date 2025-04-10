# Greek Calculator Solution in Go
Solving Greek Calculator puzzle with Golang

<img width="489" alt="calculator_puzzle" src="https://github.com/user-attachments/assets/63c9eb52-4d9c-48dc-97f2-37f946c10cc6" style="display: block; margin: 0 auto" />

## What is the Greek Calculator Puzzle?
The Greek Calculator Puzzle also known as Grecian Computer presents a device where one has to figure out how to adjust five different dials with numbers on each of them to add up to 42 in each column.

## Story behind the solution
When I first got the puzzle I thought it wouldn't take too long to solve it by hand. After realizing that it wasn't going to be as straightforward as possible, I took a pen and a notebook and noted down everything I could learn about this puzzle to see if I could narrow down the combinations or find certain patterns. That didn't work too well for me. After a while, I thought that this puzzle could be solved programmatically. My first attempts at solving this involved writing a program in JavaScript to calculate all possible combinations of the numbers available that add up to 42 to see how many combinations I would get, after which I was going to eliminate the ones that are not possible with this puzzle. That yielded way too many results. I decided to give up on this puzzle.

This is where I ended up picking it up again earlier this year. I started learning Go and thought that it would be a great opportunity to practice Go and try to solve this puzzle again. This time my approach was to completely represent the dials using multidimensional slices in Go to paint an accurate picture of the puzzle and simulate it in the program. After tinkering with it for a while, I came up with functions that would rotate rows, and dials, calculate column sums, and layer all of the dials on top of each other to see if they add to the desired sums. To my surprise, this approach worked on the first try. I was able to verify the layout of the physical puzzle which ended up being the correct solution for the puzzle.

In addition to the puzzle, I added tests and benchmarks for most of the functions that comprised the program. It was a fun way to explore Go features. In particular, I have learned a lot about Go's slices, arrays, testing, and benchmarking.

## Getting Started
### Prerequisites
You should have `go@1.24.*` installed for this project to run. Tests will only run with this version and higher due to the introduction of `b.Loop()` feature for benchmarking

### Running
To run this program locally, clone this repository, navigate to the folder, and run the main.go file:
```bash
git clone https://github.com/EvgeniiKlepilin/greek-calculator-go.git
cd greek-calculator-go
go run main.go
```

### Testing
To run tests for the calculator, you should run the standard Go test command:
```bash
go test
```
A sample output will look like the following:
```bash
PASS
ok  	evgeniiklepilin.com/greek-calculator-go	0.198s
```

To get more verbose output, run it with `-v` flag:
```bash
go test -v
```
A sample output will look like this:
```bash
=== RUN   TestLayerDials
--- PASS: TestLayerDials (0.00s)
=== RUN   TestCalculateColumn
--- PASS: TestCalculateColumn (0.00s)
=== RUN   TestRotateRow
--- PASS: TestRotateRow (0.00s)
=== RUN   TestRotateDial
--- PASS: TestRotateDial (0.00s)
PASS
ok  	evgeniiklepilin.com/greek-calculator-go	0.129s
```

### Benchmarking
To run benchmark tests on the program, run the following command:
```bash
go test -bench=.
```
A sample output will look like this:
```bash
goos: darwin
goarch: amd64
pkg: evgeniiklepilin.com/greek-calculator-go
cpu: Intel(R) Core(TM) i7-7567U CPU @ 3.50GHz
BenchmarkBruteForce-4        	     138	   8968762 ns/op
BenchmarkLayerDials-4        	 2944935	       350.3 ns/op
BenchmarkCalculateColumn-4   	221360473	         5.417 ns/op
BenchmarkRotateRow-4         	207250884	         5.833 ns/op
BenchmarkRotateDial-4        	60783166	        19.52 ns/op
PASS
ok  	evgeniiklepilin.com/greek-calculator-go	6.000s
```

