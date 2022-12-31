package main

import (
	"fmt"
	// "strconv"
)

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (addr IPAddr) String() string {
	// strings.Join(string(addr), ":")

	buf := ""

	// i0 := int(addr[0])
	// buf += strconv.Itoa(i0)
	// buf += ":"
	// i1 := int(addr[1])
	// buf += strconv.Itoa(i1)
	// buf += ":"
	// i2 := int(addr[2])
	// buf += strconv.Itoa(i2)
	// buf += ":"
	// i3 := int(addr[3])
	// buf += strconv.Itoa(i3)

	// for i, v := range addr {
	// 	if i != 0 {
	// 		buf += ":"
	// 	}
	// 	buf += strconv.Itoa(int(v))
	// }

	buf = fmt.Sprintf("%d:%d:%d:%d", addr[0], addr[1], addr[2], addr[3])

	return buf
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
