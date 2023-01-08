// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc272/submit?taskScreenName=abc272_e
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3 3
-1 -1 -6`,
			`2
2
0`,
		},
		{
			`5 6
-2 -2 -5 -7 -15`,
			`1
3
2
0
0
0`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc272/tasks/abc272_e