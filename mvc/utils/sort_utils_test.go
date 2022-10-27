package utils

import (
	"github.com/stretchr/testify/assert"
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
