parser grammar HelloParser;

options {
	tokenVocab = HelloLexer;
}

// 语法规则：程序由问候语句组成
program: greeting* EOF;

// 问候语句：hello + 名字
greeting: HELLO ID SEMICOLON;