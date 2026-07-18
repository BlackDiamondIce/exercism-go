package cards

// FavoriteCards возвращает слайс с любимыми картами: 2, 6 и 9.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem возвращает карту по указанному индексу.
// Если индекс вне границ, возвращает -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem изменяет карту в слайсе по указанному индексу.
// Если индекс вне диапазона, добавляет карту в конец.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems добавляет новые карты в самое начало слайса.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem удаляет карту из слайса по индексу, склеивая две части.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
