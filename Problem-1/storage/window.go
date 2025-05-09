package storage

type NumberWindow struct {
	size    int
	numbers []int
	unique  map[int]bool
}

func NewNumberWindow(size int) *NumberWindow {
	return &NumberWindow{
		size:    size,
		numbers: []int{},
		unique:  make(map[int]bool),
	}
}

func (nw *NumberWindow) AddNumbers(nums []int) {
	for _, num := range nums {
		if !nw.unique[num] {
			if len(nw.numbers) >= nw.size {
				old := nw.numbers[0]
				nw.numbers = nw.numbers[1:]
				delete(nw.unique, old)
			}
			nw.numbers = append(nw.numbers, num)
			nw.unique[num] = true
		}
	}
}

func (nw *NumberWindow) GetNumbers() []int {
	copyNums := make([]int, len(nw.numbers))
	copy(copyNums, nw.numbers)
	return copyNums
}