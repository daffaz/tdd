package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add two nums", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d but got %d", sum, expected)
		}
	})
}
