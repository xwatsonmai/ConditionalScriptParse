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

	lex := parse.Parse([]byte(`("广东省" + "深圳市" != "广东省" + "深圳市") and ("广东省" + "云浮市" != "广东省" + "云浮市") and ( "广东省" + "深圳市" != "广东省" + "深圳市")`))
	fmt.Println("res", lex.GetResult())
	fmt.Println("err", lex.GetError())

	//a := &big.Rat{}
	//a.SetString("1")
	//b := []*big.Rat{}

}
