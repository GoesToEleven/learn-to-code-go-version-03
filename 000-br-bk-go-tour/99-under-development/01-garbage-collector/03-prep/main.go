package main

import (
	"fmt"
	"math"
)

func main() {
	// 3 GHz = 3_000_000_000 ops / sec
	ghz := 3_000_000_000.0

	fmt.Println("ops per deci")
	fmt.Println(ghz/(10))

	fmt.Println("ops per centi")
	fmt.Println(ghz/(math.Pow(10,2)))

	fmt.Println("ops per milli")
	fmt.Println(ghz/(math.Pow(10,3)))

	fmt.Println("ops per micro")
	fmt.Println(ghz/(math.Pow(10,6)))

	fmt.Println("ops per nano")
	fmt.Println(ghz/(math.Pow(10,9)))

	fmt.Println("ops per pico")
	fmt.Println(ghz/(math.Pow(10,12)))
}

/*
The notation "3e+06" represents a numerical value in scientific notation. 
It is a compact way to express very large numbers or very small numbers, 
particularly in fields like science and engineering. In this notation:

The "3" is known as the coefficient or significand.
The "e" stands for "exponent."
The "+06" is the exponent value.
So, "3e+06" can be interpreted as follows:

3 * 10^6

In other words, it represents the number 3 followed by six zeros, which is equivalent to 3,000,000.
*/

/*
Context switches are considered to be expensive because it takes times to swap Threads on and off a core. 
The amount of latency incurrent during a context switch depends on different factors but it’s not unreasonable 
for it to take between ~1000 and ~1500 nanoseconds. Considering the hardware should be able to reasonably execute 
(on average) 12 instructions per nanosecond per core, a context switch can cost you ~12k to ~18k instructions of latency. 
In essence, your program is losing the ability to execute a large number of instructions during a context switch.
*/
// source: https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html

/*
One CPU can have multiple cores.

The number of instructions per second that can be executed per CPU core depends on several factors, including the specific CPU architecture, clock speed, and the type of instructions being executed. Generally, modern CPUs are capable of executing multiple instructions per clock cycle (IPC), and the clock speed is typically measured in gigahertz (GHz) which is billions of clock cycles per second.

GHz
measure of clock speed
clock cycles per second measured in billions
example: 3.5 GHz

Hz
measure of clock speed
clock cycles per second
example: 3_500_000_000 Hz

Note:
3.5 GHz is the same as 3_500_000_000 Hz

IPC
instructions per clock cycle

To calculate the maximum theoretical instructions per second per CPU core, you can follow these steps:

1. Find the clock speed of the CPU in Hz (hertz). For example, if you have a CPU with a clock speed of 3.5 GHz, you would convert it to Hz:

   3.5 GHz = 3,500,000,000 Hz

2. Determine the IPC of the CPU. IPC represents the number of instructions that can be executed per clock cycle. This value varies depending on the CPU architecture and workload. Let's assume an IPC of 2 as a rough estimate.

3. Calculate the maximum instructions per second (INS/s) per core:

   Maximum INS/s = (Clock Speed in Hz) * (IPC)

   Maximum INS/s = (3,500,000,000 Hz) * (2) = 7,000,000,000 INS/s

So, in this example, you can theoretically execute up to 7 billion instructions per second per CPU core.

Please note that this is a simplified calculation and represents a theoretical maximum. Real-world performance can vary significantly depending on factors like the complexity of the instructions, memory access patterns, branch prediction, and many other architectural considerations. Additionally, not all instructions take the same amount of time to execute, so the actual throughput may be lower than the theoretical maximum.
*/
