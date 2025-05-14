package main

import (
	"fmt"
)

type Person struct {
	// Exported fields
	Name string
	Age  int

	// Unexported fields
	salary int64
}

func NewPerson(name string, age int, salary int64) Person {
	return Person {
		Name: name,
		Age: age,
		salary: salary * 100,
	}
}

func main() {
	// type is Person, can't be nil
	fmt.Println("---------- Value type------------")

	yusuf := Person{
		Name:   "Yusuf",
		Age:    120,
		salary: 1,
	}

	fmt.Println(yusuf.Description())
	fmt.Println(Description(&yusuf))

	yusuf.SetSalaryValue(200)
	fmt.Println(yusuf.Description())

	yusuf.SetSalary(201)
	fmt.Println(yusuf.Description())

	fmt.Println("---------- PTR type------------")

	// type is pointer to person, can be nil
	// zero value is nil
	var budi *Person
	fmt.Println(budi.Description())
	budi = &Person{
		Name:   "Budi",
		Age:    70,
		salary: 300,
	}

	fmt.Println(Description(budi))
	fmt.Println(DescriptionValue(*budi))
}

// Accept pointer type
func Description(p *Person) string {
	return fmt.Sprintf("%s, age: %d, salary: %d", p.Name, p.Age, p.salary)
}

// Accept value type
func DescriptionValue(p Person) string {
	return fmt.Sprintf("%s, age: %d, salary: %d", p.Name, p.Age, p.salary)
}

// receiver function
func (p Person) Description() string {
	return fmt.Sprintf("%s, age: %d, salary: %d", p.Name, p.Age, p.salary)
}

func (p *Person) SetSalary(salary int64) {
	p.salary = salary
}

func (p Person) SetSalaryValue(salary int64) {
	p.salary = salary
}
