// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	exampleIns := [][]string{{`[12,345,2,6,7896]`}, {`[555,901,482,1771]`}}
	exampleOuts := [][]string{{`2`}, {`1`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, findNumbers, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}