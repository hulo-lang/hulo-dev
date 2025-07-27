// Code generated from Hello.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // Hello
import "github.com/antlr4-go/antlr/v4"

// BaseHelloListener is a complete listener for a parse tree produced by HelloParser.
type BaseHelloListener struct{}

var _ HelloListener = &BaseHelloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHelloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHelloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHelloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHelloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseHelloListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseHelloListener) ExitProgram(ctx *ProgramContext) {}

// EnterGreeting is called when production greeting is entered.
func (s *BaseHelloListener) EnterGreeting(ctx *GreetingContext) {}

// ExitGreeting is called when production greeting is exited.
func (s *BaseHelloListener) ExitGreeting(ctx *GreetingContext) {}
