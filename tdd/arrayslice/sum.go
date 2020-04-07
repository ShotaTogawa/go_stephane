package arrayslice

// Sum get arrays and iterate the numbers to return sum
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
