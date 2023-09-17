package patterns

import (
	"log"
)

type LengthChecker struct {
	StandardLength int
}

func NewLengthChecker(length int) *LengthChecker {
	return &LengthChecker{StandardLength: length}
}

func (l LengthChecker) IsValid(number string) bool {
	if len(number) != l.StandardLength {
		return false
	}
	return true
}

type Service struct {
	lengthChecker *LengthChecker
}

func NewService(number int) *Service {
	return &Service{lengthChecker: NewLengthChecker(number)}
}

func (s Service) process(number string) {
	if ok := s.lengthChecker.IsValid(number); ok {
		log.Print(ok)
	}
}

func Facade() {
	service := NewService(10)
	service.process("0123456789")
}
