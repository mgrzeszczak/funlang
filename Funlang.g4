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
    : name parameters Semicolon
    ;

parameterList
    : number
    | number Comma parameterList
    ;

parameters
    : OpenParam CloseParam
    | OpenParam parameterList CloseParam
    ;

definition
    : name AssignmentOperator function Semicolon
    ;

name
    : Nondigit (Nondigit | Digit)*
    ;

function
    : composition
    | function CompositionOperator innerFunctions
    ;

innerFunctions
    : OpenParam functionList CloseParam
    | function
    ;

functionList
    : function
    | function Comma functionList
    ;

composition
    : composition PrimitiveRecursionOperator primitiveRecursion
    | primitiveRecursion
    ;

primitiveRecursion
    : OpenParam function CloseParam
    | zero
    | successor
    | projection
    | name
    ;

zero
    : 'Z' OpenParam number CloseParam
    ;

successor
    : 'S'
    ;

projection
    : 'P' OpenParam number ',' number CloseParam
    ;

number
    : Digit+
    ;

AssignmentOperator
    : '='
    ;

CompositionOperator
    : '|'
    ;

PrimitiveRecursionOperator
    : '&'
    ;

Comma
    : ','
    ;

Semicolon
    : ';'
    ;

OpenParam
    : '('
    ;
CloseParam
    : ')'
    ;

Nondigit
    : [a-zA-Z_]
    ;

Digit
    : [0-9]
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