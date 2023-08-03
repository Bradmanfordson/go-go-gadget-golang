package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Nest struct {
	Birds string `json:"birds"`
	Frogs string `json:"frogs"`
}

type Data struct {
	Hi string `json:"hi"`
	What string `json:"what"`
	Who string `json:"who"`
	Mm []string `json:"mm"`
	Amimals Nest `json:"amimals"`
}

func main() {
	content, err := os.ReadFile("./talk_dirty_to_me.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}

	
	var yeet Data
	err = json.Unmarshal(content, &yeet)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n",yeet)
	}

	fmt.Println()

	fmt.Printf("%#v\n", string(yeet.Hi))
	fmt.Println(yeet.What)
	fmt.Println(yeet.Who)

	fmt.Println("-----------------")
	for i, v := range(yeet.Mm) {
		fmt.Println(i)
		fmt.Println(v)
	}

	fmt.Printf("%#v\n", yeet.Amimals)

	// var yoink Amimals
	// json.Unmarshal([]byte(content["amimals"]))
	// fmt.Println(yoink.Amimals)
	// for _, v := range(yoink.Amimals){
	// 	fmt.Println(v)
	// }
}