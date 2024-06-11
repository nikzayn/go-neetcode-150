package slowandfastpointers

func isHappy(num int) bool {
	slow := num
	fast := sumOfSquaredDigits(num)

	for fast != 1 || slow != fast {
		slow = sumOfSquaredDigits(slow)
		fast = sumOfSquaredDigits(sumOfSquaredDigits(fast))
	}

	if fast == 1 {
		return true
	}

	return false
}

func pow(digit, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}
