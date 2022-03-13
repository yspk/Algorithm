package strnull

import "testing"

func TestStringNull(t *testing.T) {
	t.Log(StringNull("we are happy"))
	t.Log(StringNull(" are happy"))
	t.Log(StringNull("we  "))
	t.Log(StringNull("wo     chen"))
	t.Log(StringNull("happy"))
	t.Log(StringNull(" "))
}
