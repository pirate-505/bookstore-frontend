#lang racket

(struct price 
  (value currency)) 

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
          (substring price-str (- pivot 1) len)
          (- pivot 1))]))
  (parse-try price-str "" len))


(define list-of-prices
    (map
      (lambda (arg)
        (parse-price arg))
      (vector->list (current-command-line-arguments))))

(define (result-price)
  (price 
    (foldl + 0 (map price-value list-of-prices))
    (price-currency (car list-of-prices))))

(define usage
  "usage: racket bookstore.rkt <cost1>[currency] [cost2 cost3 ...]\nExample: racket bookstore.rkt 9$ 9$\n")

(cond [(< (vector-length (current-command-line-arguments)) 2)
       (display usage)]
      [(> (length (remove-duplicates (map price-currency list-of-prices))) 1)
       (display (string-append "Currency must be the same\n"
                               usage))] 
      [else  
        (display (format "~a~a\n" 
                         (price-value (result-price))
                         (price-currency (result-price))))])
