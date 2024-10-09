package parse

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/big"
)

type exprLex struct {
	line   []byte
	pos    int
	st     bool
	err    error
	result bool
}

const eof = 0

func (x *exprLex) Lex(yylval *msSymType) int {
	if len(x.line) == 0 {
		return 0 //表示已经解析完所有的内容了
	}
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, yylval)
		case 'a', 'o', 'i':
			return x.symbol(c, yylval)
		case '"':
			str, err := x.string(c, yylval)
			if err != nil {
				x.err = err
				return 0
			}
			return str
		case '!':
			if x.peek() == '=' {
				x.pos++
				return NOT
			} else {
				return int(c)
			}
		case '<', '>':
			if x.peek() == '=' {
				x.pos++
				if c == '<' {
					return LE
				} else {
					return GE
				}
			} else {
				return int(c)
			}
		case '=', '-', '(', ')', '&', '+', '*', '/':
			if c == '=' {
				if x.last() != '!' {
					return int(c)
				}
			} else if c == '&' {
				return AND
			} else {
				return int(c)
			}
		case ' ', '\t', '\n', '\r':
		default:
			x.err = errors.New(fmt.Sprintf("unrecognized character %q", c))
			return 0
		}
	}
}

func (x *exprLex) Error(s string) {
	if x.err == nil {
		x.err = errors.New(fmt.Sprintf("parse error: %s", s))
	} else {
		x.err = fmt.Errorf("parse error: %s,%w", s, x.err)
	}
	//x.err = errors.New(fmt.Sprintf("parse error: %s", s))
	//log.Printf("parse error: %s", s)
}

func (x *exprLex) peek() rune {
	if x.pos == len(x.line) {
		return eof
	}
	return rune(x.line[x.pos])
}

func (x *exprLex) last() rune {
	//fmt.Println("last", x.pos)
	return rune(x.line[x.pos-2])
}

func (x *exprLex) next() rune {
	if x.pos == len(x.line) {
		return eof
	}
	defer func() {
		x.pos++
	}()
	return rune(x.line[x.pos])
}

func (x *exprLex) num(c rune, yylval *msSymType) int {
	//_p := x.pos
	//fmt.Println("_p", _p)
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			x.err = errors.New(fmt.Sprintf("WriteRune: %s", err))
			//log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
L:
	for {
		c = x.peek()

		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			x.pos++
			add(&b, c)
		default:
			break L
		}
	}
	yylval.param = &big.Rat{}
	_, ok := yylval.param.SetString(b.String())
	if !ok {
		x.err = errors.New(fmt.Sprintf("bad number %q", b.String()))
		//log.Printf("bad number %q", b.String())
		return eof
	}
	return PARAM
}

func (x *exprLex) symbol(c rune, yylval *msSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
L:
	for {
		c = x.peek()
		switch c {
		case 'n', 'd', 'r':
			//fmt.Println("c", string(c))
			x.pos++
			add(&b, c)
		default:
			break L
		}
	}
	if b.String() == "and" {
		return AND
	} else if b.String() == "or" {
		return OR
	} else if b.String() == "in" {
		return IN
	} else {
		log.Printf("bad symbol %q", b.String())
		return eof
	}
}

func (x *exprLex) string(c rune, yylval *msSymType) (int, error) {
	//add := func(b *bytes.Buffer, c rune) {
	//	if _, err := b.WriteRune(c); err != nil {
	//		log.Fatalf("WriteRune: %s", err)
	//	}
	//}
	escaped := false
	var b []rune
L:
	for {
		c = x.next()
		if c == eof {
			return eof, errors.New(fmt.Sprintf("bad string end %q", string(b)))
		}
		if escaped {
			b = append(b, c)
			escaped = false
			continue
		}
		switch c {
		case '\\':
			escaped = true
		case '"':
			break L
		default:
			b = append(b, c)
		}
	}
	yylval.strParam = string(b)
	return STRPARAM, nil
}

//func SetResult(res bool) {
//	fmt.Println("res", res)
//	x.result = res
//}

func (x *exprLex) GetResult() bool {
	return x.result
}

func (x *exprLex) GetError() error {
	return x.err
}

func Parse(line []byte) *exprLex {
	x := &exprLex{line: line}
	msParse(x)
	return x
}
