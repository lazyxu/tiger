# tiger compiler
## 1. 项目概览
### 1.1 项目目标
用lex写出一个tiger语言的词法分析器，用YACC的分析方法完成对某一个的语法分析，并生成语法树和中间代码

### 1.2 运行
```
go generate
./tiger testcases/test4.tig
```
### 1.3 单元测试
#### lex 
```
cd yacc
go test -v -run='Test_Lex' -args ../testcases/test4.tig
```
#### yacc 
```
cd yacc
go test -v -run='Test_Yacc' -args ../testcases/test4.tig
```
### 1.3 代码结构

* yacc
    * lex.go
    * tiger.y
* symbol
    * Symbol 符号
    * Table 符号表
* absyn
    * pos.go
    * Exp 抽象语法树的结点
* types
    * Entry 变量入口
* frame
* translate
    * Exp 通过方法 Ex, Cx 或 Nx 来生成具体的结点
    * Level 层
* tree 
    * Exp
    * Stm ir树的结点
* frag
    * ProcFrag
    * StringFrag