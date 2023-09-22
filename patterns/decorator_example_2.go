package patterns

import "fmt"

// MainDish = $50
// SideDish = $15
// Staple = $10

type IMeal interface {
	Add(iPack int) int
}

type MainDish struct {
	IMeal
}

func (maindish *MainDish) Add(iPack int) int {
	return iPack * 50
}

type SideDish struct {
	IMeal
}

func (sidedish *SideDish) Add(iPack int) int {
	return iPack * 15
}

type Staple struct {
	IMeal
}

func (staple *Staple) Add(iPack int) int {
	return iPack * 10
}

type LunchBox struct {
	price int
}

func (lb *LunchBox) Calculate() int {
	return lb.price
}

func (lb *LunchBox) Order(meal IMeal, iPack int) {
	lb.price += meal.Add(iPack)
}

func Decorator2() {
	lunchbox := LunchBox{}
	lunchbox.Order(&MainDish{}, 2)
	lunchbox.Order(&SideDish{}, 3)
	lunchbox.Order(&Staple{}, 1)
	fmt.Printf("Total calculate price: $%d \n", lunchbox.price)
}
