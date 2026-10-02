# 🚀 GoLisp: A LISP Interpreter in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/yashbaddi/golisp.svg)](https://pkg.go.dev/github.com/yashbaddi/golisp)
[![License: ISC](https://img.shields.io/badge/License-ISC-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yashbaddi/golisp)](https://goreportcard.com/report/github.com/yashbaddi/golisp)

A lightweight, robust, and tree-walking **LISP interpreter** implemented from scratch in **Go**. **GoLisp** parses and evaluates S-expressions with support for lexical scoping, first-class functions (closures), recursion, rich arithmetic, list manipulations, string literals with unicode escape sequences, and an interactive Read-Eval-Print Loop (REPL).

---

## 📑 Table of Contents

- [✨ Features](#-features)
- [🏗 Architecture & Design](#-architecture--design)
- [📦 Installation & Getting Started](#-installation--getting-started)
  - [Prerequisites](#prerequisites)
  - [Clone & Build](#clone--build)
  - [Running the Interactive REPL](#running-the-interactive-repl)
  - [Running Tests](#running-tests)
- [📖 Language Specification & Built-ins](#-language-specification--built-ins)
  - [Data Types](#data-types)
  - [Special Forms](#special-forms)
  - [Standard Library & Primitives](#standard-library--primitives)
- [💡 Code Examples](#-code-examples)
  - [1. Arithmetic & Constants](#1-arithmetic--constants)
  - [2. Variables & Lexical Scoping](#2-variables--lexical-scoping)
  - [3. Higher-Order Functions & Closures](#3-higher-order-functions--closures)
  - [4. Recursive Algorithms](#4-recursive-algorithms)
  - [5. List Processing & Functional Mapping](#5-list-processing--functional-mapping)
- [📁 Project Structure](#-project-structure)
- [⚖️ License](#️-license)

---

## ✨ Features

- **Tree-Walking Interpreter**: Direct AST/S-expression evaluation with high fidelity and simple semantics.
- **Lexical Scoping & Closures**: First-class `lambda` expressions retain access to their defining environment with hierarchical scope chains.
- **Special Forms**: Native implementation of `define`, `lambda`, and short-circuiting `if`.
- **First-Class Functions & Higher-Order Functions**: Pass functions as arguments, return functions, and map over lists.
- **Variadic Arithmetic**: Multi-operand arithmetic (`+`, `-`, `*`, `/`) with automatic `int` and `float64` type promotion.
- **Rich Primitive Library**: Includes modulo (`mod`), increment/decrement (`incf`/`decf`), power (`pow`), absolute value (`abs`), comparisons, and `pi`.
- **List Operations**: Classic LISP primitives: `list`, `cons`, `car`, `cdr`, and higher-order `map`.
- **Robust Lexer & Parser**: Unicode runes, string escape parsing (`\n`, `\t`, `\"`, `\\`, `\uXXXX`), identifiers, numbers, and boolean tokens (`#t`, `#f`).
- **Interactive REPL**: Built-in CLI REPL for live experimentation.
- **Zero External Dependencies**: Pure Go standard library implementation.

---

## 🏗 Architecture & Design

GoLisp follows a clean modular compiler/interpreter pipeline:

```
+-------------------------------------------------------------+
|                         Input Source                        |
|                     "(+ 10 (* 2 5))"                        |
+-------------------------------------------------------------+
                               |
                               v
+-------------------------------------------------------------+
|                 Lexer (internal/lexer)                      |
|  Tokenizes input into LPAREN, RPAREN, NUMBER, IDENTIFIER... |
+-------------------------------------------------------------+
                               |
                               v
+-------------------------------------------------------------+
|                Parser (internal/parser)                     |
|  Transforms token stream into nested S-expressions          |
|  (node.List, node.Symbol, int, float64, string, bool)       |
+-------------------------------------------------------------+
                               |
                               v
+-------------------------------------------------------------+
|              Evaluator (internal/evaluate)                  |
|  Executes S-Expressions against lexically-scoped Env        |
|  Handles Special Forms (if, define, lambda) & Primitives   |
+-------------------------------------------------------------+
                               |
                               v
+-------------------------------------------------------------+
|                         Output                              |
|                           20                                |
+-------------------------------------------------------------+
```

### Module Breakdown

| Package | Responsibility |
| :--- | :--- |
| `cmd/repl` | Entry point executable that connects `os.Stdin` and `os.Stdout` to the REPL. |
| `internal/token` | Defines token types (`LPAREN`, `RPAREN`, `STRING`, `NUMBER`, `BOOLEAN`, `IDENTIFIER`, `EOF`, etc.). |
| `internal/lexer` | Lexical scanner parsing runes, numbers, identifiers, escape characters, and unicode codepoints. |
| `internal/node` | Core S-expression data types: `node.Symbol` (`string`) and `node.List` (`[]any`). |
| `internal/ast` | AST nodes and string formatting representations (`ast.Program`, `ast.List`, `ast.Identifier`, `ast.NumberLiteral`). |
| `internal/parser` | Recursive-descent parser producing evaluated node structures from tokens. |
| `internal/evaluate` | Execution engine, environment chaining (`NewEnv`, `NewGlobalEnv`), special forms, and built-ins. |
| `internal/repl` | Interactive Read-Eval-Print Loop interface. |

---

## 📦 Installation & Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) **1.22+** (tested with Go 1.26+)

### Clone & Build

```bash
# Clone the repository
git clone https://github.com/yashbaddi/golisp.git
cd golisp

# Build the REPL executable
go build -o lisp ./cmd/repl
```

### Running the Interactive REPL

Start the REPL directly:

```bash
go run ./cmd/repl/main.go
```

Or run the compiled binary:

```bash
./lisp
```

Inside the REPL, type LISP expressions and inspect results in real time:

```lisp
lisp > (+ 10 20 30)
60
lisp > (define square (lambda (x) (* x x)))
<func([]interface {}) (interface {}, error)>
lisp > (square 9)
81
```

### Running Tests

Execute the comprehensive test suite across all packages:

```bash
go test -v ./...
```

Run tests with race detection and coverage:

```bash
go test -race -cover ./...
```

---

## 📖 Language Specification & Built-ins

### Data Types

| Type | LISP Syntax | Go Representation | Example |
| :--- | :--- | :--- | :--- |
| **Integer** | `123`, `-42` | `int` | `42` |
| **Float** | `3.14`, `-0.05` | `float64` | `3.14159` |
| **Boolean** | `#t` (true), `#f` (false) | `bool` | `#t` |
| **String** | `"..."` (supports `\n`, `\t`, `\uXXXX`) | `string` | `"Hello, World!"` |
| **Symbol** | `foo`, `make-adder`, `+`, `equal?` | `node.Symbol` | `x` |
| **List** | `(item1 item2 ...)` | `node.List` (`[]any`) | `(list 1 2 3)` |

### Special Forms

Special forms do not evaluate all their arguments upfront; they follow distinct execution rules:

| Form | Syntax | Description |
| :--- | :--- | :--- |
| `define` | `(define <symbol> <expr>)` | Binds `<symbol>` to the result of `<expr>` in the current lexical environment. |
| `if` | `(if <test> <then> [else])` | Evaluates `<test>`. If truthy (not `#f`, not `nil`, and not empty list `()`), evaluates `<then>`; otherwise evaluates `<else>` (or returns `nil`). Short-circuits execution. |
| `lambda` | `(lambda (<params...>) <body...> )` | Creates a first-class anonymous closure capturing the current lexical scope. Evaluates body expressions sequentially and returns the result of the last expression. |

### Standard Library & Primitives

#### ➕ Arithmetic & Math
- `(+ ...)`: Variadic sum (`(+ 1 2 3 4)` $\rightarrow$ `10`).
- `(- first ...)`: Negation or variadic subtraction (`(- 10 3 2)` $\rightarrow$ `5`, `(- 5)` $\rightarrow$ `-5`).
- `(* ...)`: Variadic product (`(* 2 3 4)` $\rightarrow$ `24`).
- `(/ first ...)`: Unary reciprocal or variadic division (`(/ 10 2)` $\rightarrow$ `5`, `(/ 2.0)` $\rightarrow$ `0.5`).
- `(mod a b)`: Integer modulo (`(mod 10 3)` $\rightarrow$ `1`).
- `(incf x [delta])`: Increment by `1` or optional delta (`(incf 5 3)` $\rightarrow$ `8`).
- `(decf x [delta])`: Decrement by `1` or optional delta (`(decf 10 4)` $\rightarrow$ `6`).
- `(abs n)`: Absolute value of an integer or float (`(abs -42)` $\rightarrow$ `42`).
- `(pow base exp)`: Exponentiation (`(pow 2 10)` $\rightarrow$ `1024`).
- `pi`: Built-in mathematical constant $\pi \approx 3.141592653589793$.

#### ⚖️ Comparisons & Predicates
- `(> a b ...)`: Strict decreasing order test.
- `(>= a b ...)`: Non-increasing order test.
- `(< a b ...)`: Strict increasing order test.
- `(<= a b ...)`: Non-decreasing order test.
- `(= a b ...)`: Numeric equality comparison.
- `(equal? a b)`: Deep structural equality (compares lists, symbols, strings, and primitives).

#### 📜 List Operations
- `(list ...)`: Constructs a new list from evaluated arguments (`(list 1 2 3)` $\rightarrow$ `(1 2 3)`).
- `(cons elem list)`: Prepends `elem` to the front of `list` (`(cons 0 (list 1 2))` $\rightarrow$ `(0 1 2)`).
- `(car list)`: Returns the first element of a non-empty list (`(car (list 10 20))` $\rightarrow$ `10`).
- `(cdr list)`: Returns the tail list containing all elements after `car` (`(cdr (list 10 20 30))` $\rightarrow$ `(20 30)`).
- `(map fn list)`: Applies unary function `fn` to every element of `list`, returning a new list.

---

## 💡 Code Examples

### 1. Arithmetic & Constants

```lisp
;; Variadic arithmetic
(+ 1 2 3 4 5)          ; => 15
(* 2 3 4 5)            ; => 120
(- 100 20 10)          ; => 70
(/ 100 2 2)            ; => 25

;; Math operations
(mod 17 5)             ; => 2
(pow 2 16)             ; => 65536
(abs -3.14)            ; => 3.14
(incf 10 5)            ; => 15
(decf 10)              ; => 9

;; Math constants
(* 2 pi 10)            ; => 62.83185307179586 (Circumference)
```

### 2. Variables & Lexical Scoping

```lisp
;; Global definitions
(define radius 5)
(define area (* pi (pow radius 2)))

;; Lexical scope isolation
(define x 100)
(define get-scoped (lambda ()
  (define x 200)
  x))

(get-scoped)           ; => 200
x                      ; => 100 (outer environment is untouched)
```

### 3. Higher-Order Functions & Closures

```lisp
;; Function returning a function (Closure / Currying)
(define make-adder (lambda (x)
  (lambda (y) (+ x y))))

(define add10 (make-adder 10))
(add10 25)             ; => 35

;; Function accepting a function
(define repeat (lambda (f)
  (lambda (x) (f (f x)))))

(define twice (lambda (x) (* 2 x)))
((repeat twice) 10)    ; => 40
```

### 4. Recursive Algorithms

```lisp
;; Factorial
(define fact (lambda (n)
  (if (<= n 1)
      1
      (* n (fact (- n 1))))))

(fact 5)               ; => 120
(fact 10)              ; => 3628800

;; Fibonacci
(define fib (lambda (n)
  (if (< n 2)
      1
      (+ (fib (- n 1)) (fib (- n 2))))))

(fib 10)               ; => 89
```

### 5. List Processing & Functional Mapping

```lisp
;; List manipulation
(define nums (list 1 2 3 4 5))
(car nums)             ; => 1
(cdr nums)             ; => (2 3 4 5)
(cons 0 nums)          ; => (0 1 2 3 4 5)

;; Mapping functions over lists
(define square (lambda (x) (* x x)))
(map square (list 1 2 3 4 5)) ; => (1 4 9 16 25)

;; Recursive range generator and Fibonacci mapping
(define range (lambda (a b)
  (if (= a b)
      (list)
      (cons a (range (+ a 1) b)))))

(map fib (range 0 10)) ; => (1 1 2 3 5 8 13 21 34 55)
```

---

## 📁 Project Structure

```
.
├── LICENSE                   # ISC License
├── README.md                 # Project documentation
├── cmd/
│   └── repl/
│       └── main.go           # CLI REPL entry point
├── go.mod                    # Go module definition (github.com/yashbaddi/golisp)
├── internal/
│   ├── ast/                  # AST node definitions & string formatters
│   │   ├── ast.go
│   │   └── ast_test.go
│   ├── evaluate/             # Evaluator, environments, builtins, special forms
│   │   ├── builtins_test.go
│   │   ├── env.go
│   │   ├── evaluate.go
│   │   ├── evaluate_test.go
│   │   ├── specialForms.go
│   │   └── specialForms_test.go
│   ├── lexer/                # Lexer scanning tokens & unicode handling
│   │   ├── errors.go
│   │   ├── lexer.go
│   │   ├── lexer_test.go
│   │   └── readToken.go
│   ├── node/                 # S-Expression node types (Symbol, List)
│   │   └── node.go
│   ├── parser/               # Recursive descent parser
│   │   ├── errors.go
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── repl/                 # Interactive REPL implementation
│   │   └── repl.go
│   └── token/                # Token types and definitions
│       ├── token.go
│       └── token_test.go
└── tests/                    # End-to-end integration test suite
    └── lisp_test.go
```

---

## ⚖️ License

This project is licensed under the **ISC License**. See the [LICENSE](LICENSE) file for details.

---

*Developed with ❤️ by [Yash Baddi](https://github.com/yashbaddi)*
