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

// Test Instance

// NewTest function is a factory function to create a new test object/instance
func NewTest(params ParamsType) *ParamsType {
	return &ParamsType{
		UnitTestPassed: 0,
		UnitTestFailed: 0,
		UnitTestTotal:  0,
		Name:           params.Name,
		TestFunc:       params.TestFunc,
		Before:         params.Before,
		After:          params.After,
	}
}

// SetTestFunction sets the test-instance function
func (testInstance *ParamsType) SetTestFunction(params TestFunction) {
	testInstance.TestFunc = params
}

// AssertEquals function asserts equality of a computation and expected result
func (testInstance *ParamsType) AssertEquals(expr interface{}, result interface{}, message string) string {
	t := testInstance.TestObject
	if expr == result {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, result, expr)
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, result, expr)
}

// AssertNotEquals function asserts inequality of a computation and expected result
func (testInstance *ParamsType) AssertNotEquals(expr interface{}, result interface{}, message string) string {
	t := testInstance.TestObject
	if expr != result {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, result, expr)
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, result, expr)
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	return fmt.Sprintf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, result, expr)
}

// AssertStrictEquals function asserts strict equality => deep equality check through stringified values
func (testInstance *ParamsType) AssertStrictEquals(expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)
	t := testInstance.TestObject
	if string(jsonExpr) == string(jsonResult) {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v, Got %v", testInstance.Name, message, string(jsonResult), string(jsonExpr))
}

// AssertNotStrictEquals function asserts strict inequality => deep equality check through stringified values
func (testInstance *ParamsType) AssertNotStrictEquals(expr interface{}, result interface{}, message string) string {
	// stringify expr and result for strict equals comparison
	jsonExpr, _ := json.Marshal(expr)
	jsonResult, _ := json.Marshal(result)
	t := testInstance.TestObject
	if string(jsonExpr) != string(jsonResult) {
		fmt.Println("Passed")
		testInstance.UnitTestPassed += 1
		return "Passed"
	}
	fmt.Printf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, string(jsonResult), string(jsonExpr))
	t.Errorf("\nFailed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, string(jsonResult), string(jsonExpr))
	fmt.Printf("\n")
	testInstance.UnitTestFailed += 1
	return fmt.Sprintf("Failed [Test-Case: %v]: %v => Expected %v and %v not to be equals", testInstance.Name, message, string(jsonResult), string(jsonExpr))
}

func (testInstance *ParamsType) UnitTestResult() UnitTestResult {
	return UnitTestResult{
		UnitTestPassed: testInstance.UnitTestPassed,
		UnitTestFailed: testInstance.UnitTestFailed,
		UnitTestTotal:  testInstance.UnitTestPassed + testInstance.UnitTestFailed,
	}
}

func (testInstance *ParamsType) RunTest() UnitTestResult {
	// validate test-case name and testFunc
	if testInstance.Name == "" || testInstance.TestFunc == nil {
		fmt.Printf("\n Test case name and test task/function are required - Testing stopped!!!")
		fmt.Printf("\n")
		return UnitTestResult{}
	}
	// run test case
	fmt.Println("Running Test: ", testInstance.Name)
	fmt.Println("================================================")
	testInstance.TestFunc()
	testInstance.UnitTestTotal = testInstance.UnitTestPassed + testInstance.UnitTestFailed
	// Test report
	fmt.Println("Summary for Test ", testInstance.Name, ":")
	fmt.Printf("\nTest Passed: %v", testInstance.UnitTestPassed)
	fmt.Printf("\nTest Failed: %v", testInstance.UnitTestFailed)
	fmt.Printf("\nTotal Test: %v\n", testInstance.UnitTestPassed+testInstance.UnitTestFailed)
	return testInstance.UnitTestResult()
}

// TestResult function captures overall testing results for all the unit tests
func TestResult(params []UnitTestResult) {
	testPassed := uint(0)
	testFailed := uint(0)
	testTotal := uint(0)
	for _, val := range params {
		testPassed += val.UnitTestPassed
		testFailed += val.UnitTestFailed
		testTotal += val.UnitTestTotal
	}
	fmt.Println("============================")
	fmt.Println("All Tests Summary Stats:")
	fmt.Println("============================")
	fmt.Printf("\nTest Passed: %v", testPassed)
	fmt.Printf("\nTest Failed: %v", testFailed)
	fmt.Printf("\nTotal Test: %v\n", testTotal)
	fmt.Printf("\n***** Test Completed *****\n")
}
