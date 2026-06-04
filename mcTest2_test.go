// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30, 2026-06-03
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mctest

import (
	"testing"
)

func TestMcTest(t *testing.T) {
	// capture all test results
	var results []UnitTestResult

	// test-data
	var (
		result1a = 100
		result2a = 200
		result3a = map[string]string{"ItemName": "Abi"}
		result4a = map[string]string{"location": "Abi"}
	)

	Expr1a := func() int { return 100 }
	Expr2a := func() int { return 200 }
	Expr3a := func() map[string]string { return result3a }
	Expr4a := func() map[string]string { return result4a }

	// empty testcase name: testing will return with printed testcase name and testFunc required
	test0 := NewTest(ParamsType{})
	test0.SetTestFunction(func() {
		test0.AssertEquals(Expr1a(), result1a, "Expected outcome: 100")
		test0.AssertEquals(Expr2a(), result2a, "Expected outcome: 200")
	})
	// run test and update results
	test0Result := test0.RunTest()
	results = append(results, test0Result)

	// instance object
	test1 := NewTest(ParamsType{
		Name: "TestSeries100a",
	})
	// set test function
	test1.SetTestFunction(func() {
		test1.AssertEquals(Expr1a(), result1a, "Expected outcome: 100")
		test1.AssertEquals(Expr2a(), result2a, "Expected outcome: 200")
	})
	// run test and update results
	test1Result := test1.RunTest()
	results = append(results, test1Result)

	test2 := NewTest(ParamsType{
		Name: "Test Series 200a",
	})
	// set test function
	test2.SetTestFunction(func() {
		test2.AssertEquals(Expr1a(), result1a, "Expected outcome: 100")
		test2.AssertEquals(Expr2a(), result2a, "Expected outcome: 200")
		test2.AssertNotEquals(Expr1a(), result2a, "Expected expr and result not equals")
		test2.AssertNotEquals(Expr2a(), result1a, "Expected expr and result not equals")
		test2.AssertNotStrictEquals(Expr3a(), result4a, "Expected outcome: not strictly equals")
		test2.AssertNotStrictEquals(Expr4a(), result3a, "Expected outcome: not strictly equals")
	})
	// run test and update results
	test2Result := test2.RunTest()
	results = append(results, test2Result)

	TestResult(results)
}
