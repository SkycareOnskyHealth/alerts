package alerts

import (
	"testing"
)

func TestErrors(t *testing.T) {
	ok, err := CheckAlert("invalid_operator", "1", "<")
	if ok || err == nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">", "invalid", "1")
	if ok || err == nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">", "1", "invalid")
	if ok || err == nil {
		t.Fatal("fail")
	}

	ok, err = CheckAlert(">", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("==", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<=", "1", "0")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("!=", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}

	ok, err = CheckAlert(">", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<", "1", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("==", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("!=", "1", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
}
