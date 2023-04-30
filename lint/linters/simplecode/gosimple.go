// FIXME: (issue #2) investigate need for this import directive.
// package simplecode // import "github.com/360EntSecGroup-Skylar/goreporter/linters/simplecode"
package simplecode

import (
	"goreporter/lint/linters/simplecode/lint/lintutil"
	"goreporter/lint/linters/simplecode/simple"
)

func Simple(path map[string]string, except string) []string {
	var res []string
	for _, p := range path {
		res = append(res, lintutil.ProcessArgs(except, "gosimple", simple.Funcs, []string{p})...)
	}
	return res
}
