package merge_sort

import "testing"

func TestSort(t *testing.T) {
	// good_slc := []int{38, 27, 43, 3, 9, 82, 10}
	good_slc := []int{-356, 328, 705, -199, -373, 108, -377, -362, 128, 98, 1, -9, -500, -607, 387, 12, 210, -600, -351, 432}
	got := sort(good_slc)

	if len(got) != len(good_slc) {
		t.Errorf("TestingSort: Expected length %v, got %v", len(good_slc), len(got))
	}

	for idx := range got {
		if idx != 0 && got[idx] < got[idx-1] {
			t.Errorf("TestingSort: Slice not sorted. Returned slice: %v", got)
		}
	}
}
