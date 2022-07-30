package appleandorange

func CountApplesAndOranges(s int, t int, a int, b int, apples []int, oranges []int) (int, int) {
	appleCount := 0
	orangeCount := 0

	for _, apple := range apples {
		dropPoint := a + apple
		if IsTheFruitInsideTheArea(dropPoint, s, t) {
			appleCount++
		}
	}

	for _, orange := range oranges {
		dropPoint := b + orange
		if IsTheFruitInsideTheArea(dropPoint, s, t) {
			orangeCount++
		}
	}

	return appleCount, orangeCount
}

func IsTheFruitInsideTheArea(fruit int, s int, t int) bool {
	if fruit >= s && fruit <= t {
		return true
	}
	return false
}
