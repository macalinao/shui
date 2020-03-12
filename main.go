package main

import (
	"fmt"
	"time"

	"github.com/turnage/graw/reddit"
)

func main() {
	script, err := reddit.NewScript("shui v0.1", 2*time.Second)
	if err != nil {
		panic(err)
	}

	harvest, err := script.ListingWithParams("/r/hydrohomies/top", map[string]string{"t": "day"})
	if err != nil {
		panic(err)
	}

	fmt.Print(harvest.Posts[0].URL)
}
