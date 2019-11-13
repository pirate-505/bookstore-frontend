#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#if __WORDSIZE == 64
# define ULONG_MAX_DIGIT_LENGTH 20
#else
# define ULONG_MAX_DIGIT_LENGTH 10
#endif

typedef struct {
    char currency;
    unsigned long value;
} Book_t;

char * calcBooksCost(Book_t * books[], size_t books_amount) {
    static char result_buf[ULONG_MAX_DIGIT_LENGTH + sizeof(char) + 1] = {0};
    unsigned long value_sum = 0;

    for (unsigned int i = 0; i < books_amount; i++) {
        value_sum += books[i]->value;
    }
    sprintf(result_buf, "%ld%c", value_sum, books[0]->currency);
    return result_buf;
}

Book_t * parseBook(const char * book_str) {
  Book_t * new_book = malloc(sizeof(Book_t));
  char cur = book_str[strlen(book_str) - 1];
  new_book->currency = (cur >= 48 && cur <= 57) ? '\0' : cur;
  new_book->value = strtol(book_str, NULL, 10);
  return new_book;
}

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    printf("usage: %s <cost1>[currency] [cost2 cost3 ...]\nExample: %s 9$ 9$\n", argv[0], argv[0]);
    return 0;
  }
  unsigned int books_amount = argc - 1;
  Book_t * cart[books_amount];

  for (int i = 1; i < argc; i++) {
    cart[i-1] = parseBook(argv[i]);
  }

  const char *sum = calcBooksCost(cart, books_amount);

  printf("Total: %s\n", sum);

  for (unsigned int i = 0; i < books_amount; i++) {
    free(cart[i]);
  }
  return 0;
}