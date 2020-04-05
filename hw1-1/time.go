package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"

)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == nil {
		fmt.Println("Now is",time)
		return

	} else {
		_,_ = fmt.Fprintf(os.Stderr, "Error message: %v\n", err)
		os.Exit(1)
	}


}
