package main

import (
	"io/ioutil"
	"fmt"
	"reflect"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)

	files, err := ioutil.ReadDir("/usr/bin")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for _, file := range files {
			fmt.Printf("%#v  -  %s\n", file, reflect.TypeOf(file))
		}
		wg.Done()
	}()

	wg.Wait()
}