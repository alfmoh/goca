package goca

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    string
	expected string
	param    string
}

type intTestCase struct {
	input    string
	expected string
	param    uint64
}

type intTestCases []intTestCase

type testCases []testCase

func runTest(tests testCases, f func(str string) string, t *testing.T) {
	for _, test := range tests {
		testName := fmt.Sprintf("%s, %s", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			res := f(test.input)
			if res != test.expected {
				t.Errorf("\n\nGot: %s \nExpected: %s\n\n", res, test.expected)
			}
		})
	}
}
func runTestWithParams(tests testCases, f func(str string, params ...string) string, t *testing.T) {
	for _, test := range tests {
		testName := fmt.Sprintf("%s, %s", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			res := f(test.input, test.param)
			if res != test.expected {
				t.Errorf("\n\nGot: %s \nExpected: %s\n\n", res, test.expected)
			}
		})
	}
}
func runTestWithIntParams(tests intTestCases, f func(str string, params uint64) string, t *testing.T) {
	for _, test := range tests {
		testName := fmt.Sprintf("%s, %s", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			res := f(test.input, test.param)
			if res != test.expected {
				t.Errorf("\n\nGot: %s \nExpected: %s\n\n", res, test.expected)
			}
		})
	}
}
