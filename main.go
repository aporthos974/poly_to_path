package main

import (
	"fmt"
	"io/ioutil"
	"polygon_to_path/converter"
)

func main() {
	content, _ := ioutil.ReadFile("reunion.svg")
	fmt.Printf(converter.Convert(string(content)))
}
