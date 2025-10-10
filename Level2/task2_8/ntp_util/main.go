package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const timeHost = "0.beevik-ntp.pool.ntp.org"

func main() {
	timeNow, err := ntp.Time(timeHost)
	if err != nil {
		println(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Time: %v\n", timeNow)
}
