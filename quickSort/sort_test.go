package quickSort

import "testing"

func TestQuickSort(t *testing.T) {
	item := RandArray(10)
	data := QuickSort(item)
	t.Logf("%+v", data)
	
	data2 := QuickSort2(item)
	t.Logf("%+v", data2)
	
	SelectionSort(item, 10)
	t.Logf("%+v", item)
	
}

