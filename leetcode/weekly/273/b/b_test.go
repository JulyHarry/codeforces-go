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
			`3`, `[0,1]`, `"RRDDLU"`, 
			`[1,5,4,3,1,0]`,
		},
		{
			`2`, `[1,1]`, `"LURD"`, 
			`[4,1,0,0]`,
		},
		{
			`1`, `[0,0]`, `"LRUD"`, 
			`[0,0,0,0]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, executeInstructions, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-273/problems/execution-of-all-suffix-instructions-staying-in-a-grid/