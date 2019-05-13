package splitter

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const arrElements = 999

func genArray(elems int) []int {
	var res []int

	for i := 1; i <= elems; i++ {
		res = append(res, i)
	}

	return res
}

func countElems(arr [][]int) (int, bool) {
	var toCount []int
	for i := range arr {
		for j := range arr[i] {
			toCount = append(toCount, arr[i][j])
		}
	}

	if len(toCount) == arrElements {
		return len(toCount), true
	} else {
		return len(toCount), false
	}
}

func TestArr(t *testing.T) {
	Convey("Array should split correctly", t, func() {
		Convey("Divided array on 3 arrays must be joined and equal to arrElements", func() {
			testArr := genArray(arrElements)
			data := SplitToQueue(testArr, 3)
			count, _ := countElems(data)
			So(count, ShouldEqual, arrElements)
		})
		Convey("Divided array on 6 arrays must be joined and equal to arrElements", func() {
			testArr := genArray(arrElements)
			data := SplitToQueue(testArr, 6)
			count, _ := countElems(data)
			So(count, ShouldEqual, arrElements)
		})
		Convey("Divided array on 12 arrays must be joined and equal to arrElements", func() {
			testArr := genArray(arrElements)
			data := SplitToQueue(testArr, 12)
			count, _ := countElems(data)
			So(count, ShouldEqual, arrElements)
		})
	})
}
