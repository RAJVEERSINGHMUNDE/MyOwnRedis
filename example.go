package main

import (
	"fmt"
)

func main() {

	const n = 500000000
	const d = 3e20 / n
	fmt.Print(int64(d))

}
