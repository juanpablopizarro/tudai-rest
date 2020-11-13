package greeter

import "testing"

// TestSayHello ...
func TestSayHello(t *testing.T) {
	// Instancio el servicio
	var s Greeter = greeter{}

	// defino lo que espero
	expected := "Hello Alice!"

	// ejecuto el servicio y obtengo el resultado real (actual)
	actual := s.SayHello("Alice")

	// comparo actual con expected
	if expected != actual {
		t.Errorf("\"%v\" != \"%v\"", actual, expected)
	}

	t.Log(":D")

}
