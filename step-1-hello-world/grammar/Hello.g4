grammar Hello;

// 语法规则：程序由问候语句组成
program: greeting* EOF;

// 问候语句：hello + 名字
greeting: 'hello' ID ';';

// 词法规则
ID: [a-zA-Z_][a-zA-Z0-9_]*;  // 标识符
WS: [ \t\r\n]+ -> skip;      // 空白字符