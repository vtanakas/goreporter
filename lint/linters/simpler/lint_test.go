package simpler

import (
	"testing"

	"goreporter/lint/linters/simpler/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, NewChecker(), "")
}
