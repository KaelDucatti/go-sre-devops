package section01

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPerson(t *testing.T) {
	t.Run("Success Cases", func (t *testing.T) {
		t.Run("should return Person instance", func(t *testing.T) {
			require := require.New(t)

			p, err := NewPerson(4800, "Kael", 26)

			require.NoError(err)
			require.Equal(float32(4800), p.Wage())
			require.Equal("Kael", p.Name())
			require.Equal(26, p.Age())
		})
	})
	t.Run("Validation Errors", func(t *testing.T) {
		t.Run("should return wage error", func(t *testing.T) {
			require := require.New(t)

			_, err := NewPerson(-1, "Kael", 26)

			require.Error(err)
			require.EqualError(err, "Wage cannot be less than 0")
		})
		t.Run("should return name error", func(t *testing.T) {
			require := require.New(t)

			_, err := NewPerson(4800, "", 26)

			require.Error(err)
			require.EqualError(err, "Name cannot be void")
		})
		t.Run("should return age error", func(t *testing.T) {
			require := require.New(t)

			testCases := []struct {
				testName	string
				wage		float32
				name		string
				age			int
			}{
				{"age is negative", 4800, "Kael", -1},
				{"age is zero", 4800, "Kael", -0},
			}

			for _, test := range testCases {
				t.Run(test.testName, func(t *testing.T) {
					_, err := NewPerson(test.wage, test.name, test.age)

					require.Error(err)
					require.EqualError(err, "Age must be greater than 0")
				})
			}
		})
	})
}

func TestSetters(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		t.Run("should update wage and return nil", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)
			
			err := p.SetWage(5200)

			require.NoError(err)
			require.Equal(float32(5200), p.Wage())
		})
		t.Run("should update name and return nil", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)
			
			err := p.SetName("Mikael")
	
			require.NoError(err)
			require.Equal("Mikael", p.Name())
		})
		t.Run("should update age and return nil", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)

			err := p.SetAge(27)

			require.NoError(err)
			require.Equal(int(27), p.Age())
		})
	})
	t.Run("Error Validation", func(t *testing.T) {
		t.Run("should return wage error", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)

			err := p.SetWage(-1)

			require.Error(err)
			require.EqualError(err, "Wage cannot be less than 0")
			require.Equal(float32(4800), p.Wage())
		})
		t.Run("should return name error", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)

			err := p.SetName("")

			require.Error(err)
			require.EqualError(err, "Name cannot be void")
			require.Equal("Kael", p.Name())
		})
		t.Run("should return age error", func(t *testing.T) {
			require := require.New(t)
			p, _ := NewPerson(4800, "Kael", 26)

			testCases := []struct {
				testName 	string
				Age			int
			}{
				{"age is zero", 0},
				{"age is negative", -1},
			}

			for _, test := range testCases {
				t.Run(test.testName, func(t *testing.T) {
					err := p.SetAge(test.Age)
					require.Error(err)
					require.EqualError(err, "Age must be greater than 0")
					require.Equal(int(26), p.Age())
				})
			}
		})
	})
}

func TestGetters(t *testing.T) {
	require := require.New(t)
	p, _ := NewPerson(4800, "Kael", 26)

	require.Equal(p.Wage(), float32(4800))
	require.Equal(p.Name(), "Kael")
	require.Equal(p.Age(), 26)
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