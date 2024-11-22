package brackets_test

import (
	"testing"

	"github.com/Friedfisch/brackets"
)

func cB1(t *testing.T, n string, s string, e bool) {
	r := brackets.IsBalanced(s)
	if r != e {
		t.Errorf("%s: %s should be %v", n, s, e)
	} else {
		t.Logf("%s ok", n)
	}
}

func TestSolve(t *testing.T) {
	cB1(t, "sf 5", "(***(()", false)

	cB1(t, "sok 1", "(*", true)
	cB1(t, "sok 2", "(*))", true)
	cB1(t, "sok 3", "****", true)
	cB1(t, "sok 4", "(***", true)
	cB1(t, "sok 5", "((*)*()", true)
	cB1(t, "sf 1", "((*", false)
	cB1(t, "sf 2", ")*(", false)
	cB1(t, "sf 3", "*(", false)
	cB1(t, "sf 4", "***(", false)
	cB1(t, "sf 5", "(***(()", false)

	cB1(t, "empty", "", true)
	cB1(t, "no data", "123", true)
	cB1(t, "f 1", ")", false)
	cB1(t, "f 2", "(", false)
	cB1(t, "f 3", ")(", false)
	cB1(t, "f 4", "(())(())(((()))()()()())))()", false)
	cB1(t, "ok 1", "()", true)
	cB1(t, "ok 2", "(())", true)
	cB1(t, "ok 3", "(())(())(((()))()()()())()", true)
}
