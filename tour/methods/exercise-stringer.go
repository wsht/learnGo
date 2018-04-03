package methods

import (
	"fmt"
	"strings"
	"unsafe"
)

type IPAddr [4]byte

func (ip IPAddr) Stringer() string {
	fmt.Println(ip)
	str := ""

	for _, value := range ip {
		// fmt.Printf("%T", value)

		str = str + fmt.Sprintf("%v", value) + "."
	}
	str = strings.TrimRight(str, ".")
	fmt.Println(str)

	return str
}

func bytes2str(b byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ExerciseStringer() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
		a.Stringer()
	}
}
