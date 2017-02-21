package pkg

import (
	"fmt"
	"log"
)

func fn() {
	var s string
	fn2 := func() string { return "" }
	_, _ = fmt.Printf(fn2()) // MATCH /should use print-style function/
	fmt.Sprintf(fn2())       // MATCH /should use print-style function/
	log.Printf(fn2())        // MATCH /should use print-style function/
	_, _ = fmt.Printf(s)     // MATCH /should use print-style function/
	_, _ = fmt.Printf(s, "")

	_, _ = fmt.Printf(fn2(), "")
	_, _ = fmt.Printf("")
	_, _ = fmt.Printf("", "")
}
