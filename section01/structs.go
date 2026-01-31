package section01

import "errors"

type Person struct {
	wage	float32
	name	string
	age		int
}

func NewPerson(wage float32, name string, age int) (*Person, error) {
	if err := validateWage(wage); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateAge(age); err != nil {
		return nil, err
	}

	return &Person{
		wage: wage,
		name: name,
		age: age,
	}, nil
}

func (p *Person) GetWage() float32 { return p.wage }
func (p *Person) GetName() string { return p.name }
func (p *Person) GetAge() int { return p.age }

func (p *Person) SetWage(wage float32) error {
	if err := validateWage(wage); err != nil {
		return err
	}
	p.wage = wage
	return nil
}

func (p *Person) SetName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	p.name = name
	return nil
}

func (p *Person) SetAge(age int) error {
	if err := validateAge(age); err != nil {
		return err
	}
	p.age = age
	return nil
}

func validateWage(wage float32) error {
	if wage < 0 {
		return errors.New("Wage cannot be less than 0")
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Name cannot be void")
	}
	return nil
}

func validateAge(age int) error {
	if age <= 0 {
		return errors.New("Age must be greater than 0")
	}
	return nil
}