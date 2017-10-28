grammar Funlang;

program
    : definitions tasks
    ;

definitions
    : definition*
    ;

tasks
    : task*
    ;

task
    : name parameters ';'
    ;

args
    : number
    | number ',' args
    ;

parameters
    : OpenParam CloseParam
    | OpenParam args CloseParam
    ;

definition
    : name '=' function ';'
    ;

name
    : Nondigit (Nondigit | Digit)*
    ;

function
    : composition
    | function '|' OpenParam functions CloseParam
    ;

functions
    : function
    | function ',' functions
    ;

composition
    : composition '&' primitiveRecursion
    | primitiveRecursion
    ;

primitiveRecursion
    : OpenParam function CloseParam
    | zero
    | successor
    | projection
    | name
    ;

OpenParam
    : '('
    ;
CloseParam
    : ')'
    ;

zero
    : '$Z' OpenParam number CloseParam
    ;

successor
    : '$S'
    ;

projection
    : '$P' OpenParam number ',' number CloseParam
    ;

Nondigit
    : [a-zA-Z_]
    ;

Digit
    : [0-9]
    ;

number
    : Digit+
    ;

Whitespace
    :   [ \t]+
        -> skip
    ;

Newline
    :   (   '\r' '\n'?
        |   '\n'
        )
        -> skip
    ;

BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;

LineComment
    :   '//' ~[\r\n]*
        -> skip
    ;