%{
package yacc
import(
    "github.com/MeteorKL/tiger/absyn"
    "github.com/MeteorKL/tiger/symbol"
)
const(
    EOF = 0
    Debug = 0
    ErrorVerbose = true
)
var Absyn_root absyn.Exp;
%}

%union {
	ty         absyn.Ty
	namety     absyn.Namety
	sym        symbol.Symbol
	Var        absyn.Var
	exp        absyn.Exp
	expList    absyn.ExpList
	dec        absyn.Dec
	decList    absyn.DecList
	field      absyn.Field
	fieldList  absyn.FieldList
	fundec     absyn.Fundec
	fundecList absyn.FundecList
	nametyList absyn.NametyList
	efield     absyn.Efield
	efieldList absyn.EfieldList

	ival int
	sval string
}

%type <ty> ty
%type <namety> tydec
%type <sym> id
%type <Var> lvalue
%type <exp> root exp let cond
%type <expList> arglist arg explst
%type <dec> dec vardec tydecs fundecs
%type <decList> decs
%type <field> tyfield
%type <fieldList> tyfields ty_field
%type <fundec> fundec
%type <efieldList> recordlist record

%token <ival> INT
%token <sval> ID STRING

%token 
  /* punctuation mark */
  COMMA COLON SEMICOLON LPAREN RPAREN LBRACK RBRACK LBRACE RBRACE DOT PLUS MINUS TIMES DIVIDE EQ NEQ LT LE GT GE AND OR ASSIGN
  /* reserved word */
  reserved_word_beg
  WHILE FOR TO BREAK LET IN END FUNCTION VAR TYPE ARRAY IF THEN ELSE DO OF NIL
  reserved_word_end
  ILLEGAL


%nonassoc LOW
%nonassoc THEN DO TYPE FUNCTION ID 
%nonassoc ASSIGN LBRACK ELSE OF COMMA
%left OR
%left AND
%nonassoc EQ NEQ LE LT GT GE
%left PLUS MINUS
%left TIMES DIVIDE
%left UMINUS

%start program
%%
program:    root                          {Absyn_root=$1}
            ;

root:       /* empty */                   {$$=nil}
            | exp                         {$$=$1}
            ;

exp:        INT                           {$$=&absyn.IntExp{EM_tokPos, $1}}
            | STRING                      {$$=&absyn.StringExp{EM_tokPos, $1}}
            | NIL                         {$$=&absyn.NilExp{EM_tokPos}}
            | id LPAREN arglist RPAREN    {$$=&absyn.CallExp{EM_tokPos, $1, $3}}
            | lvalue                      {$$=&absyn.VarExp{EM_tokPos, $1}}
            | LPAREN explst RPAREN        {$$=&absyn.SeqExp{EM_tokPos, $2}}
            | cond                        {$$=$1}
            | let                         {$$=$1}
            | exp OR exp                  {$$=&absyn.IfExp{EM_tokPos, $1, &absyn.IntExp{EM_tokPos,1}, $3}}
            | exp AND exp                 {$$=&absyn.IfExp{EM_tokPos, $1, $3, &absyn.IntExp{EM_tokPos,0}}}
            | exp LT exp                  {$$=&absyn.OpExp{EM_tokPos, absyn.LtOp, $1, $3}}
            | exp GT exp                  {$$=&absyn.OpExp{EM_tokPos, absyn.GtOp, $1, $3}}
            | exp LE exp                  {$$=&absyn.OpExp{EM_tokPos, absyn.LeOp, $1, $3}}
            | exp GE exp                  {$$=&absyn.OpExp{EM_tokPos, absyn.GeOp, $1, $3}}
            | exp PLUS exp                {$$=&absyn.OpExp{EM_tokPos, absyn.PlusOp, $1, $3}}
            | exp MINUS exp               {$$=&absyn.OpExp{EM_tokPos, absyn.MinusOp, $1, $3}}
            | exp TIMES exp               {$$=&absyn.OpExp{EM_tokPos, absyn.TimesOp, $1, $3}}
            | exp DIVIDE exp              {$$=&absyn.OpExp{EM_tokPos, absyn.DivideOp, $1, $3}}
            | exp EQ exp                  {$$=&absyn.OpExp{EM_tokPos, absyn.EqOp, $1, $3}}
            | exp NEQ exp                 {$$=&absyn.OpExp{EM_tokPos, absyn.NeqOp, $1, $3}}
            | id LBRACK exp RBRACK OF exp {$$=&absyn.ArrayExp{EM_tokPos, $1, $3, $6}}
            | id LBRACE recordlist RBRACE {$$=&absyn.RecordExp{EM_tokPos, $1, $3}}
            | lvalue ASSIGN exp           {$$=&absyn.AssignExp{EM_tokPos, $1, $3}}
            | MINUS exp %prec UMINUS      {$$=&absyn.OpExp{EM_tokPos, absyn.MinusOp, &absyn.IntExp{EM_tokPos, 0}, $2}}
            | BREAK                       {$$=&absyn.BreakExp{EM_tokPos}}
            ;

recordlist: /* empty */                   {$$=nil}
            | record                      {$$=$1}
            ;

