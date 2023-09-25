package patterns

import "log"

type IdServiceSingleton struct {
	idService *IdService
}

func (singleService *IdServiceSingleton) GetIdService() *IdService {
	// singleton checks
	if singleService.idService == nil {
		singleService.idService = NewIdService()
	}
	return singleService.idService
}

type IdService struct {
	counter int
}

func NewIdService() *IdService {
	return &IdService{counter: 0}
}

func (service *IdService) Next() int {
	service.counter++
	return service.counter
}

func buildCar(id int) {
	log.Printf("build car: %d", id)
}

func buildMotobike(id int) {
	log.Printf("build motobike: %d", id)
}

func Singleton() {
	var singleton IdServiceSingleton
	s0 := singleton.GetIdService()
	buildCar(s0.Next())
	buildCar(s0.Next())

	s1 := singleton.GetIdService()
	buildMotobike(s1.Next())
	buildMotobike(s1.Next())

	if s0 == s1 {
		log.Print("s0 and s1 is the same")
	}
}
