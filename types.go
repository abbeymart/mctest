package mctest

import "testing"

// TestFunction ***** types *****
type TestFunction func()

// ParamsType make params public
type ParamsType struct {
	Name           string
	TestFunc       TestFunction
	UnitTestPassed uint
	UnitTestFailed uint
	UnitTestTotal  uint
	Before         string
	After          string
	TestObject     *testing.T
}

type UnitTestResult struct {
	UnitTestPassed uint
	UnitTestFailed uint
	UnitTestTotal  uint
}
