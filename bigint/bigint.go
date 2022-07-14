package bigint

import (
	"errors"
	"math/big"
	"strings"
)

type Bigint struct {
	value string
}

// declare error
var ErrorBadInput = errors.New("bad input")

// validate function
func validate(s *string) error {
	allowed := "1234567890"

	str := *s

	// shows error if first letter/number of string is not +- or number
	if !strings.Contains("+-"+allowed, string(str[0])) {
		return ErrorBadInput
	}

	// shows error if string  contains letters or signs (starts checking from index 1)
	for i := 1; i < len(str); i++ {
		if !strings.Contains(allowed, string(str[i])) {
			return ErrorBadInput
		}

	}

	return nil
}

// function to validate and return correct form of string
func NewInt(num string) (Bigint, error) {

	// removes leading '0's and '+' signs from string
	s := strings.TrimLeft(num, "0+")

	if err := validate(&s); err != nil {
		return Bigint{}, err
	}
	return Bigint{value: s}, nil
}

// function to subtract numbers then return result
func Add(a, b Bigint) Bigint {

	ba, bb := big.NewInt(10), big.NewInt(10)
	ad, _ := ba.SetString(a.value,10)
	
	ab, _ := bb.SetString(b.value, 10)

	var sum big.Int
	sum.Add(ad, ab)

	return Bigint{
		value: sum.String(),
	}
}

// function to add numbers then return result
func Sub(a, b Bigint) Bigint {

	ba, bb := big.NewInt(10), big.NewInt(10)

	ad, _ := ba.SetString(a.value,10)
	
	ab, _ := bb.SetString(b.value, 10)

	var diff big.Int
	diff.Sub(ad, ab)

	return Bigint{
		value: diff.String(),
	}
}

// function to return the result of multiplicaton of numbers
func Multiply(a, b Bigint) Bigint {

	ba, bb := big.NewInt(10), big.NewInt(10)

	ad, _ := ba.SetString(a.value,10)
	
	ab, _ := bb.SetString(b.value, 10)

	var multp big.Int
	multp.Mul(ad, ab)

	return Bigint{
		value: multp.String(),
	}
}

// function to return the result of finding Mod of numbers
func Mod(a, b Bigint) Bigint {

	ba, bb := big.NewInt(10), big.NewInt(10)

	ad, _ := ba.SetString(a.value,10)
	
	ab, _ := bb.SetString(b.value, 10)

	var mod big.Int
	mod.Mod(ad, ab)

	return Bigint{
		value: mod.String(),
	}
}

// function to return absolute value of number
func (x *Bigint) Abs() Bigint {

	val := x.value
	if val[0] == '-' {
		return Bigint{

			value: val[1:],
		}
	}
	return Bigint{
		value: val,
	}
}

func (z *Bigint) Value() string {

	return z.value
}
