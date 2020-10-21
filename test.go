// example1.go
package main

import (
  "os"
	"fmt"
	mail "net/mail"
	forkedmail "./go/src/net/mail"
	goawaymail "./net/mail"
)

func main() {
	addressList, err := mail.ParseAddressList(os.Args[1])
	fmt.Printf("GO: %+v %v\n", addressList, err)
	addressList2, err2 := forkedmail.ParseAddressList(os.Args[1])
	fmt.Printf("GP: %+v %v\n", addressList2, err2)
	addressList3, err3 := goawaymail.ParseAddressList(os.Args[1])
	fmt.Printf("GA: %+v %v\n", addressList3, err3)
}
