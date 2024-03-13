package sugar_test

import (
	"testing"

	. "github.com/thienhaole92/sugar"
)

func TestFind(t *testing.T) {
	{
		value, index := Find([]string{}, func(s string) bool {
			return len(s) != 0
		})
		AssertEqual(t, -1, index)
		AssertEqual(t, "", value)
	}
	{
		value, index := Find([]string{"a", "ab", "abc"}, func(s string) bool {
			return len(s) > 1
		})
		AssertEqual(t, 1, index)
		AssertEqual(t, "ab", value)
	}
}
