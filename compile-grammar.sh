#!/bin/bash

antlr4 -visitor -no-listener -package generated Funlang.g4 -Dlanguage=Go
mkdir -p generated
mv funlang_*.go generated
rm *.tokens
