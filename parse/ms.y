%{
package parse

import (
        "math/big"
)

func setResult(l msLexer, v bool) {
	l.(*exprLex).result = v
}
%}
%union {
param  *big.Rat
strParam string
check bool
symbol string
}

%type <param> expr expr1 expr2 expr3
%type <strParam> exprstr exprstr1
%type <check> res item numItem strItem
%token '>' '<' '=' '!' '(' ')' '+' '-' '*' '/'
%token <param> PARAM
%token <strParam> STRPARAM
%token <symbol> NOT AND OR GE LE
%%
start:
	res
	{
	setResult(mslex,$1)
	}
res:
	item
	{
	$$ = $1
	}
|	item AND item
	{
	$$ = $1 && $3
	}
|	item NOT item
        {
        $$ = $1 != $3
        }
|	item '=' item
	{
	$$ = $1 == $3
	}
|	item OR item
	{
	$$ = $1 || $3
	}
item:
	numItem
	{
	$$ = $1
	}
|	strItem
	{
	$$ = $1
	}
strItem:
	exprstr '=' exprstr
	{
	$$ = $1 == $3
	}
|	exprstr NOT exprstr
 	{
 	$$ = $1 != $3
 	}
exprstr:
	exprstr1
	{
	$$ = $1
	}
|	exprstr1 '+' exprstr1
	{
	$$ = $1 + $3
	}
|	'(' exprstr ')'
	{
	$$ = $2
	}
exprstr1:
	STRPARAM
numItem:
	expr '>' expr
	{
	if $1.Cmp($3) == 1 {
	  $$ = true
	}else{
	  $$ = false
	}
	}
|	expr '<' expr
	{
	if $1.Cmp($3) == -1 {
	  $$ = true
	}else{
	  $$ = false
	}
	}
|	expr '=' expr
	{
	if $1.Cmp($3) == 0 {
	  $$ = true
	}else{
	  $$ = false
	}
	}
|	expr GE expr
	{
	check := $1.Cmp($3)
	if check == 1 || check == 0{
	$$ = true
	}else{
	$$ = false
	}
	}
|	expr LE expr
	{
	check := $1.Cmp($3)
	if check == -1 || check == 0{
	$$ = true
	}else{
	$$ = false
	}
	}
|	expr NOT expr
        {
       	if $1.Cmp($3) != 0 {
       	  $$ = true
       	}else{
       	  $$ = false
       	}
        }
|	'(' item ')'
	{
	$$ = $2
	}
expr:
	expr1
|	'-'expr
	{
	$$ = $2.Neg($2)
	}
expr1:
	expr2
|	expr1 '+' expr2
	{
		$$ = $1.Add($1, $3)
	}
|	expr1 '-' expr2
	{
		$$ = $1.Sub($1, $3)
	}
expr2:
	expr3
|	expr2 '*' expr3
	{
		$$ = $1.Mul($1, $3)
	}
|	expr2 '/' expr3
	{
		$$ = $1.Quo($1, $3)
	}

expr3:
	PARAM
|	'(' expr ')'
	{
		$$ = $2
	}
%%