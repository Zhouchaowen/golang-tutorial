package ch_1

import "testing"

// go test
func TestAdd0(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) returned %d, expected %d", result, expected)
	}
}

func TestAdd1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		result := add(2, 3)
		expected := 5
		if result != expected {
			t.Errorf("add(2, 3) returned %d, expected %d", result, expected)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		result := add(3, 3)
		expected := 6
		if result != expected {
			t.Errorf("add(3, 3) returned %d, expected %d", result, expected)
		}
	})
}

func TestAdd2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1+2",
			args{
				a: 1,
				b: 2,
			},
			3,
		},
		{
			"10+10",
			args{
				a: 10,
				b: 10,
			},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
