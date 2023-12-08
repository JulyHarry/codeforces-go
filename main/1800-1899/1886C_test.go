// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1886/problem/C
// https://codeforces.com/problemset/status/1886/problem/C
func Test_cf1886C(t *testing.T) {
	testCases := [][2]string{
		{
			`3
cab
6
abcd
9
x
1`,
			`abx`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1886C)
}