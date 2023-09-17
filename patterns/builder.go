package patterns

import "log"

type human struct {
	age      int
	height   int
	eyeColor string
}

func NewHuman() human {
	return human{}
}

func NewHumanWithFields(age int, height int, eyeColor string) human {
	return human{age: age, height: height, eyeColor: eyeColor}
}

func (h human) withEyeColor(color string) human {
	h.eyeColor = color
	return h
}

func (h human) withAge(age int) human {
	h.age = age
	return h
}

func (h human) withHeight(height int) human {
	h.height = height
	return h
}

func Builder() {
	human := NewHuman().withAge(24).withHeight(180).withEyeColor("brown")
	log.Print(human)
}
