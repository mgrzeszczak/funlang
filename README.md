# funlang

Programming language that allows to define primitive recursive functions and calculate them.

# Operators
* __|__ - functional composition
* __&__ - primitive recursion

# Basic functions
* __Z__(x) - zero function, where x - number of arguments
* __S__ - successor function
* __P__(size,which) - projection functions, returns 'which' argument from 'size' arguments (counting from 0)

# Usage
1. Build funlang using `go build` or use precompiled binary
2. Execute scripts using `./funlang script.fl`

# Example program
```
// definitions
def add =  P(1,0) & (S | P(3,2));
add(100,200);

def double = add |(P(1,0),P(1,0));
double(2);

def multiply = Z(1) & (add | ( P(3,2), P(3,0) ) );
def power = (S | Z(1)) & (multiply | (P(3,2), P(3,0)));
def predecessor = Z(0) & P(2,0);
def sign = Z(0) & ( S | Z(2));
def nsign = (S | Z(0)) & (Z(2));
def ld = P(1,0) & (predecessor | P(3,2));
def absdiff = add | (ld | (P(2,0),P(2,1)), ld | (P(2,1),P(2,0)));

// executions
multiply(2,5);
power(2,10);
predecessor(2);
sign(0);
sign(1);
sign(100);
nsign(0);
nsign(1);
ld(10,5);
ld(9,10);
absdiff(10,5);
absdiff(5,10);
```

Running this script gives the following result:
```
$ ./run.sh
add[100 200] = 300
double[2] = 4
multiply[2 5] = 10
power[2 10] = 1024
predecessor[2] = 1
sign[0] = 0
sign[1] = 1
sign[100] = 1
nsign[0] = 1
nsign[1] = 0
ld[10 5] = 5
ld[9 10] = 0
absdiff[10 5] = 5
absdiff[5 10] = 5
```

# License
```
MIT License

Copyright (c) 2017 Maciej Grzeszczak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
