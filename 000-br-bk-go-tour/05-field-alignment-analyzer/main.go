// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(fieldalignment.Analyzer) }

// source:
// https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment

/*
to run this and have it analyze a file hypothetically called 'main2.go'
do this at the command line:

mac
./main <path/to/main2.go>

windows
main.exe <path/to/main2.go>
*/
