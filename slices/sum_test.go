package main

import (
	"reflect"

	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4})
	want := []int{6, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
