// Code generated from Hello.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // Hello
import "github.com/antlr4-go/antlr/v4"

// HelloListener is a complete listener for a parse tree produced by HelloParser.
type HelloListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterGreeting is called when entering the greeting production.
	EnterGreeting(c *GreetingContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitGreeting is called when exiting the greeting production.
	ExitGreeting(c *GreetingContext)
}
