// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/600/B
// https://codeforces.com/problemset/status/600/problem/B
func Test_cf600B(t *testing.T) {
	testCases := [][2]string{
		{
			`5 4
1 3 5 7 9
6 4 2 8`,
			`3 2 1 4`,
		},
		{
			`5 5
1 2 1 2 5
3 1 4 1 5`,
			`4 2 4 2 5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf600B)
}
