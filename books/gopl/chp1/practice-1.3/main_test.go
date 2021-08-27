package main

import "testing"

func Test_printArgs1(t *testing.T) {
	testCases := []struct {
		description string
		params      []string
		exception   string
	}{
		{
			description: "Test_printArgs1-case1",
			params:      []string{"case1", "1", "2", "3"},
			exception:   "case1 1 2 3",
		},
		{
			description: "Test_printArgs1-case2",
			params:      []string{"case2", "4", "5", "6"},
			exception:   "case2 4 5 6",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			s := printArgs1(tc.params)
			if s != tc.exception {
				t.Errorf("unexpected s: %s", s)
			}
		})
	}
}

func Test_printArgs2(t *testing.T) {
	testCases := []struct {
		description string
		params      []string
		exception   string
	}{
		{
			description: "Test_printArgs2-case1",
			params:      []string{"case1", "1", "2", "3"},
			exception:   "case1 1 2 3",
		},
		{
			description: "Test_printArgs2-case2",
			params:      []string{"case2", "4", "5", "6"},
			exception:   "case2 4 5 6",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			s := printArgs1(tc.params)
			if s != tc.exception {
				t.Errorf("unexpected s: %s", s)
			}
		})
	}
}
