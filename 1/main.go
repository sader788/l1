package main

import "fmt"

type Human struct {
	firstName  string
	secondName string
	age        int
}

type Action struct {
	name string
	Human
	h Human
}

type Action2 struct {
	name string
	Human
	h Human
}

type NamePrinter interface {
	printName()
}

func (h *Human) say(str string) {
	fmt.Println(str)
}

func (h *Human) printName() {
	fmt.Println(h.firstName)
}

func (a *Action) printName() {
	fmt.Println(a.name)
}

func main() {
	act := Action{
		name:  "action",
		Human: Human{"John", "Williams", 47},
		h: Human{
			firstName:  "Vanya",
			secondName: "Pupkin",
			age:        18,
		},
	}

	fmt.Println("--------------------example1--------------------")
	//Если метод переопределен для дочерней структуры, то будет вызван ее метод
	act.printName()
	//Чтобы вызвать метод родительской структуры, нужно сделатт это явно
	act.Human.printName()
	//Также можно вызвать метод у поля типа родительской структуры
	act.h.printName()

	act2 := Action2{
		name:  "action",
		Human: Human{"John", "Williams", 47},
		h: Human{
			firstName:  "Vanya",
			secondName: "Pupkin",
			age:        18,
		},
	}
	fmt.Println("--------------------example2--------------------")
	//Если метод не переопределен дочерней структуры, то будет вызван метод родительской структуры
	act2.printName()
	act2.Human.printName()
	act2.h.printName()
}
