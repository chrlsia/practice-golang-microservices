package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//Initialization:
	els := []int{9, 8, 7, 6, 5}
	//Execution:
	BubbleSort(els)
	//Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}
func TestBubbleSortWBestCase(t *testing.T) {
	//Initialization:
	els := []int{5, 6, 7, 8, 9}
	//Execution:
	BubbleSort(els)
	//Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortNilSlice(t *testing.T) {

	BubbleSort(nil)

}

//getElements creates a slice of n elements of integers
func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])

}

func BenchmarkBubbleSort10(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els) //8.986 ns/op
	}
}

func BenchmarkSort10(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		//sort function for  []int
		sort.Ints(els) //~ 251.3 ns/op
	}
}

func BenchmarkBubbleSort1000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els) // 690.6 ns/op
	}
}

func BenchmarkSort1000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els) //47054 ns/op
	}
}

func BenchmarkBubbleSort50000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els) // 690.6 ns/op
	}
}

func BenchmarkSort50000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els) //47054 ns/op
	}
}

func BenchmarkBubbleSort100000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		//BenchmarkBubbleSort100000-4   	       1	15113150236 ns/op
		BubbleSort(els) //
	}
}

func BenchmarkSort100000(b *testing.B) { //pop up suggestion cause is the only one exported
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		//BenchmarkSort100000-4   	     154	   7406533 ns/op
		sort.Ints(els)
	}
}

func Sort(els []int) {
	if len(els) < 50000 {
		BubbleSort(els)
		return
	} else {
		sort.Ints(els)
		return
	}
}

/*
Συμπέρασμα: Για λίγα data(κατω απο 1000) η BubbleSort είναι καλύτερη.
Αλλά για πολλά data η sort.Ints είναι πολύ καλύτερη
*/
