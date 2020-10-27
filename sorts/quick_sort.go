package sorts

// 挖坑，填坑，分治算法
//
// 1. 将最左边的元素挖出，做为基准元素。
// 2. 从最右边元素开始往左遍历，直到找到一个小于基准元素的元素。并将该元素挖出，放到左边，左边往右移动一个元素。
// 3. 从最左边元素开始往右遍历，直到找到一个大于基准元素的元素。将改元素挖出，放到右边，右边往左边移动一个元素。
// 4. 重复执行第2，第3步，直到左右移动到同一个元素。把基准元素放入到该元素。
// 5. 将数组按第4步元素分成两半，重新执行第1, 2, 3, 4步。
func QuickSort(arr []int) {
	quickSortPartial(arr, 0, len(arr)-1)
}

func quickSortPartial(arr []int, left int, right int) {
	if left >= right {
		return
	}

	i := left
	j := right
	pivot := arr[i]
	for i < j {
		for ; i < j && arr[j] >= pivot; j-- {
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		for ; i < j && arr[i] < pivot; i++ {
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	quickSortPartial(arr, left, i-1)
	quickSortPartial(arr, i+1, right)
}
