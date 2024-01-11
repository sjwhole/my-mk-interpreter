# Monkey Language Interpreter

A Go implementation of the Monkey programming language interpreter, following the structure from "Writing An Interpreter In Go" by Thorsten Ball.

## Table of Contents
- [Features](#features)
- [Components Overview](#components-overview)
- [Getting Started](#getting-started)
- [Usage Examples](#usage-examples)
- [Project Structure](#project-structure)
- [References](#references)

## Features

- Lexical analysis with tokenization
- Pratt parser for expression parsing
- Abstract Syntax Tree (AST) construction
- Tree-walking evaluation model
- Variable bindings and environment management
- Support for functions, conditionals, and return statements
- Built-in REPL (Read-Eval-Print Loop)

## Components Overview

### 1. Lexer
- **Algorithm**: Finite-state automaton
- Converts source code input into tokens
- Handles whitespace skipping and token position tracking
- Supports identifiers, literals, and operators

### 2. Parser
- **Algorithm**: Pratt parsing (Top-down operator precedence)
- Recursive descent implementation
- Produces an Abstract Syntax Tree (AST)
- Handles operator precedence and associativity

### 3. AST
- Defines Node interfaces for program structure
- Implements expression and statement nodes
- Supports modification through a `Modify()` method

### 4. Evaluator
- **Algorithm**: Tree-walking interpreter
- Executes the AST through recursive evaluation
- Manages environment and variable bindings
- Handles error propagation and return values

### 5. Object System
- Defines runtime objects (integers, booleans, strings, etc.)
- Manages environment chaining for scopes
- Implements function objects and error handling

## Getting Started

### Prerequisites
- Go 1.20+

### Installation
```bash
git clone https://github.com/sjwhole/my-mk-interpreter.git
cd my-mk-interpreter
```

## Usage Examples
### Build monkey executable
```bash
go build 
```

### Starting the REPL
```bash
./monkey
```

### Running a Monkey Program
Also, you can also run a Monkey program from a file. (like python)
```bash
./monkey <path-to-file>
```

### Example Monkey Program
```
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x-1) + fibonacci(x-2)
        }
    }
};
puts(fibonacci(10));
```

```bash
./monkey src/fibo.mk 
# expected output: 55
```

## Project Structure
```
monkey/
├── ast/          # Abstract Syntax Tree definitions
├── evaluator/    # Execution engine
├── lexer/        # Tokenization
├── object/       # Runtime object system
├── parser/       # Syntax parsing
├── repl/         # Read-Eval-Print Loop
└── token/        # Token definitions
```

## References
- Thorsten Ball's "Writing An Interpreter In Go"
- Pratt Parsing technique
- Monkey language specification