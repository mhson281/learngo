package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var (
        sum int
        i = min
    )

	for {
		if i % 2 != 0 {
            i++
			continue
		} else if i > max {
            break
        }
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
        i++
	}
	fmt.Printf(" = %d\n", sum)
}

