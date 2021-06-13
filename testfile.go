package main

import (
	"fmt"
	"io/ioutil"
	"log"
)



func myfunc() {

	content, err := ioutil.ReadFile("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
