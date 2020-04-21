//
//  main.c
//  chopsticks
//
//  Created by Terry Yin on 21/4/20.
//  Copyright © 2020 Terry Yin. All rights reserved.
//

#include <stdio.h>
#include <assert.h>

typedef struct TestCase {
  int people_count;
  int chopstick_count;
  int chopstick_lengths[5000];
} TestCase;

int find_pair(int count, const int * lengths) {
  int result = 0;
  for(int i = 0; i <= count - 3; i++) {
    int d = (lengths[i] - lengths[i + 1]);
    d = d * d;
    if (i == 0 || result > d) result = d;
  }
  return result;
}

int solve_one_case(TestCase * test_case) {
  int result = 0;
  for(int p = 0; p < test_case->people_count; p++) {
    result += find_pair(test_case->chopstick_count - p * 3, &test_case->chopstick_lengths[p * 3]);
  }
  return result;
}

int solver(int total, TestCase * test_cases) {
  return solve_one_case(test_cases);;
}

TestCase * for_people(int p, TestCase * result) {
  TestCase tc = {p,0,{}};
  *result = tc;
  return result;
}

TestCase * push(TestCase * result, int n, int length) {
  int i = result->chopstick_count;;
  for(; i < result-> chopstick_count + n; i++)
    result->chopstick_lengths[i] = length;
  result->chopstick_count = i;
  return result;
}

void expect_eq(int expect, int actual, const char * message) {
  if (expect != actual) {
    fprintf(stderr, "expect %d, but got %d. %s\n", expect, actual, message);
  }
}

void test_all() {
  TestCase tc;
  expect_eq(0, solver(1, push(for_people(1, &tc), 3, 5)), "all same");
  expect_eq(1, solver(1, push(push(for_people(1, &tc), 1, 4), 2, 5)), "one shorter");
  expect_eq(0, solver(1, push(push(for_people(1, &tc), 1, 4), 3, 5)), "one shorter but there are more");
  expect_eq(0, solver(1, push(push(for_people(1, &tc), 2, 4), 2, 5)), "two short two long");

  expect_eq(1, solver(1, push(push(push(for_people(2, &tc), 3, 4), 1, 5), 2, 6)), "two people and 2n has different length");
  expect_eq(0, solver(1, push(push(for_people(2, &tc), 4, 4), 2, 5)), "two people and 2n has also same length");

  printf("Done.\n");
  return;
}

int main(int argc, const char * argv[]) {
  test_all();
  return 0;
}
