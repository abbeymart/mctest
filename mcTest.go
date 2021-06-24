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

// OptionValue make params public
type OptionValue struct {
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
	passedTest     = 0
	failedTest     = 0
)

// AssertEquals ***** functions *****
// assert equals
func AssertEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	if expr == result {
		fmt.Println("Passed")
		unitTestPassed += 1
		passedTest += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
	fmt.Printf("\n")
	unitTestFailed += 1
	failedTest += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, result, expr)
}

// AssertNotEquals assert not equals
func AssertNotEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	if expr != result {
		fmt.Println("Passed")
		unitTestPassed += 1
		passedTest += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
	fmt.Printf("\n")
	unitTestFailed += 1
	failedTest += 1
	return fmt.Sprintf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, result, expr)
}

// AssertStrictEquals assert not strict equals => deep equality check through stringified values
func AssertStrictEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)

	if string(jsonExpr) == string(jsonResult) {
		fmt.Println("Passed")
		unitTestPassed += 1
		passedTest += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	unitTestFailed += 1
	failedTest += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", caseName, message, string(jsonResult), string(jsonExpr))
}

// AssertNotStrictEquals assert strict equals => deep equality check through stringified values
func AssertNotStrictEquals(t *testing.T, expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)

	if string(jsonExpr) != string(jsonResult) {
		fmt.Println("Passed")
		unitTestPassed += 1
		passedTest += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	unitTestFailed += 1
	failedTest += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v and %v not to be equals", caseName, message, string(jsonResult), string(jsonExpr))
}

func McTest(options OptionValue) {
	var (
		testName string
		testFunc TestFunction
	)
	if options.Name != "" {
		testName = options.Name
	} else {
		testName = "Unknown"
	}
	// make current testName accessible from textFunc
	caseName = testName

	if options.TestFunc != nil {
		testFunc = options.TestFunc
	} else {
		testFunc = nil
	}
	fmt.Println("Running Test: ", testName)
	fmt.Println("================================================")
	if testFunc != nil {
		testFunc()
	} else {
		fmt.Printf("\n No test task/function specified - Test skipped!!!")
		fmt.Printf("\n")
	}
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
	fmt.Println("Test Passed: ", passedTest)
	fmt.Println("Test Failed: ", failedTest)
	fmt.Println("Total Test: ", passedTest+failedTest)
	// reset test counts
	passedTest = 0
	failedTest = 0
	fmt.Println("***** Test Completed *****")
}
