#!/bin/bash

cd src
go build
mkdir -p ../binaries
mv funlang ../binaries
../binaries/funlang ../examples/example.fl
