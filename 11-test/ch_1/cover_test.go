package ch_1

import "testing"

func TestCover(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		result := cover(2, 3)
		expected := 5
		if result != expected {
			t.Errorf("add(2, 3) returned %d, expected %d", result, expected)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		result := cover(3, 3)
		expected := 6
		if result != expected {
			t.Errorf("add(3, 3) returned %d, expected %d", result, expected)
		}
	})
}
