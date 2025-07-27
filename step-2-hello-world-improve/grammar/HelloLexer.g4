lexer grammar HelloLexer;

HELLO: 'hello';
SEMICOLON: ';';

// 词法规则
ID: [a-zA-Z_][a-zA-Z0-9_]*; // 标识符
WS: [ \t\r\n]+ -> skip; // 空白字符