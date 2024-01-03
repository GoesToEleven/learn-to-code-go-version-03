# "Drop by drop, the bucket gets filled."

# Video
- [video #1 on youtube](https://youtu.be/HBtu9Dsjp80)
- [video #2 on youtube](https://youtu.be/ut_REn0xFPM)

# Takeaways
1. Monotonic time
2. Time unites: second, decisecond, centisecond, millisecond, microsecond, nanosecond, picosecond
3. UTC
4. super fun code with bufio.NewScanner() for `formatting time` (see the code base)
5. [SOON TO BE RELEASED described in this video]() `01/02 03:04:05PM '06 -0700` 

# Table of Contents
1. [Monotonic time](#monotonic-time)
2. [Time Units](#time-units)
3. [UTC](#utc)
4. [Coupons for Go courses](#coupons-for-go-courses)

# Monotonic time
In computing, a clock is considered "monotonic" if it ***only ever advances forward*** in time, never backward. A monotonic clock guarantees that the values it returns will be strictly increasing over time, which is **useful for measuring elapsed time between events in a reliable way.** However, unlike real-time clocks, monotonic clocks are not set to any particular standard time and may not persist their values across reboots.

In Go, the `time` package provides a way to access a monotonic clock through its `time.Now()` function. The `time.Time` struct returned by this function embeds both a wall clock time and a monotonic clock reading. **If you take the difference between two `time.Time` values, Go will use the monotonic clock reading to compute the duration, if available.**

Here's an example that demonstrates how to measure elapsed time using the monotonic clock:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()  // Get the current time, including monotonic clock reading

	// Simulate some work
	time.Sleep(2 * time.Second)

	end := time.Now()  // Get the current time again

	elapsed := end.Sub(start)  // Compute the elapsed time using monotonic clock

	fmt.Printf("Elapsed time: %s\n", elapsed)
}
```
### Monotonic time continued
In this example, `start` and `end` are `time.Time` structs that contain both the wall time and the monotonic time. When `end.Sub(start)` is computed, it uses the monotonic clock readings to accurately measure the time elapsed, even if the system clock has been adjusted during that interval.

Using a monotonic clock is important for reliable time measurement because wall clock time can be adjusted backward or forward, for example, when synchronizing with a time server or due to daylight saving time adjustments. These adjustments could lead to incorrect or even negative duration calculations if you were only relying on the wall clock time.

# Time Units
- **Second (s)**: The second is the base unit of time in the International System of Units (SI).
- **Decisecond (ds)**: 1 decisecond = 0.1 seconds.
- **Centisecond (cs)**: 1 centisecond = 0.01 seconds.
- **Millisecond (ms)**: 1 millisecond = 0.001 seconds. This unit of time is often used in computing to measure response time.
- **Microsecond (μs)**: 1 microsecond = 0.000001 seconds = **1×10^-6 seconds.** This unit of time is often used in scientific research and in telecommunications.
- **Nanosecond (ns)**: 1 nanosecond = 0.000000001 seconds = **1×10^-9 seconds.** This unit of time is - frequently used in computing and in the timing of electronic circuits.
- **Picosecond (ps)**: 1 picosecond = 0.000000000001 seconds = **1×10^-12 seconds.** This unit of time is used in ultrafast science, like in laser physics and in the measurement of molecular dynamics.

# UTC
UTC stands for Coordinated Universal Time, and it serves as the primary time standard by which the world regulates clocks and time. Although the acronym "UTC" doesn't align with either the English or French phrases for Coordinated Universal Time, the name was chosen as a compromise to accommodate various languages.

### Historical Background
Before the widespread use of UTC, Greenwich Mean Time (GMT) was the standard. GMT was originally based on the mean solar time at the Prime Meridian, which passes through Greenwich, London. However, GMT is now considered somewhat outdated and has been largely replaced by UTC, which is more accurate and precise. UTC was formally introduced in 1960 as a modern replacement for GMT and has been updated several times since to improve its accuracy.

### Technical Basis
UTC is an atomic time scale, based on International Atomic Time (TAI) with leap seconds added at irregular intervals to compensate for the slowing rotation of the Earth. These leap seconds are added to ensure that UTC remains within 0.9 seconds of mean solar time, also known as UT1. The addition of leap seconds is decided by international agreement and is announced several months in advance. This is why UTC is often considered more precise than GMT: it incorporates the benefits of both atomic timekeeping and astronomical observations.

### Time Zones
UTC serves as the reference for time zones worldwide. Local time zones are defined as UTC plus or minus a certain number of hours and minutes. For example, Eastern Standard Time (EST) is UTC-5, meaning it is 5 hours behind UTC. Some regions also observe daylight saving time, temporarily shifting their local time.

### UTC in Computing
In computing and telecommunications, UTC is crucial for data synchronization. Many protocols and standards, including those for the internet, are based on UTC. When computers synchronize their clocks via the Network Time Protocol (NTP), they are adjusting to a time source that is coordinated with UTC.

### Converting UTC to UTC-5 (EST)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a time in UTC
	utcTime := time.Date(2023, time.September, 2, 14, 0, 0, 0, time.UTC)
	fmt.Printf("UTC time: \t %v \n", utcTime)

	// Define EST (UTC-5)
	est := time.FixedZone("EST", -5*60*60)

	// Convert the time to EST
	estTime := utcTime.In(est)
	fmt.Printf("Time in EST: \t %v \n", estTime)
}
```

### Leap Seconds and Computing
The addition of leap seconds can pose challenges for computing systems, especially those requiring time to always move forward. To deal with this, some systems use a modified version of UTC, such as "smoothed" leap seconds or simply ignoring them altogether, depending on the needs of the application.

### Summary
UTC is the standard for timekeeping used worldwide, based on both atomic time and astronomical observations. It serves as the basis for local time zones and is crucial for various aspects of modern life, including computing, telecommunications, and international coordination.

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce) 
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses


![global utc time zones](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/02-time-pkg/UTC-time-zones.png)
source: https://en.wikipedia.org/wiki/Time_zone#/media/File:World_Time_Zones_Map.png