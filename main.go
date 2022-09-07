package main

import (
	"ConditionalScriptParse/parse"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}
		//fmt.Println("debug", string(line))
		//fmt.Println("line", len(line))
		//lex := &exprLex{line: line}
		//msParse(lex)
		lex := parse.Parse(line)
		fmt.Println("res", lex.GetResult())
		fmt.Println("err", lex.GetError())
	}

	//msParse(&exprLex{line: []byte("11 < 22")})
}