record:     id EQ exp                     {$$=absyn.EfieldListInsert(&absyn.Efield_{$1, $3}, nil)}
            | id EQ exp COMMA record      {$$=absyn.EfieldListInsert(&absyn.Efield_{$1, $3}, $5)}
            ;

let:        LET decs IN explst END        {$$=&absyn.LetExp{EM_tokPos, $2, &absyn.SeqExp{EM_tokPos, $4}}}
            ;

arglist:    /* empty */                   {$$=nil}
            | arg                         {$$=$1}
            ;

arg:        exp                           {$$=absyn.ExpListInsert($1, nil)}
            | exp COMMA arg               {$$=absyn.ExpListInsert($1, $3)}
            ;

            /* decs -> {dec} */
decs:       /* empty */                   {$$=nil}
            | dec decs                    {$$=absyn.DecListInsert($1, $2)}
            ;

            /* dec -> tydec */
            /*     -> vardec */
            /*     -> fundec */
dec:        tydecs                        {$$=$1}
            | vardec                      {$$=$1}
            | fundecs                     {$$=$1}
            ;
            
tydecs:     tydec %prec LOW               {$$=&absyn.TypeDec{EM_tokPos, absyn.NametyListInsert($1, nil)}}
            | tydec tydecs                {$$=&absyn.TypeDec{EM_tokPos, absyn.NametyListInsert($1, $2.(*absyn.TypeDec).Type)}}
            ;

            /* tydec -> type type-id = ty */
tydec:      TYPE id EQ ty                 {$$=&absyn.Namety_{$2, $4}}
            ;

            /* ty -> type-id */
            /*    -> { tyfields } */
            /*    -> array of type-id */
ty:         id                            {$$=&absyn.NameTy{EM_tokPos, $1}}
            | LBRACE tyfields RBRACE      {$$=&absyn.RecordTy{EM_tokPos, $2}}
            | ARRAY OF id                 {$$=&absyn.ArrayTy{EM_tokPos, $3}}
            ;

            /* tyfields -> empty */
            /*          -> id:type-id{,id:type-id} */
tyfields:   /* empty */                   {$$=nil}
            | ty_field                    {$$=$1}
            ;

ty_field:   tyfield                       {$$=absyn.FieldListInsert($1, nil)}
            | tyfield COMMA ty_field      {$$=absyn.FieldListInsert($1, $3)}
            ;

tyfield:    id COLON id                   {$$=&absyn.Field_{EM_tokPos, $1, $3, true}}
            ;

            /* vardec -> var id:=exp*/
            /*        -> var id:type-id:=exp */
vardec:     VAR id ASSIGN exp             {$$=&absyn.VarDec{EM_tokPos, $2, nil, $4, true}}
            | VAR id COLON id ASSIGN exp  {$$=&absyn.VarDec{EM_tokPos, $2, $4, $6, true}}
            ;

fundecs:    fundec %prec LOW              {$$=&absyn.FunctionDec{EM_tokPos, absyn.FundecListInsert($1, nil)}}
            | fundec fundecs              {$$=&absyn.FunctionDec{EM_tokPos, absyn.FundecListInsert($1, $2.(*absyn.FunctionDec).Function)}}
            ;

            /* fundec -> function id (tyfields) = exp */
            /*        -> function id (tyfields) : type-id = exp */
fundec:     FUNCTION id LPAREN tyfields RPAREN EQ exp             {$$=&absyn.Fundec_{EM_tokPos, $2, $4, nil, $7}}
            | FUNCTION id LPAREN tyfields RPAREN COLON id EQ exp  {$$=&absyn.Fundec_{EM_tokPos, $2, $4, $7, $9}}
            ;

id:         ID                            {$$=symbol.SymbolInsert($1)}
            ;
            
            /* lvalue -> id */
            /*        -> lvalue.id */
            /*        -> lvalue[exp] */
lvalue:     id %prec LOW                  {$$=&absyn.SimpleVar{EM_tokPos, $1}}
            | lvalue DOT id               {$$=&absyn.FieldVar{EM_tokPos, $1, $3}}
            | id LBRACK exp RBRACK        {$$=&absyn.SubscriptVar{EM_tokPos, &absyn.SimpleVar{EM_tokPos, $1}, $3}}
            | lvalue LBRACK exp RBRACK    {$$=&absyn.SubscriptVar{EM_tokPos, $1, $3}}
            ;

explst:     /* empty */                   {$$=nil}      
            | exp                         {$$=absyn.ExpListInsert($1, nil)}
            | exp SEMICOLON explst        {$$=absyn.ExpListInsert($1, $3)}
            ;

cond:       IF exp THEN exp ELSE exp                              {$$=&absyn.IfExp{EM_tokPos, $2, $4, $6}}
            | IF exp THEN exp                                     {$$=&absyn.IfExp{EM_tokPos, $2, $4, nil}}
            | WHILE exp DO exp                                    {$$=&absyn.WhileExp{EM_tokPos, $2, $4}}
            | FOR id ASSIGN exp TO exp DO exp                     {$$=&absyn.ForExp{EM_tokPos, $2, $4, $6, $8, true}}
            ;
