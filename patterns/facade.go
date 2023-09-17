package patterns

import (
	"log"
	"strings"
)

type LengthChecker struct {
	StandardLength int
}

func NewLengthChecker(length int) *LengthChecker {
	return &LengthChecker{StandardLength: length}
}

func (l LengthChecker) IsValid(number string) bool {
	return len(number) == l.StandardLength
}

type PrefixValidator struct {
	prefix string
}

func NewPrefixValidator(prefix string) *PrefixValidator {
	return &PrefixValidator{prefix: prefix}
}

func (validator PrefixValidator) IsValid(number string) bool {
	return strings.HasPrefix(number, validator.prefix)
}

type Service struct {
	lengthChecker   *LengthChecker
	prefixValidator *PrefixValidator
}

func NewService(number int, prefix string) *Service {
	return &Service{
		lengthChecker:   NewLengthChecker(number),
		prefixValidator: NewPrefixValidator(prefix),
	}
}

func (s Service) process(number string) {
	ok := s.lengthChecker.IsValid(number) && s.prefixValidator.IsValid(number)
	log.Print(ok)
}

func Facade() {
	service := NewService(10, "123-")
	service.process("123-567890")
	service.process("123456789")
}
