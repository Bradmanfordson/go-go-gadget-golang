package main

import (
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

var nums = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var letters = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func print_nums() {
	for i := 0; i < 10; i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("print_nums pid", unix.Getpid())
	fmt.Println("print_nums tid", unix.Gettid())

}

func print_letters() {
	for i := 0; i < 10; i++ {
		fmt.Println(letters[i])
	}
	fmt.Println("print_letters pid", unix.Getpid())
	fmt.Println("print_letters tid", unix.Gettid())

}

func main() {

	go print_nums()

	fmt.Println("main pid", unix.Getpid())
	fmt.Println("main tid", unix.Gettid())

	go print_letters()

	for i := 4; i > 0; i-- {
		fmt.Println("Finished in: ", i, " seconds.")
		time.Sleep(1 * time.Second)
	}
}
    
