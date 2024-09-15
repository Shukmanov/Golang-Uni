package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Id   int
}

type Manager struct {
	Employee
	Department string
}

func (student Person) Stalking() string {
	return fmt.Sprintf("Im %s and  %v years old.", student.Name, student.Age)
}

func (student Person) OmedetoOrGreeting() string {
	return fmt.Sprintf("You are kyoufu no maaaad Saentistoo %s!", student.Name)
}

func (e Employee) Work() string {
	return fmt.Sprintf("%v %s", e.Id, e.Name)
}
func main() {
	// Ex 1
	fmt.Println("Hello, Worlds!")

	//Ex 2
	var chislo int = 90
	var stroka string = "T"
	var flout float32 = 0.029939828 //8322993
	var dauble float64 = 0.0299398288322993
	var boool bool = true

	dauble2 := 0.99398288322993
	flout2 := 0.993982883 //22993
	stroka2 := "Q"
	chislo2 := 94
	boool2 := true

	fmt.Printf("%v %v %v %v %v %v %v %v %t %t\n", chislo, stroka, flout, dauble, dauble2, flout2, stroka2, chislo2, boool, boool2)

	//ex 3
	resultIf := IFcheckNum(flout)
	fmt.Println(resultIf)

	fmt.Println(FORNaturalSum())

	day := "Monday"
	fmt.Println(SWITCHSchedule(day))

	//ex 4 function
	number1, number2 := 93882, 930848

	fmt.Println(SumOfTwo(number1, number2))

	fmt.Println(reversString(stroka, stroka2))

	fmt.Println(coeffTwoNum(number1, number2))

	// OOP ex_1
	kbtu21B031280 := Person{"Kuanysh", 20}

	fmt.Println(kbtu21B031280.Name, kbtu21B031280.Age)

	fmt.Println(kbtu21B031280.Stalking())

	fmt.Println(kbtu21B031280.OmedetoOrGreeting())

	kbtu20B019384 := Person{"Kyouma", 21}
	fmt.Println(kbtu20B019384.OmedetoOrGreeting())

	// OOP ex_2
	Darkhan := Employee{"Darkhan", 1}
	fmt.Println(Darkhan.Work())

	managerNUMBER1 := Manager{Employee{"Darkhan", 1}, "University"}
	fmt.Println(managerNUMBER1.Work())

}

// Ex 3 IF FOR SWITCH
func IFcheckNum(flout float32) string {
	if flout > 0 {
		return "Positive"

	} else if flout < 0 {
		return "Negative"

	} else {
		return "Zero"
	}
}

func FORNaturalSum() int {
	numSum := 0
	for i := 0; i < 10; i++ {
		numSum += i
	}
	return numSum
}

func SWITCHSchedule(classToday string) string {
	switch classToday {
	case "Monday":
		return "Class at 9am and 6pm"
	case "Tuesday":
		return "you can chill"
	case "Wednesday":
		return "Class at 9am"
	case "Thursday":
		return "Class at 10am to 1pm"
	case "Friday":
		return "Class at 8am to 5pm"
	case "Saturday":
		return "You can chill"
	case "Sunday":
		return "HOMEWORKS"
	default:
		return "Day?"
	}
}

// Ex 4 functions
func SumOfTwo(number1, number2 int) int {
	return number1 + number2
}

func reversString(letter1, letter2 string) (string, string) {
	return letter2, letter1
}

func coeffTwoNum(number1, number2 int) (int, int) {
	if number1 > number2 {
		return number1 / number2, number1 % number2
	} else if number1 < number2 {
		return number2 / number1, number2 % number1
	} else {
		return 1, 0
	}
}

