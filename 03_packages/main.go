package main

import (
	"fmt"
	"math"
	"github.com/luanrivello/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println("packages", math.Ceil(math.Abs(-2.1)))
	fmt.Println("packages", math.Abs(math.Ceil(-2.1)))
	fmt.Println(strutil.Reverse("packages"))
}
