package goca

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	expected := "testCase"
	secExpected := "TeStCase"
	var tests = testCases{
		{"test Case", expected, ""},
		{"TEsT CasE", expected, ""},
		{"TeST_cAse", expected, ""},
		{"TeST\\cAse", expected, ""},
		{"TeStCase", secExpected, ""},
	}

	runTest(tests, CamelCase, t)
}
func TestCapitalize(t *testing.T) {
	expected := "Test Case"
	secExpected := "Testcase"
	var tests = testCases{
		{"test Case", expected, ""},
		{"TEsT CasE", expected, ""},
		{"TeST_cAse", expected, ""},
		{"TeST\\cAse", expected, ""},
		{"TeStCase", secExpected, ""},
	}

	runTest(tests, Capitalize, t)
}
func TestDecapitalize(t *testing.T) {
	expected := "tEST cASE"
	secExpected := "tESTCASE"
	var tests = testCases{
		{"test Case", expected, ""},
		{"TEsT CasE", expected, ""},
		{"TeST_cAse", expected, ""},
		{"TeST\\cAse", expected, ""},
		{"TeStCase", secExpected, ""},
	}

	runTest(tests, Decapitalize, t)
}

func TestKebabCase(t *testing.T) {
	expected := "test-case"
	result := KebabCase(" tesT CaSe")

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestKebabCaseSeparator(t *testing.T) {
	var tests = testCases{
		{"test Case", "test-case", "-"},
		{"TEsT CasE", "test~case", "~"},
		{"TeST_cAse", "test!case", "!"},
		{"TeST\\cAse", "test&case", "&"},
	}
	runTestWithParams(tests, KebabCase, t)
}

func TestSwapCase(t *testing.T) {
	var test = testCases{
		{"tEsT cAsE", "TeSt CaSe", ""},
		{"tEsT_cAsE", "TeSt_CaSe", ""},
		{"tEsT 2 -cAsE", "TeSt 2 -CaSe", ""},
	}
	runTest(test, SwapCase, t)
}

func TestTitleCase(t *testing.T) {
	var tests = testCases{
		{"test Case", "Test Case", ""},
		{"TEsT CasE", "Test Case", "*"},
		{"TEsT*CasE", "Test*case", "*"},
		{"TeST_cAse", "Test_case", "_"},
		{"TeST\\cAse", "Test Case", "A"},
		{"TeStCase", "Testcase", ""},
	}

	runTestWithParams(tests, TitleCase, t)
}

func TestCharAt(t *testing.T) {

	var tests = intTestCases{
		{"test Case", "t", 0},
		{"TEsT CasE", "E", 1},
		{"TeST_cAse", "_", 4},
		{"TeST\\cAse", "\\", 4},
		{"TeST\\cAse", "TeST\\cAse", 100},
		{"Hello, 世界", "界", 8},
	}

	runTestWithIntParams(tests, CharAt, t)
}

func TestCodeAt(t *testing.T) {
	var tests = testCases{
		{"a", "61", ""},
		{"\\", "5c", ""},
		{"界", "e7958c", ""},
		{"Й", "d099", ""},
	}
	runTest(tests, HexAt, t)
}

func TestFirst(t *testing.T) {
	var tests = intTestCases{
		{"test Case", "", 0},
		{"TEsT CasE", "TEs", 3},
		{"TeST_cAse", "TeST_", 5},
		{"TeST\\cAse", "TeST\\", 5},
		{"TeST\\cAse", "TeST\\cAse", 100},
		{"Hello, 世界", "Hello, 世", 8},
	}
	runTestWithIntParams(tests, First, t)
}

func TestLast(t *testing.T) {
	var tests = intTestCases{
		{"test Case", "", 0},
		{"TEsT CasE", "asE", 3},
		{"TeST_cAse", "_cAse", 5},
		{"TeST\\cAse", "\\cAse", 5},
		{"TeST\\cAse", "TeST\\cAse", 100},
		{"Hello, 世界", "世界", 2},
	}
	runTestWithIntParams(tests, Last, t)
}

func TestPruneNoParam(t *testing.T) {
	result := Prune("Test case", 5)
	expected := "Test..."

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestPrune(t *testing.T) {
	type LocalTestCase struct {
		input     string
		strLength uint64
		end       string
		expected  string
	}

	var tests = []LocalTestCase{
		{"Test case", 9, "$.$", "Test case$.$"},
		{"TesT cAse", 2000, "~~~", "TesT cAse~~~"},
		{"TesT cAse", 2, "~~~", "~~~"},
		{"TesT cAse", 6, "---", "TesT---"},
		{"TesT cAse", 6, "世界", "TesT世界"},
	}
	for _, test := range tests {
		result := Prune(test.input, test.strLength, test.end)
		if result != test.expected {
			t.Errorf("\n\nGot: %s \nExpected: %s\n\n", result, test.expected)
		}
	}
}

func TestSlice(t *testing.T) {
	type LocalTestCase struct {
		input      string
		startIndex float64
		endVal     uint64
		expected   string
	}
	tests := []LocalTestCase{
		{"Test casE", 0, 3, "Tes"},
		{"Test casE", -5, 0, " casE"},
		{"Test casE", -5, 2000, " casE"},
		{"Test casE", 1, 7, "est ca"},
		{"Test casE", 1, 2000, "est casE"},
		{"Test casE", 1, 0, "est casE"},
		{"Test casE", 0, 0, "Test casE"},
		{"Hello, 世界", 0, 8, "Hello, 世"},
	}

	for _, test := range tests {
		result := Slice(test.input, test.startIndex, test.endVal)
		if result != test.expected {
			t.Errorf("\n\nGot: %s \nExpected: %s\n\n", result, test.expected)
		}
	}
}

func TestCount(t *testing.T) {
	type LocalTestCase struct {
		input    string
		expected int
	}
	tests := []LocalTestCase{
		{"Test", 4},
		{"Test ", 5},
		{"Hello, 世界", 9},
	}

	for _, test := range tests {
		result := Count(test.input)
		if result != test.expected {
			t.Errorf("\n\nGot: %d \nExpected: %d\n\n", result, test.expected)
		}
	}
}

func TestCountSubStrings(t *testing.T) {
	type LocalTestCase struct {
		input    string
		param    string
		expected int
	}

	tests := []LocalTestCase{
		{"Test Case", "Te", 1},
		{"tEsT CAsE", "tEsT ", 1},
		{"tEsT TEsT", "EsT", 2},
		{"TestTestTestTestTEsT", "Test", 4},
		{"TestTestTestTestTEsT", "T", 6},
		{"TesT cAse世界", "界", 1},
	}

	for _, test := range tests {
		result := CountSubStrings(test.input, test.param)
		if test.expected != result {
			t.Errorf("\n\nGot: %d \nExpected: %d\n\n", result, test.expected)
		}
	}
}
