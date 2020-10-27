package sorts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 7, 5, 10, 1, 2, 3, 11, 4}
	QuickSort(arr)
	assert.Equal(t, []int{1, 2, 3, 3, 4, 5, 7, 10, 11}, arr)
}
