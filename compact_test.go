package compact

import "testing"

type foo []string

func Test_Strings_on_foo(t *testing.T) {
	loose := foo{"", "a", "", "b", "", "c", ""}

	if len(loose) != 7 {
		t.Fatalf("%v", len(loose))
	}

	compact := Strings(loose)

	if len(compact) != 3 {
		t.Fatalf("didn't compact %v", len(compact))
	}

	if compact[0] != "a" {
		t.Fatalf("%v", compact[0])
	}

	if compact[1] != "b" {
		t.Fatalf("%v", compact[1])
	}

	if compact[2] != "c" {
		t.Fatalf("%v", compact[2])
	}

	// basicically, just shouldn't blow up
	Strings(nil)
	Strings(foo{})
	Strings(foo{""})
	Strings(foo{"", ""})
}

func Test_Strings_on_string_slice(t *testing.T) {
	loose := []string{"", "a", "", "b", "", "c", ""}

	if len(loose) != 7 {
		t.Fatalf("%v", len(loose))
	}

	compact := Strings(loose)

	if len(compact) != 3 {
		t.Fatalf("didn't compact %v", len(compact))
	}

	if compact[0] != "a" {
		t.Fatalf("%v", compact[0])
	}

	if compact[1] != "b" {
		t.Fatalf("%v", compact[1])
	}

	if compact[2] != "c" {
		t.Fatalf("%v", compact[2])
	}

	// basicically, just shouldn't blow up
	Strings(nil)
	Strings([]string{})
	Strings([]string{""})
	Strings([]string{"", ""})
}
