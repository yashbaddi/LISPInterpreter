import lispint from "./lisp.js";

// lispint("()");
console.log(lispint("(+ 1 1)"));
// console.log(lispint("(dw 3 3 )"));
console.log(lispint("(define circle-area (lambda (r) (* pi (* r r))))"));
console.log(lispint("(circle-area 3)"));
console.log(
  lispint("(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))")
);
console.log(lispint("(fact 10)"));
console.log(lispint("(fact 100)"));
console.log(lispint("(circle-area (fact 10))"));
console.log(lispint("(define first car)"));
console.log(lispint("(define rest cdr)"));
console.log(
  lispint(
    "(define count (lambda (item L) (if L (+ (equal? item (first L)) (count item (rest L))) 0)))"
  )
);
console.log(lispint("(count 0 (list 0 1 2 3 0 0))"));
console.log(
  lispint(
    "(count (quote the) (quote (the more the merrier the bigger the better)))"
  )
);
lispint("(define twice (lambda (x) (* 2 x)))");
console.log(lispint("(twice 5)"));
lispint("(define repeat (lambda (f) (lambda (x) (f (f x)))))");
console.log(lispint("((repeat twice) 10)"));
console.log(lispint("((repeat (repeat twice)) 10)"));
console.log(lispint("((repeat (repeat (repeat twice))) 10)"));
console.log(lispint("((repeat (repeat (repeat (repeat twice)))) 10)"));
console.log(lispint("(pow 2 16)"));
console.log(
  lispint(
    "(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))"
  )
);
console.log(
  lispint(
    "(define range (lambda (a b) (if (= a b) 1 (cons a (range (+ a 1) b)))))"
  )
);

console.log(lispint("(range 0 10)"));
console.log(lispint("(map fib (range 0 10))"));
console.log(lispint("(map fib (range 0 20))"));
