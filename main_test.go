package main

import "testing"

func Test_fuelDistance(t *testing.T) {
	tests := []struct {
		name            string
		fuelConsumption int
		fuelAvailable   int
		want            int
	}{
		{"Road rest", 10, 20, 200},
	}
	for _, test := range tests {

		got := fuelDistance(test.fuelConsumption, test.fuelAvailable)
		if got != test.want {
			t.Error("rest road", "got:", got, "want", test.want)
		}

	}
}
