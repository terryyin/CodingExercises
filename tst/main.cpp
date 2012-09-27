#include "CppUTest/CommandLineTestRunner.h"

int main(int ac, char** av) {

	int res = CommandLineTestRunner::RunAllTests(ac, av);
	return res;
}
