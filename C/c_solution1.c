#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <errno.h>

#if __WORDSIZE == 64
# define ULONG_MAX_DIGIT_LENGTH 20
#else
# define ULONG_MAX_DIGIT_LENGTH 10
#endif
#define BOOK_COST_BUFSIZE (ULONG_MAX_DIGIT_LENGTH + sizeof(char) + 1)

typedef struct {
    char currency;
    unsigned long value;
} Book_t;

char * calcBooksCost(Book_t * books[], size_t books_amount) {
    static char result_buf[BOOK_COST_BUFSIZE] = {0};
    unsigned long value_sum = 0;

    for (unsigned int i = 0; i < books_amount; i++) {
        unsigned long sum_before = value_sum;
        value_sum += books[i]->value;
        if (value_sum < sum_before) {
            fprintf(stderr, "Sum value overflow!\n");
            return NULL;
        }
    }
    snprintf(result_buf, BOOK_COST_BUFSIZE - 1, "%lu%c", value_sum,
            books[0]->currency);
    return result_buf;
}

Book_t * parseBook(const char * book_str) {
  Book_t * new_book = malloc(sizeof(Book_t));
  char cur = book_str[strlen(book_str) - 1];

  new_book->currency = isdigit(cur) ? '\0' : cur;
  char *end_ptr = NULL;
  new_book->value = strtoul(book_str, &end_ptr, 10);
  if (end_ptr == book_str || errno == ERANGE || errno == EINVAL) {
    fprintf(stderr, "Value parsing error!\n");
    return NULL;
  }
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
    cart[i-1] = NULL;
    cart[i-1] = parseBook(argv[i]);
    if (!cart[i-1]) {
      printf("fail on %d\n", i);
      for (int j = 0; j < i - 1; j++) {
        printf("free %d\n", j);
        free(cart[j]);
      }
      return 1;
    }
  }

  const char *sum = calcBooksCost(cart, books_amount);
  if (sum)
    printf("Total: %s\n", sum);

  for (unsigned int i = 0; i < books_amount; i++) {
    free(cart[i]);
  }
  if (!sum)
    return 1;
  return 0;
}
