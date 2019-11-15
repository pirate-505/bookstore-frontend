#lang racket

(provide (all-defined-out))
(struct price 
  (value currency)
  #:transparent) 

;; parse-book : string -> book
;; turn string like 8$ into price with value and currency
;; return #f if string is invalid
(define (parse-price price-str)
  (define len (string-length price-str))
  (define (parse-try trial-value trial-currency pivot)
    (cond 
      [(= (string-length trial-value)
          0) 
       #f]
      [(string->number trial-value)
       (price (string->number trial-value) trial-currency)]
      [else 
        (parse-try 
          (substring price-str 0 (- pivot 1))
          (substring price-str (- pivot 1) pivot)
          (- pivot 1))]))
  (parse-try price-str "" len))


(define list-of-prices
    (map
      (lambda (arg)
        (parse-price arg))
      (vector->list (current-command-line-arguments))))

(define result-price
  (price 
    (foldl + 0 (map price-value list-of-prices))
    "$"))

(if (< (vector-length (current-command-line-arguments)) 2)
  (and
    (display "usage: racket bookstore.rkt <cost1>[currency] [cost2 cost3 ...]\nExample: racket bookstore.rkt 9$ 9$\n")
    (exit))
    (display (format "~a~a\n" 
                 (price-value result-price)
                 (price-currency result-price))))
