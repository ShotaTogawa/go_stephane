package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	name string
	age  int
	contactInfo
}

func main() {
	sam := person{name: "sam", age: 10}
	// person := person{"sam", 10} 上記と同じ
	fmt.Println(sam.name, sam.age)

	jim := person{
		name: "Jim",
		age:  20,
		contactInfo: contactInfo{
			email:   "jim@test.com",
			zipCode: 94000,
		},
	}

	// &variable: Give me the memory address of the value this variable is pointing at
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	jim.updateName("Jimmy")
	jim.print()

	// var alex person

	// alex.name = "Alex"
	// alex.age = 15
	// alex.contactInfo.email = "test@test.com"
	// alex.contactInfo.zipCode = 12345
	// fmt.Printf("%+v", alex)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newName string) {
	// ここで実行された内容は、RAMにコピーされる。
	// *pointer: Give me the value this memory address is pointing at
	(*pointerToPerson).name = newName
}
