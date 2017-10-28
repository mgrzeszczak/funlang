#!/bin/bash

go build
mkdir -p binaries
mv funlang binaries
./binaries/funlang example.fl
