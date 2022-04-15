package main

func main() {
	check := 5
	println(SquareOfSum(check))
	println(SumOfSquares(check))
}

func SquareOfSum(n int) int {
	var result int
	if n <= 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func SumOfSquares(n int) int {
	var result int
	if n <= 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}
