package main

import (
	"fmt"
	"github.com/xwatsonmai/ConditionalScriptParse/parse"
)

func main() {
	//in := bufio.NewReader(os.Stdin)
	//for {
	//	if _, err := os.Stdout.WriteString("> "); err != nil {
	//		log.Fatalf("WriteString: %s", err)
	//	}
	//	line, err := in.ReadBytes('\n')
	//	if err == io.EOF {
	//		return
	//	}
	//	if err != nil {
	//		log.Fatalf("ReadBytes: %s", err)
	//	}
	//	//fmt.Println("debug", string(line))
	//	//fmt.Println("line", len(line))
	//	//lex := &exprLex{line: line}
	//	//msParse(lex)
	//	lex := parse.Parse(line)
	//	fmt.Println("res", lex.GetResult())
	//	fmt.Println("err", lex.GetError())
	//	//strings.Index()
	//}

	lex := parse.Parse([]byte(`((""="{}" or ""="" or -1=0 or -1=2 or -1=3) and (1=1 or 1=3 or 1=5))`))
	fmt.Println("res", lex.GetResult())
	fmt.Println("err", lex.GetError())

	//a := &big.Rat{}
	//a.SetString("1")
	//b := []*big.Rat{}

}
