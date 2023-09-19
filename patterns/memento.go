package patterns

import "log"

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

type Person struct {
	lifePoints int
}

func NewPerson(lifePoints int) *Person {
	return &Person{lifePoints: lifePoints}
}

func (p *Person) Display() {
	log.Print("life point: ", p.lifePoints)
}

func (p *Person) Reduce(reducePoint int) {
	p.lifePoints = p.lifePoints - reducePoint
	log.Print("reduced life point: ", p.lifePoints)
}

type PersonOriginator struct {
	person Person
}

func NewPersonOriginator(person Person) *PersonOriginator {
	return &PersonOriginator{person: person}
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

func Memento() {
	h := NewPerson(100)
	h.Display()
	h.Reduce(50)
	h.Display()

	o := NewPersonOriginator(*h)
	m := o.Save()

	h.Reduce(25)
	h.Display()

	o.Restore(m)
	snapshotPerson := o.Get()
	snapshotPerson.Display()
}
