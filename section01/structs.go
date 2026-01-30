package section01

import "errors"

type Person struct {
	Wage	float32
	Name	string
	Age		int
}

func NewPerson(wage float32, name string, age int) (*Person, error) {
	if wage <= 0 {
		return nil, errors.New("Wage must be greater than 0")
	}
	if name == "" {
		return nil, errors.New("Name cannot be void")
	}
	if age <= 0 {
		return nil, errors.New("Age must be greater than 0")
	}

	return &Person{
		Name: name,
		Age: age,
		Wage: wage,
	}, nil
}