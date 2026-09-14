
````
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
````
🐒 Monkey Lang — A Go Virtual Machine/Compiler

This project is an implementation of the Monkey programming language, built by following the book "Write a Compiler in Go"

---

📘 Overview

Monkey is a small, interpreted, dynamically-typed language with C-like syntax and JavaScript-like semantics.
It supports variables, functions, closures, arrays, hashes, and control flow structures like if and return.

This Go project implements:

Lexer — tokenizes the Monkey source code

Parser — builds an abstract syntax tree (AST)

Evaluator — executes the AST nodes

REPL — interactive Read-Eval-Print Loop for experimentation

---
🧩 Project Structure
````
Monkey-Lang/
├── main.go          # Entry point, starts the REPL
├── repl/            # REPL (interactive shell)
├── lexer/           # Lexical analyzer
├── parser/          # Parser to build AST
├── ast/             # AST node definitions
├── object/          # Runtime objects and environment
├── evaluator/       # Evaluation logic
├── token/           # Token definitions
└── go.mod           # Go module file
````
🚀 Getting Started
1. Clone the Repository

````bash
   git clone https://github.com/Sina-karimi81/Monkey-lang-compiler.git
   cd Monkey-lang-compiler
````
2. Run the REPL
```bash
go run main.go
```
---

💻 Example Session
````monkey
>> let a = 5 * 5;
>> a
25
>> let add = fn(x, y) { x + y };
>> add(10, 5)
15
>> if (10 > 5) { return true; } else { return false; }
true
````
---
⚙️ Requirements
- Go 1.20+
- No external dependencies
---
🧠 Learning Resource

This interpreter is built by following:

📖 [Write an Interpreter in Go](https://interpreterbook.com/)

---
🛠️ Future Work

- [x] Add built-in functions to type
- [x] Add support for arrays
- [x] Add support for hashes/objects
- [ ] Add support for interpreting files
- [ ] complete the standard library
- [ ] improve error handling