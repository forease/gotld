package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/forease/gotld"
)

var url = flag.String("d", "", "check url")

func main() {
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: -d www.forease.net")
		os.Exit(1)
	}

	tld, domain, err := gotld.GetTld(*url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("TLD: %s, Domain: %s\n", tld.Tld, domain)
}
