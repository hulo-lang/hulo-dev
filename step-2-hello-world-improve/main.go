package main

import (
	"fmt"
	"hulo-dev/generated"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// 输入文本
	input := "hello world;\nhello hulo;"

	// 创建输入流
	inputStream := antlr.NewInputStream(input)

	// 创建词法分析器
	lexer := generated.NewHelloLexer(inputStream)

	// 创建标记流
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	// 创建语法分析器
	parser := generated.NewHelloParser(tokenStream)

	// 解析程序
	tree := parser.Program()

	// 创建监听器
	listener := &HelloListener{}

	// 遍历语法树
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println("解析完成！")
}

// 自定义监听器
type HelloListener struct {
	*generated.BaseHelloListener
}

// 进入问候语句时调用
func (l *HelloListener) EnterGreeting(ctx *generated.GreetingContext) {
	name := ctx.ID().GetText()
	fmt.Printf("问候: %s\n", name)
}
