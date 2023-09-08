package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)
}

/*
Analyze this file using the `pprof` tool:

```bash
go tool pprof mem.pprof
```
*/