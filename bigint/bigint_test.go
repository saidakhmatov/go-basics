package bigint_test

import (
	"fmt"
	"testing"

	"github.com/saidakhmatov/golang/bigint"
)


func TestAddTD(t *testing.T) {

	tests := []struct {
		label string
		a, b  string
		want  string
		err   error
	}{
		{"test 1", "+++++++++111", "222", "333", nil},
		{"test 2", "000123", "1", "124", nil},
		{"test 3", "0000000000000000-111", "222", "111", nil},
		{"test 4", "sdafa", "adfs", "", bigint.ErrorBadInput},
		{"test 5", "7c89", "43+34", "", bigint.ErrorBadInput},

	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			var err error
			var a, b bigint.Bigint

			a, err = bigint.NewInt(tt.a)

			b, err = bigint.NewInt(tt.b)

			testcase := fmt.Sprintf("add: (%s,%s)", a.Value(), b.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), tt.err.Error())

				}

			} else {
				ans := bigint.Add(a, b)
				if ans.Value() != tt.want {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, ans.Value(), tt.want)
				}
			}
		})
	}
}

func TestSubTD(t *testing.T) {

	tests := []struct {
		label string
		a, b  string
		want  string
		err   error
	}{
		{"test 1", "+++++++++111", "222", "-111", nil},
		{"test 2", "000123", "1", "122", nil},
		{"test 3", "0000000000000000-111", "222", "-333", nil},
		{"test 4", "sdafa", "adfs", "", bigint.ErrorBadInput},
		{"test 5", "7c89", "43+34", "", bigint.ErrorBadInput},

	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			var err error
			var a, b bigint.Bigint

			a, err = bigint.NewInt(tt.a)

			b, err = bigint.NewInt(tt.b)

			testcase := fmt.Sprintf("sub: (%s,%s)", a.Value(), b.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), tt.err.Error())

				}

			} else {
				ans := bigint.Sub(a, b)
				if ans.Value() != tt.want {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, ans.Value(), tt.want)
				}
			}
		})
	}
}

func TestMultTD(t *testing.T) {

	tests := []struct {
		label string
		a, b  string
		want  string
		err   error
	}{
		{"test 1", "+++++++++111", "222", "24642", nil},
		{"test 2", "000123", "1", "123", nil},
		{"test 3", "0000000000000000-111", "222", "-24642", nil},
		{"test 4", "sdafa", "adfs", "", bigint.ErrorBadInput},
		{"test 5", "7c89", "43+34", "", bigint.ErrorBadInput},

	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			var err error
			var a, b bigint.Bigint

			a, err = bigint.NewInt(tt.a)

			b, err = bigint.NewInt(tt.b)

			testcase := fmt.Sprintf("Mult: (%s,%s)", a.Value(), b.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), tt.err.Error())

				}

			} else {
				ans := bigint.Multiply(a, b)
				if ans.Value() != tt.want {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, ans.Value(), tt.want)
				}
			}
		})
	}
}

func TestModTD(t *testing.T) {

	tests := []struct {
		label string
		a, b  string
		want  string
		err   error
	}{
		{"test 1", "+++++++++111", "222", "111", nil},
		{"test 2", "000123", "33", "24", nil},
		{"test 3", "00000000000000004654655464655", "89233", "18101", nil},
		{"test 4", "sdafa", "adfs", "", bigint.ErrorBadInput},
		{"test 5", "7c89", "43+34", "", bigint.ErrorBadInput},

	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			var err error
			var a, b bigint.Bigint

			a, err = bigint.NewInt(tt.a)

			b, err = bigint.NewInt(tt.b)

			testcase := fmt.Sprintf("Mod: (%s,%s)", a.Value(), b.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), tt.err.Error())

				}

			} else {
				ans := bigint.Mod(a, b)
				if ans.Value() != tt.want {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, ans.Value(), tt.want)
				}
			}
		})
	}
}

func TestAbsTD(t *testing.T) {

	tests := []struct {
		label string
		a  string
		want  string
		err   error
	}{
		{"test 1", "+++++++++111", "111", nil},
		{"test 2", "000-123", "123", nil},
		{"test 3", "0000000000000000-111", "111",  nil},
		{"test 4", "sdafa", "adfs",  bigint.ErrorBadInput},
		{"test 5", "7c89", "43+34",  bigint.ErrorBadInput},

	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			var err error
			var a bigint.Bigint

			a, err = bigint.NewInt(tt.a)

			testcase := fmt.Sprintf("abs: (%s)", a.Value())

			if err != nil {
				if tt.err == nil {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), "nil")
				} else if err.Error() != tt.err.Error() {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, err.Error(), tt.err.Error())

				}

			} else {
				
				ans := a.Abs()
				if ans.Value() != tt.want {
					t.Errorf("test: %s, result: %s, expected: %s ", testcase, ans.Value(), tt.want)
				}
			}
		})
	}
}

