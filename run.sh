#!/bin/bash

cd src
go build -o funlang
mv funlang ../binaries
../binaries/funlang ../examples/example.fl
