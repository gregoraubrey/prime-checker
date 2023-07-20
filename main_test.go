package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, expected false but got true.", 0)
	}

	if msg != "0 is not prime." {
		t.Errorf("with %d as test parameter, expected msg of '0 is not prime.' but got msg of '%s'.", 0, msg)
	}

	result, msg = isPrime(97)
	if !result {
		t.Errorf("with %d as test parameter, expected true but got false.", 97)
	}

	if msg != "97 is prime!" {
		t.Errorf("with %d as test parameter, expected msg of '97 is prime!' but got msg of '%s'.", 97, msg)
	}

	result, msg = isPrime(100)
	if result {
		t.Errorf("with %d as test parameter, expected false but got true.", 100)
	}

	if msg != "100 is not prime because it is divisible by 2." {
		t.Errorf("with %d as test parameter, expected msg of '100 is not prime because it is divisible by 2.' but got msg of '%s'.", 100, msg)
	}
}
