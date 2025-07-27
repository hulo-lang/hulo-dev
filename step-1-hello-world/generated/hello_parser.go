// Code generated from Hello.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // Hello
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type HelloParser struct {
	*antlr.BaseParser
}

var HelloParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func helloParserInit() {
	staticData := &HelloParserStaticData
	staticData.LiteralNames = []string{
		"", "'hello'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"program", "greeting",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 17, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 5, 0, 6, 8, 0, 10, 0, 12, 0,
		9, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0,
		15, 0, 7, 1, 0, 0, 0, 2, 12, 1, 0, 0, 0, 4, 6, 3, 2, 1, 0, 5, 4, 1, 0,
		0, 0, 6, 9, 1, 0, 0, 0, 7, 5, 1, 0, 0, 0, 7, 8, 1, 0, 0, 0, 8, 10, 1, 0,
		0, 0, 9, 7, 1, 0, 0, 0, 10, 11, 5, 0, 0, 1, 11, 1, 1, 0, 0, 0, 12, 13,
		5, 1, 0, 0, 13, 14, 5, 3, 0, 0, 14, 15, 5, 2, 0, 0, 15, 3, 1, 0, 0, 0,
		1, 7,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// HelloParserInit initializes any static state used to implement HelloParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHelloParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HelloParserInit() {
	staticData := &HelloParserStaticData
	staticData.once.Do(helloParserInit)
}

// NewHelloParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHelloParser(input antlr.TokenStream) *HelloParser {
	HelloParserInit()
	this := new(HelloParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HelloParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Hello.g4"

	return this
}

// HelloParser tokens.
const (
	HelloParserEOF  = antlr.TokenEOF
	HelloParserT__0 = 1
	HelloParserT__1 = 2
	HelloParserID   = 3
	HelloParserWS   = 4
)

// HelloParser rules.
const (
	HelloParserRULE_program  = 0
	HelloParserRULE_greeting = 1
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllGreeting() []IGreetingContext
	Greeting(i int) IGreetingContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HelloParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HelloParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HelloParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(HelloParserEOF, 0)
}

func (s *ProgramContext) AllGreeting() []IGreetingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGreetingContext); ok {
			len++
		}
	}

	tst := make([]IGreetingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGreetingContext); ok {
			tst[i] = t.(IGreetingContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Greeting(i int) IGreetingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreetingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreetingContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *HelloParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HelloParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HelloParserT__0 {
		{
			p.SetState(4)
			p.Greeting()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(10)
		p.Match(HelloParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGreetingContext is an interface to support dynamic dispatch.
type IGreetingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsGreetingContext differentiates from other interfaces.
	IsGreetingContext()
}

type GreetingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreetingContext() *GreetingContext {
	var p = new(GreetingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HelloParserRULE_greeting
	return p
}

func InitEmptyGreetingContext(p *GreetingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HelloParserRULE_greeting
}

func (*GreetingContext) IsGreetingContext() {}

func NewGreetingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreetingContext {
	var p = new(GreetingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HelloParserRULE_greeting

	return p
}

func (s *GreetingContext) GetParser() antlr.Parser { return s.parser }

func (s *GreetingContext) ID() antlr.TerminalNode {
	return s.GetToken(HelloParserID, 0)
}

func (s *GreetingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreetingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreetingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.EnterGreeting(s)
	}
}

func (s *GreetingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.ExitGreeting(s)
	}
}

func (p *HelloParser) Greeting() (localctx IGreetingContext) {
	localctx = NewGreetingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HelloParserRULE_greeting)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Match(HelloParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(13)
		p.Match(HelloParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(14)
		p.Match(HelloParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
