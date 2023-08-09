package main

import (
	"fmt"
	"sync"
	"reflect"
)

func printy(message *string, wg *sync.WaitGroup){
	fmt.Println(*message)
	fmt.Println(message)
	*message = "oops"
	wg.Done()
}

func main() {
	var msg string = "hello bradmanfordson"
	var wg sync.WaitGroup
	fmt.Printf("%#v - %s\n", wg, reflect.TypeOf(wg))

	wg.Add(1)

	fmt.Println(&msg)
	fmt.Println(msg)
	fmt.Println("----------")
	go printy(&msg, &wg)
	wg.Wait()
	fmt.Println(msg)
}