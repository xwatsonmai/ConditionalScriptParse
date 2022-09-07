# 条件表达式脚本（DSL）解析
## 一、项目特定
使用goyacc，lex进行解析的脚本，代码易看，容易拓展。
## 二、项目目标
给完全不懂编程的人士，提供一套开箱即用的，配置流程或指标条件判断的脚本。省去复杂的表单配置与解析过程。减少配置人员与开发人员的心智负担，减轻系统的复杂度
## 三、目前支持的特性
### 1. 数字类型的条件表达式比较
### 2. 四则运算
### 3. 字符串类型的对比
### 4. 字符串拼接对比

## 四、未来目标
### 1. 字符串比较支持 （已实现）
### 2. 代码内调用支持 （已实现）
### 3. 完善的测试用例 （已实现）
### 4. 对应的前端脚本编写组件 （咕咕咕）

## 五、使用
### 1.拉包
```shell
 go get -u github.com/xwatsonmai/ConditionalScriptParse/parse@v1.0.1 
```
### 2.在代码中使用

```golang
package main

import (
	"fmt"
	"github.com/xwatsonmai/ConditionalScriptParse/parse"
	"strconv"
	"strings"
)

func main() {
	// 假设aVal，bVal为接口请求返回参数
	aVal := 1
	bVal := 2

	// 配置好的条件表达式
	checkStr := "{aVal} > {bVal}"

	// 配置好的表达式变量关系
	checkStr = strings.ReplaceAll(checkStr, "{aVal}", strconv.Itoa(aVal))
	checkStr = strings.ReplaceAll(checkStr, "{bVal}", strconv.Itoa(bVal))
	line := []byte(checkStr)
	lex := parse.Parse(line)
	err := lex.GetError()
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	res := lex.GetResult() // bool
	fmt.Println("result is ", res)
}
```

## 六、基础符号
### 1.四则运算：+，-，*，/
### 2.范围判断「(」... 「)」
### 3.并:and ，或:or
### 4.等于: =
### 5.字符串需要是'"'进行包裹


## 用例：
```shell
> 1 > 1
false
> 1 =  1
true
> (1+1)>(2+2)
false
> (1=1) or (2>2)
true
> (1=1) and (2>2)
false
> "你好" = "你好"
true
```
* 更多的测试用例可在parse/lexer_test.go中查看