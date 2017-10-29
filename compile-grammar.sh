#!/bin/bash

antlr4 -visitor -no-listener -package generated Funlang.g4 -Dlanguage=Go
mkdir -p src/generated
mv funlang_*.go src/generated
rm *.tokens
