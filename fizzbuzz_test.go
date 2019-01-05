package fizzbuzz_test

import (
	"testing"
	"github.com/WarisR/fizzbuzz"
)

func TestFizzBuzzShouldSay1(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSay2(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizzWhenInputIs3(t *testing.T) {
	result := fizzbuzz.Say(3)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSay4(t *testing.T) {
	result := fizzbuzz.Say(4)
	expected := "4"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayBuzzWhenInputIs5(t *testing.T) {
	result := fizzbuzz.Say(5)
	expected := "Buzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizzWhenInputIs6(t *testing.T) {
	result := fizzbuzz.Say(6)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
