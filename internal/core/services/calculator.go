package services

import "microservices/internal/domain"

type calculatorService struct {
	calculator domain.Calculator
}

// สร้างตัวอย่างของตัวคำนวณ
func NewCalculatorService() *calculatorService {
	return &calculatorService{calculator: domain.Calculator{}}
}

// การบวก
func (s *calculatorService) Add(a, b int) int {
	s.calculator.Result = a + b
	return s.calculator.Result
}

// การลบ
func (s *calculatorService) Subtract(a, b int) int {
	s.calculator.Result = a - b
	return s.calculator.Result
}

func (s *calculatorService) Error(err error) {
	s.calculator.Error = err.Error()
}
