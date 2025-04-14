package generics

import (
	"reflect"
	"testing"
)

func runReverseTest[T comparable](t *testing.T, input, expected []T) {
	t.Helper()
	result := Reverse(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", input, result, expected)
	}
}

func TestReverse(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		runReverseTest(t, []int{1, 2, 3}, []int{3, 2, 1})
	})

	t.Run("string slice", func(t *testing.T) {
		runReverseTest(t, []string{"a", "b", "c"}, []string{"c", "b", "a"})
	})
}
