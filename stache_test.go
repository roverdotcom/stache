package main

import "testing"

func TestParseContextStrings(t *testing.T) {
	s := parseContextStrings([]string{"test=one=two"})
	if s["test"] != "one=two" {
		t.Error(s)
	}

	s = parseContextStrings([]string{"test=one", "test=two"})
	if s["test"] != "two" {
		t.Error(s)
	}

	s = parseContextStrings([]string{"a=one", "b=two"})
	if s["a"] != "one" && s["b"] != "two" {
		t.Error(s)
	}

	s = parseContextStrings([]string{"hello", "world=test"})
	if s["world"] != "test" {
		t.Error(s)
	}
}
