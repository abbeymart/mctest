// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect Test Package

package mctest

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestInstance type provides the instance object for the mcTest
type TestInstance struct {
	CaseName       string
	UnitTestPassed uint
	UnitTestFailed uint
	TestPassed     uint
	TestFailed     uint
	TestObject     *testing.T
}

// AssertEquals function asserts equality of a computation and expected result
func (testInstance *TestInstance) AssertEquals(expr interface{}, result interface{}, message string) string {
	t := testInstance.TestObject
	if expr == result {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		testInstance.TestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, result, expr)
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	testInstance.TestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, result, expr)
}

// AssertNotEquals function asserts inequality of a computation and expected result
func (testInstance *TestInstance) AssertNotEquals(expr interface{}, result interface{}, message string) string {
	t := testInstance.TestObject
	if expr != result {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		testInstance.TestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, result, expr)
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	testInstance.TestFailed += 1
	return fmt.Sprintf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, result, expr)
}

// AssertStrictEquals function asserts strict equality => deep equality check through stringified values
func (testInstance *TestInstance) AssertStrictEquals(expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)
	t := testInstance.TestObject
	if string(jsonExpr) == string(jsonResult) {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		testInstance.TestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	testInstance.TestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
}

// AssertNotStrictEquals function asserts strict inequality => deep equality check through stringified values
func (testInstance *TestInstance) AssertNotStrictEquals(expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)
	t := testInstance.TestObject
	if string(jsonExpr) != string(jsonResult) {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		testInstance.TestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	testInstance.TestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.CaseName, message, string(jsonResult), string(jsonExpr))
}

func (testInstance *TestInstance) McTest(params ParamsType) {
	testName := params.Name
	testFunc := params.TestFunc

	// make current testName accessible from textFunc
	testInstance.CaseName = testName

	// validate case/testName and testFunc
	if testName == "" || testFunc == nil {
		fmt.Printf("\n Test case name and test task/function are required - Testing stopped!!!")
		fmt.Printf("\n")
		return
	}
	// run test case
	fmt.Println("Running Test: ", testName)
	fmt.Println("================================================")
	testFunc()
	// Test report
	fmt.Println("Summary for Test ", testName, ":")
	fmt.Printf("\nTest Passed: %v", testInstance.UnitTestPassed)
	fmt.Printf("\nTest Failed: %v", testInstance.UnitTestFailed)
	fmt.Printf("\nTotal Test: %v\n", testInstance.UnitTestPassed+testInstance.UnitTestFailed)
	// Reset unit test counts
	testInstance.UnitTestPassed = 0
	testInstance.UnitTestFailed = 0
}

func (testInstance *TestInstance) PostTestResult() {
	fmt.Println("============================")
	fmt.Println("All Tests Summary Stats:")
	fmt.Println("============================")
	fmt.Printf("\nTest Passed: %v", testInstance.TestPassed)
	fmt.Printf("\nTest Failed: %v", testInstance.TestFailed)
	fmt.Printf("\nTotal Test: %v\n", testInstance.TestPassed+testInstance.TestFailed)
	// reset test counts
	testInstance.TestPassed = 0
	testInstance.TestFailed = 0
	fmt.Printf("\n***** Test Completed *****\n")
}
