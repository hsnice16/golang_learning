package gotest

import "testing"

func Test_Divison_1(t *testing.T) {
	// try a unit test on function
	if i, e := Divison(6, 2); i != 3 || e != nil {
		// If it is not as expected, then the test has fail
		t.Error("divison function tests do not pass")
	} else {
		// record the expected information
		t.Log("first test passed")
	}
}

func Test_Divison_2(t *testing.T) {
	t.Error("just does not pass")
}