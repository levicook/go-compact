package compact

import "testing"

func Test_Strings(t *testing.T) {
	loose := []string{"a", "", "b", "", "c"}

	if len(loose) != 5 {
		t.Fatalf("%s", len(loose))
	}

	compact := Strings(loose)

	if len(compact) != 3 {
		t.Fatalf("%s", len(compact))
	}

	if len(loose) != 5 {
		t.Fatalf("%s", len(loose))
	}
}
