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
### 1.3 代码完成情况
* [✔︎] absyn 
    * exp
    * dec
    * pos
    * ty
    * var
* [✔︎] env 
    * tenv venv
    * entry
* [✔︎] frame
    * access
    * frame 
    * frag
        * ProcFrag
        * StringFrag
* semant
    * SEM_transProg()
* [✔︎] symbol
    * Symbol 符号
* [✔︎] table
    * Table 表
* [✔︎] temp
    * temp
    * label
* [✔︎] yacc
    * lex.go
    * tiger.y
* [✔︎] types
    * ty 类型
* [✔︎] translate
    * Exp 通过方法 Ex, Cx 或 Nx 来生成具体的结点
    * Level 层
    * Access
* [✔︎] tree 
    * exp
    * op 
    * stm 
* [✔︎] util