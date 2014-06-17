package main

import "fmt"

func noError(e error) {
	if e != nil {
		panic(e)
	}
}

func log(label string, obj ...interface{}) {
	fmt.Println(label)
	for _, o := range obj {
		fmt.Println(o)
	}
}
