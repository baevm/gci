package object

import "testing"

func Test_StringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}

	test1 := &String{Value: "Testing Test"}
	test2 := &String{Value: "Testing Test"}

	if hello1.Hashkey() != hello2.Hashkey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if test1.Hashkey() != test2.Hashkey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.Hashkey() == test1.Hashkey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
