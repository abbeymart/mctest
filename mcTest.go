// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect Test Package

package mctest

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestFunction ***** types *****
type TestFunction func()

// ParamsType make params public
type ParamsType struct {
	Name     string
	TestFunc TestFunction
	Before   string
	After    string
}

// ***** variables *****
var (
	caseName       = ""
	unitTestPassed = 0
	unitTestFailed = 0
	testPassed     = 0
	testFailed     = 0
)

// AssertEquals function asserts equality of a computation and expected result
func AssertEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	if expr == result {
		fmt.Println("Passed")
		unitTestPassed += 1
		testPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
	fmt.Printf("\n")
	unitTestFailed += 1
	testFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
}

// AssertNotEquals function asserts inequality of a computation and expected result
func AssertNotEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	if expr != result {
		fmt.Println("Passed")
		unitTestPassed += 1
		testPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
	fmt.Printf("\n")
	unitTestFailed += 1
	testFailed += 1
	return fmt.Sprintf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
}

// AssertStrictEquals function asserts strict equality => deep equality check through stringified values
func AssertStrictEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)

	if string(jsonExpr) == string(jsonResult) {
		fmt.Println("Passed")
		unitTestPassed += 1
		testPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	unitTestFailed += 1
	testFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
}

// AssertNotStrictEquals function asserts strict inequality => deep equality check through stringified values
func AssertNotStrictEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)

	if string(jsonExpr) != string(jsonResult) {
		fmt.Println("Passed")
		unitTestPassed += 1
		testPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	unitTestFailed += 1
	testFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
}

func McTest(params ParamsType) {
	testName := params.Name
	testFunc := params.TestFunc

	// make current testName accessible from textFunc
	caseName = testName

	// validate caseName and testFunc
	if caseName == "" || testFunc == nil {
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
	fmt.Println("Test Passed: ", unitTestPassed)
	fmt.Println("Test Failed: ", unitTestFailed)
	fmt.Println("Total Test: ", unitTestPassed+unitTestFailed)
	// Reset unit test counts
	unitTestPassed = 0
	unitTestFailed = 0
}

func PostTestResult() {
	fmt.Println("============================")
	fmt.Println("All Tests Summary Stats:")
	fmt.Println("============================")
	fmt.Println("Test Passed: ", testPassed)
	fmt.Println("Test Failed: ", testFailed)
	fmt.Println("Total Test: ", testPassed+testFailed)
	// reset test counts
	testPassed = 0
	testFailed = 0
	fmt.Println("***** Test Completed *****")
}
