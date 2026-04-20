package main

import "testing"

// Note: We can call 'hello' function from hello_world.go in this file because they are part of the same package.

func TestHello(t *testing.T) {
	// Subtest 1: Basic greeting
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("omer", "English")
		expect := "Hello, omer"

		assertCorrectMessage(t, got, expect)
	})

	// Subtest 2: Empty string
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		expect := "Hello, World"
		assertCorrectMessage(t, got, expect)
	})

	// Subtest 3: Different langauge (spanish, french, arabic, ...)
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		expect := "Hola, Elodie"
		assertCorrectMessage(t, got, expect)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Kylian Mbappe", "French")
		expect := "Bonjour, Kylian Mbappe"
		assertCorrectMessage(t, got, expect)
	})

	t.Run("in Arabic", func(t *testing.T) {
		got := Hello("Omer", "Arabic")
		expect := "Marhaba, Omer"
		assertCorrectMessage(t, got, expect)
	})
}

func assertCorrectMessage(t testing.TB, got string, expect string) {
	if got != expect {
		t.Errorf("got %q but expect %q", got, expect)
	}
}