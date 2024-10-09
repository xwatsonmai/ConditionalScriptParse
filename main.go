package main

import (
	"fmt"
	"github.com/xwatsonmai/ConditionalScriptParse/parse"
)

func main() {
	line := []byte(`"是"="是"`)
	result := parse.Parse(line)
	if err := result.GetError(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result.GetResult())
	}

	//a := &big.Rat{}
	//a.SetString("1")
	//b := []*big.Rat{}

}
