package sugar_test

import (
	"fmt"
	"strings"
	"testing"
)

func fail(t *testing.T, failMsg string, tags ...string) {
	t.Helper()
	if len(tags) == 0 {
		t.Error(failMsg)
	} else {
		t.Errorf("%s, tags: %s", failMsg, strings.Join(tags, ", "))
	}
}

func AssertEqual[T comparable](t *testing.T, expected T, actual T, tags ...string) {
	t.Helper()
	if expected != actual {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
	}
}

func AssertSliceEqual[T comparable](t *testing.T, expected []T, actual []T) {
	t.Helper()

	if expected == nil || actual == nil {
		if expected == nil && actual == nil {
			return
		}
		t.Errorf("expected: %v, actual: %v", expected, actual)
		return
	}

	if len(expected) != len(actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
		return
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected: %v, actual: %v", expected, actual)
			return
		}
	}
}
