// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mctest

import (
	"testing"
)

func TestMcTest(t *testing.T) {
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

	// instance object
	tInstance := TestInstance{
		TestObject: t,
	}
	// empty testcase name: testing will return with printed testcase name and testFunc required
	tInstance.McTest(ParamsType{
		Name: "Test Series 100a",
		TestFunc: func() {
			tInstance.AssertEquals(Expr1a(), result1a, "Expected outcome: 100")
			tInstance.AssertEquals(Expr2a(), result2a, "Expected outcome: 200")
		},
	})
	tInstance.McTest(ParamsType{
		Name: "Test Series 200a",
		TestFunc: func() {
			tInstance.AssertEquals(Expr1a(), result1a, "Expected outcome: 100")
			tInstance.AssertEquals(Expr2a(), result2a, "Expected outcome: 200")
			tInstance.AssertNotEquals(Expr1a(), result2a, "Expected expr and result not equals")
			tInstance.AssertNotEquals(Expr2a(), result1a, "Expected expr and result not equals")
			tInstance.AssertNotStrictEquals(Expr3a(), result4a, "Expected outcome: not strictly equals")
			tInstance.AssertNotStrictEquals(Expr4a(), result3a, "Expected outcome: not strictly equals")
		},
	})

	tInstance.PostTestResult()
}
