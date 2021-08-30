// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`["3","6","7","10"]`, `4`, 
			`"3"`,
		},
		{
			`["2","21","12","1"]`, `3`, 
			`"2"`,
		},
		{
			`["0","0"]`, `2`, 
			`"0"`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, kthLargestNumber, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-256/problems/find-the-kth-largest-integer-in-the-array/