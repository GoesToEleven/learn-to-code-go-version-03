package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

func bidOn(ctx context.Context, url string) Bid {

	// with this: ch := make(chan Bid)
	// this is a bug: a goroutine leak
	// run the profiler and see how many goroutines you see
	// ch := make(chan Bid)

	// fix the bug with a buffered channel of 1
	ch := make(chan Bid, 1)

	go func() {
		bid := bestBid(url)
		ch <- bid
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

var defaultBid = Bid{
	AdURL: "http://adsRus.com/default",
	Price: 3,
}

func bestBid(url string) Bid {
	// simulate work
	d := 100 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsЯus.com/ad17",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int // In ¢
}
