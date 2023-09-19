package patterns

import (
	"fmt"
	"log"
)

type IMemento interface {
	Restore() interface{}
}

type PersonMemento struct {
	person Person
}

func NewPersonMemento(person Person) *PersonMemento {
	return &PersonMemento{person: person}
}

func (memento PersonMemento) Restore() interface{} {
	return memento.person
}

type CareTaker struct {
	content []IMemento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{content: make([]IMemento, 0)}
}

func (c *CareTaker) Store(memento IMemento) {
	c.content = append(c.content, memento)
}

func (c *CareTaker) Last() (IMemento, error) {
	if len(c.content) == 0 {
		return nil, fmt.Errorf("empty memento")
	}
	lastIndex := len(c.content) - 1
	lastMemento := c.content[lastIndex]
	c.content = c.content[:lastIndex]
	return lastMemento, nil
}

type Person struct {
	lifePoints int
}

func NewPerson(lifePoints int) *Person {
	return &Person{lifePoints: lifePoints}
}

func (p *Person) Display() {
	log.Print("display life point: ", p.lifePoints)
}

func (p *Person) Reduce(reducePoint int) {
	log.Printf("reducing: %d", reducePoint)
	p.lifePoints = p.lifePoints - reducePoint
	log.Print("reduced life point: ", p.lifePoints)
}

type Originator interface {
	Save() IMemento
	Restore(IMemento)
}

type PersonOriginator struct {
	person Person
}

func NewPersonOriginator() *PersonOriginator {
	return &PersonOriginator{person: *NewPerson(100)}
}

func (originator *PersonOriginator) Save() IMemento {
	return NewPersonMemento(originator.person)
}

func (originator *PersonOriginator) Restore(m IMemento) {
	log.Print("Restore the saved person.")
	originator.person = m.Restore().(Person)
}

func (originator *PersonOriginator) Get() Person {
	return originator.person
}

func (originator *PersonOriginator) Set(person Person) {
	originator.person = person
}

func Memento() {
	// Init
	c := NewCareTaker()
	o := NewPersonOriginator()

	// first memento
	m0 := o.Save()
	c.Store(m0)

	h0 := o.Get()
	h0.Display()
	// reduce the life lifePoints
	h0.Reduce(75)

	// second memento
	o.Set(h0)
	m1 := o.Save()
	c.Store(m1)

	// reduce again
	h0.Reduce(50)
	h0.Display() // negative value

	// Now do the Restore
	lastMemento, err := c.Last()
	if err != nil {
		log.Fatal("restore m1 is failing")
	}
	o.Restore(lastMemento)

	restoreH1 := o.Get()
	restoreH1.Display() // lifePoints: 25

	lastMemento, err = c.Last()
	if err != nil {
		log.Fatal("restore m1 is failing")
	}
	o.Restore(lastMemento)

	restoreH0 := o.Get()
	restoreH0.Display() // lifePoints: 25
}
