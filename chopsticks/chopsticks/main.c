//
//  main.c
//  chopsticks
//
//  Created by Terry Yin on 21/4/20.
//  Copyright © 2020 Terry Yin. All rights reserved.
//

#include <stdio.h>
#include <stdarg.h>
#include <assert.h>

typedef struct TestCase {
  int people_count;
  int chopstick_count;
  int chopstick_lengths[5000];
} TestCase;

int find_pair(int count, const int * lengths) {
  int result = 0;
  for(int i = 0; i <= count; i++) {
    int d = (lengths[i] - lengths[i + 1]);
    d = d * d;
    if (i == 0 || result > d) result = d;
  }
  return result;
}

int solve_one_case(TestCase * test_case) {
  int result = 0;
  for(int p = 0; p < test_case->people_count; p++) {
    int count = test_case->chopstick_count - p * 2;
    count -= (test_case-> people_count - p) * 3;
    result += find_pair(count, &test_case->chopstick_lengths[p * 2]);
  }
  return result;
}

int solver(int total, TestCase * test_cases) {
  return solve_one_case(test_cases);;
}

TestCase * example(TestCase * result, int people_count, int chopstick_count, ...) {
  va_list argp;
  TestCase tc = {people_count, chopstick_count};
  *result = tc;

  va_start(argp, chopstick_count);
  for(int i = 0; i < chopstick_count; i++)
    result->chopstick_lengths[i] = va_arg(argp, int);
  va_end(argp);

  return result;
}

void expect_eq(int expect, int actual, const char * message) {
  if (expect != actual) {
    fprintf(stderr, "expect %d, but got %d. %s\n", expect, actual, message);
  }
}

void test_all() {
  TestCase tc;
  expect_eq(1, solver(1, example(&tc, 1, 3, /**/ 4, 5, 5)), "one shorter");
  expect_eq(0, solver(1, example(&tc, 1, 3, /**/ 5, 5, 5)), "all same");
  expect_eq(0, solver(1, example(&tc, 1, 4, /**/ 4, 5, 5, 5)), "one shorter but there are more");
  expect_eq(0, solver(1, example(&tc, 1, 4, /**/ 4, 4, 5, 5)), "two short two long");

  expect_eq(1, solver(1, example(&tc, 2, 6, /**/ 4, 4, 4, 5, 6, 6)), "two people and 2n has different length");
  expect_eq(0, solver(1, example(&tc, 2, 6, /**/ 4, 4, 4, 4, 5, 5)), "two people and 2n has also same length");
  expect_eq(1, solver(1, example(&tc, 2, 6, /**/ 1, 2, 4, 4, 4, 4)), "two people only 2n has same length");

  expect_eq(23, solver(1, example(&tc, 9, 40, /**/ 1, 8, 10, 16, 19, 22, 27, 33, 36, 40, 47, 52, 56, 61, 63, 71, 72, 75, 81, 81, 84, 88, 96, 98,
          103, 110, 113, 118, 124, 128, 129, 134, 134, 139, 148, 157, 157, 160, 162, 164)), "final test");

  printf("Done.\n");
  return;
}

int main(int argc, const char * argv[]) {
  test_all();
  return 0;
}
