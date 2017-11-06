package pathmonitor_test

import (
	"github.com/dargad/pathmonitor/internal/app/pathmonitor"
	"testing"
)

func TestPlaceholderReplacement(t *testing.T) {
	r := pathmonitor.ReplacePlaceholders("abc", "cde")
	AssertEqual(t, r, "abc", "failed")

	r = pathmonitor.ReplacePlaceholders("abc {}", "cde")
	AssertEqual(t, r, "abc cde", "")

	r = pathmonitor.ReplacePlaceholders("abc '{}'", "cde")
	AssertEqual(t, r, "abc '{}'", "")

	r = pathmonitor.ReplacePlaceholders("abc {}'", "cde")
	AssertEqual(t, r, "abc cde'", "")

	r = pathmonitor.ReplacePlaceholders("abc '{}", "cde")
	AssertEqual(t, r, "abc 'cde", "")

	r = pathmonitor.ReplacePlaceholders("abc '{}", "cde")
	AssertEqual(t, r, "abc 'cde", "")

	r = pathmonitor.ReplacePlaceholders("'{}' abc", "cde")
	AssertEqual(t, r, "'{}' abc", "")

	r = pathmonitor.ReplacePlaceholders("'{} abc", "cde")
	AssertEqual(t, r, "'cde abc", "")

	r = pathmonitor.ReplacePlaceholders("{}' abc", "cde")
	AssertEqual(t, r, "cde' abc", "")

	r = pathmonitor.ReplacePlaceholders("abc {} abc", "cde")
	AssertEqual(t, r, "abc cde abc", "")

	r = pathmonitor.ReplacePlaceholders("abc '{}' abc", "cde")
	AssertEqual(t, r, "abc '{}' abc", "")
}
