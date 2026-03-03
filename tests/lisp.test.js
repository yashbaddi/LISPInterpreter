import { describe, it, expect } from "vitest";
import lispint from "../lisp.js";

describe("Basic Arithmetic", () => {
  it("should add numbers correctly", () => {
    expect(lispint("(+ 1 1)")).toBe(2);
    expect(lispint("(+ 1 2 3 4)")).toBe(10);
  });

  it("should subtract numbers correctly", () => {
    expect(lispint("(- 10 3)")).toBe(7);
    expect(lispint("(- 10 2 1)")).toBe(7);
  });

  it("should multiply numbers correctly", () => {
    expect(lispint("(* 2 3 4)")).toBe(24);
  });

  it("should handle division (reversed user logic)", () => {
    expect(lispint("(/ 10 2)")).toBe(0.2); // (2 / 10)
    expect(lispint("(/ 100 10 2)")).toBe(20); // (10 / 100) = 0.1, then (2 / 0.1) = 20
  });

  it("should handle mod (reversed user logic)", () => {
    expect(lispint("(mod 10 3)")).toBe(3); // 3 % 10
    expect(lispint("(mod 3 10)")).toBe(1); // 10 % 3
  });

  it("should handle power/exponentiation", () => {
    expect(lispint("(pow 2 10)")).toBe(1024);
    expect(lispint("(pow 2 16)")).toBe(65536);
  });
});

describe("Variable Definitions and Scope", () => {
  it("should define variables", () => {
    lispint("(define x 10)");
    expect(lispint("(+ x 5)")).toBe(15);
  });

  it("should handle nested definitions and lambdas", () => {
    lispint("(define twice (lambda (x) (* 2 x)))");
    expect(lispint("(twice 5)")).toBe(10);
  });

  it("should handle higher-order functions", () => {
    lispint("(define repeat (lambda (f) (lambda (x) (f (f x)))))");
    lispint("(define twice (lambda (x) (* 2 x)))");
    expect(lispint("((repeat twice) 10)")).toBe(40);
    expect(lispint("((repeat (repeat twice)) 10)")).toBe(160);
    expect(lispint("((repeat (repeat (repeat twice))) 10)")).toBe(2560);
    expect(lispint("((repeat (repeat (repeat (repeat twice)))) 10)")).toBe(
      655360,
    );
  });

  it("should handle lexical scoping", () => {
    lispint("(define x 10)");
    lispint("(define make-adder (lambda (x) (lambda (y) (+ x y))))");
    lispint("(define add5 (make-adder 5))");
    expect(lispint("(add5 10)")).toBe(15);
    expect(lispint("x")).toBe(10); // original x should be unchanged
  });
});

describe("Conditionals and Comparisons", () => {
  it("should handle booleans", () => {
    expect(lispint("#t")).toBe(true);
    expect(lispint("#f")).toBe(false);
  });

  it("should handle if conditions", () => {
    expect(lispint("(if (> 10 5) 1 0)")).toBe(1);
    expect(lispint("(if (< 10 5) 1 0)")).toBe(0);
  });

  it("should handle nested if conditions", () => {
    expect(lispint("(if #t (if #f 3 2) 7)")).toBe(2);
  });

  it("should handle comparison operators", () => {
    expect(lispint("(>= 10 10)")).toBe(true);
    expect(lispint("(<= 5 10)")).toBe(true);
    expect(lispint("(= 42 42)")).toBe(true);
    expect(lispint("(equal? 100 100)")).toBe(true);
  });

  it("should handle recursive factorial", () => {
    lispint("(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))");
    expect(lispint("(fact 5)")).toBe(120);
    expect(lispint("(fact 10)")).toBe(3628800);
  });

  it("should handle recursive fibonacci", () => {
    lispint(
      "(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))",
    );
    expect(lispint("(fib 10)")).toBe(89);
  });
});

describe("List Operations", () => {
  it("should handle list creation and manipulation", () => {
    lispint("(define my-list (list 1 2 3 0 0))");
    expect(lispint("(car my-list)")).toBe(1);
    expect(lispint("(cdr my-list)")).toEqual([2, 3, 0, 0]);
  });

  it("should handle list processing (count example)", () => {
    lispint("(define first car)");
    lispint("(define rest cdr)");
    lispint(
      "(define count (lambda (item L) (if L (+ (equal? item (first L)) (count item (rest L))) 0)))",
    );
    expect(lispint("(count 0 (list 0 1 2 3 0 0))")).toBe(3);

    // Testing with quote
    expect(
      lispint(
        "(count (quote the) (quote (the more the merrier the bigger the better)))",
      ),
    ).toBe(4);
  });

  it("should handle map and range", () => {
    lispint(
      "(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))",
    );
    lispint(
      "(define range (lambda (a b) (if (= a b) (quote ()) (cons a (range (+ a 1) b)))))",
    );

    const fibs = lispint("(map fib (range 0 10))");
    expect(fibs).toEqual([1, 1, 2, 3, 5, 8, 13, 21, 34, 55]);
  });
});

describe("Strings and Quotes", () => {
  it("should parse strings", () => {
    expect(lispint('"Hello World"')).toBe("Hello World");
  });

  it("should handle quoted symbols", () => {
    expect(lispint("(quote sym)")).toBe("sym");
  });

  it("should handle quoted lists", () => {
    expect(lispint("(quote (a b c))")).toEqual(["a", "b", "c"]);
  });
});
