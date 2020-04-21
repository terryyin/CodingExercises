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

int solve_one_case(TestCase * test_case) {
  int top = test_case->chopstick_count - test_case->people_count * 3;
  int result = 0;
  for(int i = 0; i <= top; i++) {
    int d = (test_case->chopstick_lengths[i] - test_case->chopstick_lengths[i + 1]);
    d = d * d;
    if (i == 0 || result > d) result = d;
  }
  return result;
}

int solver(int total, TestCase * test_cases) {
    return solve_one_case(test_cases);;
}

TestCase * people_chopsticks(int p, int n, TestCase * result) {
    TestCase tc = {p,n,{5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5}};
    *result = tc;
    return result;
}

TestCase * shorter(int n, TestCase * result) {
    for(int i=0; i < n; i++)
      result->chopstick_lengths[i] = 4;
    return result;
}

void expect_eq(int expect, int actual, const char * message) {
  if (expect != actual) {
    fprintf(stderr, "expect %d, but got %d. %s\n", expect, actual, message);
  }
}

void test_all() {
    TestCase tc;
    expect_eq(0, solver(1, people_chopsticks(1, 3, &tc)), "all same");
    expect_eq(1, solver(1, shorter(1, people_chopsticks(1, 3, &tc))), "one shorter");
    expect_eq(0, solver(1, shorter(1, people_chopsticks(1, 4, &tc))), "one shorter but there are more");
    expect_eq(0, solver(1, shorter(2, people_chopsticks(1, 4, &tc))), "two short two long");
    printf("Done.\n");
    return;
}

int main(int argc, const char * argv[]) {
    test_all();
    return 0;
}
