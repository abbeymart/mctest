// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mctest

import "testing"

// test-data
var (
	result1 = 100
	result2 = 200
	result3 = map[string]string{"ItemName": "Abi"}
	result4 = map[string]string{"location": "Abi"}
)

func Expr1() int               { return 100 }
func Expr2() int               { return 200 }
func Expr3() map[string]string { return result3 }
func Expr4() map[string]string { return result4 }
func TestMcCases(t *testing.T) {
	// empty testcase name: testing will return with printed testcase name and testFunc required
	McTest(ParamsType{
		Name: "",
		TestFunc: func() {
			AssertEquals(t, Expr1(), result1, "Expected outcome: 100")
			AssertEquals(t, Expr2(), result2, "Expected outcome: 200")
		},
	})
	// Nil testFunc: testing will return with printed testcase name and testFunc required
	McTest(ParamsType{
		Name:     "Nil test-func",
		TestFunc: nil,
	})
	McTest(ParamsType{
		Name: "Test Series 200",
		TestFunc: func() {
			AssertEquals(t, Expr1(), result1, "Expected outcome: 100")
			AssertEquals(t, Expr2(), result2, "Expected outcome: 200")
			AssertNotEquals(t, Expr1(), result2, "Expected expr and result not equals")
			AssertNotEquals(t, Expr2(), result1, "Expected expr and result not equals")
			AssertNotStrictEquals(t, Expr3(), result4, "Expected outcome: not strictly equals")
			AssertNotStrictEquals(t, Expr4(), result3, "Expected outcome: not strictly equals")
		},
	})

	PostTestResult()
}
