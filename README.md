# Monkey Interpreter in Go 🐵

This repository contains my implementation of the Monkey programming language,
following along with the book [*Writing an Interpreter in
Go*](https://interpreterbook.com/) by Thorsten Ball.

The goal of this project is to deepen my understanding of how programming
languages work by building one from scratch — covering key components like
lexing, parsing, evaluation, and building a REPL — all written in Go.

## 📚 About the Book

> *Writing an Interpreter in Go* is a hands-on guide that walks you through the
> process of creating an interpreter for a programming language, starting from
> tokenizing source code all the way to executing it in a simple REPL.

## 🛠️ Features Implemented

* [x] Lexer: Converts source code into tokens
* [] Parser: Builds an abstract syntax tree (AST)
* [] Evaluator: Executes the AST with proper scoping
* [] REPL: Interactive shell for running Monkey code
* [x] Support for:

  * Integer arithmetic
  * Boolean logic
  * Conditionals (`if`, `else`)
  * Functions & closures
  * Let bindings
  * First-class functions

## 🚀 Getting Started

```bash 
git clone https://github.com/yourusername/monkey-interpreter-go.git 
cd monkey-interpreter-go go run main.go 
```

Then start typing Monkey code in the REPL!

```monkey
>> let add = fn(a, b) { a + b; }; add(2, 3);
5 
```

## 🧠 Purpose

This project is **for learning purposes only**. It’s a way to explore how
interpreters work under the hood, while sharpening my Go skills and
understanding of language design.

## 🗂️ Structure

* `lexer/` – Tokenization logic
* `parser/` – AST and parsing rules
* `evaluator/` – Evaluation engine
* `repl/` – Read-Eval-Print Loop
* `object/` – Object system (e.g. integers, booleans, functions)
* `main.go` – Entry point for the interpreter

## 🙌 Acknowledgements

Big thanks to [Thorsten Ball](https://twitter.com/thorstenball) for writing such
a fantastic book and making interpreters approachable and fun to build.

