package ports

type CalculatorService interface {
	Add(a, b int) int
	Subtract(a, b int) int
}
