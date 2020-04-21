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
  int d = (test_case->chopstick_lengths[0] - test_case->chopstick_lengths[1]);
  return d * d;
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
      result->chopstick_lengths[0] = 4;
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
    printf("Done.\n");
    return;
}

int main(int argc, const char * argv[]) {
    test_all();
    return 0;
}
