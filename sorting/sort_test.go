package sorting_test

import (
	"github.com/Kardbord/ubiquity/sliceutils"
	"github.com/Kardbord/ubiquity/sorting"
	"testing"
)

func TestSort(t *testing.T) {
	for impl := sorting.SortImpl(0); impl < sorting.MergeImplEnd; impl++ {
		if !sliceutils.Equal(sorting.Sort(impl, "d", "a", "b", "c"), []string{"a", "b", "c", "d"}) {
			t.Error()
		}
	}
}
