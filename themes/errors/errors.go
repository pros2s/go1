package errors

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
)

type Animal struct {
	animalType string
	age        int
}

func NewAnimal(animalType string, age int) *Animal {
	return &Animal{
		animalType: animalType,
		age:        age,
	}
}

func (a *Animal) ChangeAge(age int) (int, error) {
	if age < 0 {
		return age, errors.New("incorrect age: less than zero")
	}

	if age > 250 {
		return age, errors.New("incorrect age: more than 250")
	}

	a.age = age

	return age, nil
}

func AnimalErrorsTest() {
	giraffe := NewAnimal("giraffe", 150)
	pp.Println(giraffe)

	elephant := NewAnimal("elephant", 25)

	for range 5 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter age of elephant:")
		scanner.Scan()

		scannedAge, err := strconv.Atoi(scanner.Text())
		if err != nil {
			err = errors.New("invalid age")
			pp.Println(err.Error())
			break
		}

		elephantAge, elephantErr := elephant.ChangeAge(scannedAge)

		if elephantErr != nil {
			pp.Println(elephantErr.Error())
		} else {
			pp.Println(elephantAge)
		}
	}
}
