package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	// try a unit test on function
	if i, e := Division(6, 2); i != 3 || e != nil {
		// If it is not as expected, then the test has fail
		t.Error("divison function tests do not pass")
	} else {
		// record the expected information
		t.Log("first test passed")
	}
}

func Test_Division_2(t *testing.T) {
	// try a unit test on function
	if _, e := Division(6, 0); e == nil {
		// If it is not as expected, then the error
		t.Error("Divison did not work as expected.")
	} else {
		// record some of the information you expect to record
		t.Log("one test passed.", e)
	}
}