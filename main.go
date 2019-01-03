package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
	"io/ioutil"
)

func main() {
	banner()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Target: ")

	targetHostname, _ := reader.ReadString('\n')

	targetHostname = strings.Replace(targetHostname, "\n", "", -1)


	addr,err := net.LookupIP("ispycode.com")
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}


	mxs, err := net.LookupMX(targetHostname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting MX Lookup...")

	fmt.Printf("MX Records for %s: \n", targetHostname)

        // dig +short google.com mx
	for _, mx := range mxs {
		if(Contains(mx.Host, "CloudFlare")) {
			fmt.Println(Red(mx.Host))
		} else {
                        fmt.Println(Green(mx.Host))
		}
	}

	
}

func banner() {
    b, err := ioutil.ReadFile("banner.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}

func Contains(s, substr string) bool {
    return strings.Index(s, substr) >= 0
}
