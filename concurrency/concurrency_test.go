

package concurrency

import (
	"reflect"
	"sort"
	"testing"
)


func TestSquareWorker(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 4, 9, 16, 25}


	actual := SquareWorker(input)

	sort.Ints(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}