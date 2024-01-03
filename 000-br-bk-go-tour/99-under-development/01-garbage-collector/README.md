
# "If I miss one day of practice, I notice it. If I miss two days, the critics notice it. If I miss three days, the audience notices it"
#### ~ Ignacy (Jan) Paderewski, a French philosopher and mathematician. 

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. s

# Table of Contents

1. [Re](#re)


# Hardware architecture and Go

- Machines have CPUs
- CPUs have cores
- Cores have hardware threads

For example, a computer could have 1 CPU with 4 cores and 2 threads per core. This CPU would thus have 8 threads available for execution. Each thread is seen by Go as a **virtual core.** When you run this code, you will see the number of **virtual cores** available:

```go
fmt.Println(runtime.NumCPU())
``` 

### Diagrams

Notation is used in diagrams to depict how Go works with hardware:

- M: OS thread on the **M**achine   (8 in our example)
- P: virtual core **P**rocessor     (8 in our example)
- G: **G**oroutine                  

The nice thing about those letters is that they form the acronym MPG. Combined, those letters help determine the performance of your Go program: the OS threads on the machine (M) are seen as virtual core processors (P) by Go, and then goroutines (G) are assigned to each of those processors (P).

**Important to know: Multiple goroutines can be assigned to each processor.**

### The Go scheduler

There are two ways Goroutines (G) are assigned to the processors (P). This is also known as "scheduling" goroutines, and the "Go scheduler" takes care of this scheduling during runtime. This scheduling is done using run queues. Here are the two run queues that are used: 

- GRQ (global run queue)
- LRQ (local run queue)

The LRQ is attached to a processor (P). 

Goroutines wait in the GRQ until they are assigned to an LRQ. In the LRQ, the LRQ determines which goroutine should execute.

***"The Go scheduler is part of the Go runtime, and the Go runtime is built into your application."***
[Bill Kennedy](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)

### Goroutine states

- waiting
- runnable
- executing

```
Waiting: the Goroutine is stopped and waiting for something: waiting for the OS system calls such as file-based calls, or synchronization events such as atomic and mutex operations. These latencies are a root cause for bad performance.

Runnable: the Goroutine wants time on an M so it can run its instructions. If you have a lot of Goroutines that want time, then Goroutines have to wait longer to get time. Also, the individual amount of time any given Goroutine gets is shortened as more Goroutines compete for time. This scheduling latency can also be a cause of bad performance.

Executing: This means the Goroutine has been placed on an M and is executing its instructions.

[Bill Kennedy](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
```

# Switching between goroutines aka context switching

- Context switching: when the scheduler takes one goroutine off an M and puts another one on it.



# Code review check

### R
- If 


# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# "The very essence of success is practice."
#### ~ Ignacy (Jan) Paderewski

![Ignacy (Jan) Paderewski](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/06-garbage-collector/images/Ignacy_Jan_Paderewski_-_Project_Gutenberg_eText_15604.png)

source: https://upload.wikimedia.org/wikipedia/commons/d/d4/Ignacy_Jan_Paderewski_-_Project_Gutenberg_eText_15604.png