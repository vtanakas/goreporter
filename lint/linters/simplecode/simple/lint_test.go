package simple

import (
	"testing"

	"goreporter/lint/linters/simplecode/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, Funcs, "../../")
}
