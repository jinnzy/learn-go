package test

import "testing"

func TestAdd(t *testing.T)  {
	r := add(1,4)
	if (r != 6) {
		t.Fatalf("add(2,4) error,expect:%d,actual:%d",6,r)
	}
	t.Logf("test add success")

}