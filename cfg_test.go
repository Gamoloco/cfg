package cfg

import (
	"reflect"
	"testing"
)

type nested2 struct {
	A string
	B int
	C bool
	D []string
	E float64
}

type nested1 struct {
	A string
	B int
	C bool
	D []string `json:"D" cfg:"optional"`
	E float64
	F nested2
}

type test struct {
	A string
	B int
	C bool
	D []string
	E float64

	F nested1
}

func testValid(t *testing.T) {
	var tes test
	err := Load("./tests/test_valid.json", &tes)

	if err != nil {
		t.Error(err)
	} else {
		testCompare := test{
			A: "hello",
			B: 4,
			C: true,
			D: []string{"a", "b"},
			E: 4.444,
			F: nested1{
				A: "hello",
				B: 4,
				C: true,
				D: []string{"a", "b"},
				E: 4.444,
				F: nested2{
					A: "hello",
					B: 4,
					C: true,
					D: []string{"a", "b"},
					E: 4.444,
				},
			},
		}

		if !reflect.DeepEqual(tes, testCompare) {
			t.Error("Configuration file struct does not match struct")
		}
	}
}

func testMissing(t *testing.T) {
	var tes test
	err := Load("./tests/test_missing.json", &tes)

	if err == nil {
		t.Error(err)
	} else {
		if err.Error() != "Missing required field: D" {
			t.Error("Invalid missing error")
		}
	}
}

func testMissingNested(t *testing.T) {
	var tes test
	err := Load("./tests/test_missing_nested.json", &tes)

	if err == nil {
		t.Error(err)
	} else {
		if err.Error() != "Missing required field: F.F.D" {
			t.Error("Invalid missing error")
		}
	}
}

func testOptional(t *testing.T) {
	var tes test
	err := Load("./tests/test_optional.json", &tes)

	if err != nil {
		t.Error(err)
	} else {
		testCompare := test{
			A: "hello",
			B: 4,
			C: true,
			D: []string{"a", "b"},
			E: 4.444,
			F: nested1{
				A: "hello",
				B: 4,
				C: true,
				E: 4.444,
				F: nested2{
					A: "hello",
					B: 4,
					C: true,
					D: []string{"a", "b"},
					E: 4.444,
				},
			},
		}

		if !reflect.DeepEqual(tes, testCompare) {
			t.Error("Configuration file struct does not match struct")
		}
	}
}

func TestLoad(t *testing.T) {
	testValid(t)
	testMissing(t)
	testMissingNested(t)
	testOptional(t)
}
