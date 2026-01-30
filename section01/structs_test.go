package section01

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStruct(t *testing.T) {
	t.Run("Validation Errors", func(t *testing.T) {
		t.Run("should return error when wage is equal to or less than zero", func (t *testing.T) {
			require := require.New(t)
	
			testCases := []struct {
				testName 	string
				wage		float32
				name		string
				age			int
			}{
				{"wage is zero", 0, "Kael", 26},
				{"wage is negative", -1, "Kael", 26},
			}
			
			for _, test := range testCases {
				t.Run(test.testName, func(t *testing.T) {
					_, err := NewPerson(test.wage, test.name, test.age)
					require.Error(err)
					require.EqualError(err, "Wage must be greater than 0")
				})
			}
		})
		t.Run("should return error when age is equal to or less than zero", func (t *testing.T) {
			require := require.New(t)
	
			testCases := []struct {
				testName 	string
				wage		float32
				name		string
				age			int
			}{
				{"Age is zero", 4800, "Kael", 0},
				{"Age is negative", 4800, "Kael", -1},
			}
			
			for _, test := range testCases {
				t.Run(test.testName, func(t *testing.T) {
					_, err := NewPerson(test.wage, test.name, test.age)
					require.Error(err)
					require.EqualError(err, "Age must be greater than 0")
				})
			}
		})
		t.Run("should return error when name is void", func (t *testing.T) {
			require := require.New(t)

			_, err := NewPerson(4800, "", 26)

			require.Error(err)
			require.EqualError(err, "Name cannot be void")
		})
	})
}

func ExampleNewPerson() {
	person, _ := NewPerson(4800, "Kael", 26)
	fmt.Println(person)
	// Output: &{4800 Kael 26}
}

func BenchmarkNewPerson(b *testing.B) {
	for b.Loop() {
		NewPerson(4800, "Kael", 26)
	}
}