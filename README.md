# 🚀 LISP Interpreter in JavaScript

A lightweight and efficient LISP interpreter implemented in Node.js, supporting core functional programming concepts such as recursion, high-order functions, and lexically scoped environments.

---

## ✨ Features

- **Core Special Forms**: `define`, `lambda`, `if`, `quote`.
- **Arithmetic Operators**: `+`, `-`, `*`, `/`, `mod`, `incf`, `decf`, `abs`, `pow`.
- **Comparison Operators**: `>`, `>=`, `<`, `<=`, `=`, `equal?`.
- **List Operations**: `list`, `cons`, `car`, `cdr`, `map`.
- **Recursion**: Full support for recursive function definitions (e.g., Factorial, Fibonacci).
- **Lexical Scoping**: Functions (lambdas) create their own local environments.
- **Interactive REPL**: Built-in Read-Eval-Print Loop for real-time experimentation.

---

## 🛠️ Getting Started

### Prerequisites

- **Node.js** (v14 or higher recommended)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yashbaddi/LISPInterpreter.git
   cd LISPInterpreter
   ```

2. (Optional) Install dependencies (none currently required beyond standard Node modules):
   ```bash
   npm install
   ```

### Running the Interpreter

To start the interactive REPL:

```bash
node lisp.js
```

Once in the REPL, you can type LISP expressions and see the results immediately:

```lisp
LISP > (+ 5 10)
15
LISP > (define square (lambda (x) (* x x)))
LISP > (square 4)
16
```

---

## 📚 Examples & Usage

### ⚙️ Basic Arithmetic

```lisp
(+ 1 2 3 4)      ; returns 10
(* 2 5 10)       ; returns 100
(pow 2 10)       ; returns 1024
```

### 🧠 Functional Programming (Lambdas & Higher-Order)

```lisp
(define twice (lambda (x) (* 2 x)))
(twice 5)        ; returns 10

(define repeat (lambda (f) (lambda (x) (f (f x)))))
((repeat twice) 10) ; returns 40
```

### 🔄 Recursion

```lisp
;; Factorial Implementation
(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))
(fact 5) ; returns 120

;; Fibonacci Implementation
(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))
(fib 10) ; returns 89
```

### 📜 List Manipulation

```lisp
(define my-list (list 1 2 3 4))
(car my-list)           ; returns 1
(cdr my-list)           ; returns (2 3 4)
(map (lambda (x) (* x x)) my-list) ; returns (1 4 9 16)
```

---

## 📂 Project Structure

- `lisp.js`: Main entry point and REPL implementation.
- `expressionParser.js`: Responsible for parsing complex LISP expressions and nested brackets.
- `valueParser.js`: Handles parsing of individual values (numbers, strings, booleans).
- `specialForms.js`: Implements core LISP constructs like `define`, `if`, and `lambda`.
- `symbols.js`: Defines the global environment and built-in primitive functions.
- `test.js`: Contains a suite of test cases showcasing various functionalities.

---

## ⚖️ License

This project is licensed under the ISC License.

---

_Developed with ❤️ by Yash Baddi_
