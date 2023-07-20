package main

import "testing"

func BenchmarkPrintList(b *testing.B) {
	var bList linkedList
	for i := 0; i < 1_000; i++ {
		bList.addNodeToEnd(i)
	}
	for i := 0; i < b.N; i++ {
		bList.printList()
	}
}

func BenchmarkAddNodeToEnd(b *testing.B) {
	var bList linkedList
	for i := 0; i < b.N; i++ {
		bList.addNodeToEnd(i)
	}
}

/* https://pkg.go.dev/testing#hdr-Benchmarks
benchmarks are executed by the "go test" command
when its -bench flag is provided.

to run the benchmark test
enter this at the command line

go test -bench .

for help, enter this at the command line
go help
go help testfunc
go help testflag


// -run none		doesn't run any tests
// -benchtime 3s	run enough iterations of each benchmark to take 3s
go test -run none -bench . -benchtime 3s

-bench regexp
            Run only those benchmarks matching a regular expression.
            By default, no benchmarks are run.
            To run all benchmarks, use '-bench .' or '-bench=.'.
            The regular expression is split by unbracketed slash (/)
            characters into a sequence of regular expressions, and each
            part of a benchmark's identifier must match the corresponding
            element in the sequence, if any. Possible parents of matches
            are run with b.N=1 to identify sub-benchmarks. For example,
            given -bench=X/Y, top-level benchmarks matching X are run
            with b.N=1 to find any sub-benchmarks matching Y, which are
            then run in full.


-benchtime t
          Run enough iterations of each benchmark to take t, specified
          as a time.Duration (for example, -benchtime 1h30s).
          The default is 1 second (1s).
          The special syntax Nx means to run the benchmark N times
          (for example, -benchtime 100x).


Second (s): The second is the base unit of time in the International System of Units (SI).
Decisecond (ds): 1 decisecond = 0.1 seconds.
Centisecond (cs): 1 centisecond = 0.01 seconds.
Millisecond (ms): 1 millisecond = 0.001 seconds. This unit of time is often used in computing to measure response time.
Microsecond (μs): 1 microsecond = 0.000001 seconds = 1×10^-6 seconds. This unit of time is often used in scientific research and in telecommunications.
Nanosecond (ns): 1 nanosecond = 0.000000001 seconds = 1×10^-9 seconds. This unit of time is frequently used in computing and in the timing of electronic circuits.
Picosecond (ps): 1 picosecond = 0.000000000001 seconds = 1×10^-12 seconds. This unit of time is used in ultrafast science, like in laser physics and in the measurement of molecular dynamics.

*/
