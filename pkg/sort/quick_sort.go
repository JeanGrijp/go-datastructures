package sort

// Troca dois elementos do array
func trocar(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// Particiona o array usando a abordagem Hoare
func particionar(array []int, left, right int) int {
	pivo := array[left]
	i := left - 1
	j := right + 1

	for {
		for {
			j--
			if array[j] <= pivo {
				break
			}
		}
		for {
			i++
			if array[i] >= pivo {
				break
			}
		}
		if i < j {
			trocar(array, i, j)
		} else {
			return j
		}
	}
}

// Implementa QuickSort recursivamente
func QuickSort(array []int, left, right int) {
	if left < right {
		p := particionar(array, left, right)
		QuickSort(array, left, p)
		QuickSort(array, p+1, right)
	}
}
