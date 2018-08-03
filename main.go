package main

import (
	"fmt"
	"os"
)

func main() {
	host := os.Getenv("SITEHOST")
	fmt.Printf("host: %s", host)
}
