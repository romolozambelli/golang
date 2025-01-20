package main

import "fmt"

func weekday(number int) string {
	switch number {
	case 1:
		return "Monday"

	case 2:
		return "Tuesdeay"

	case 3:
		return "Wednesday"

	case 4:
		return "Thursday"

	case 5:
		return "Friday"

	default:
		return "Holiday"
	}
}

func weekday2(number int) string {
	switch {
	case number == 1:
		return "Monday"
	default:
		return "Holiday"
	}
}

func weekday3(number int) string {

	var weekday string

	switch {
	case number == 1:
		weekday = "Monday"
		fallthrough // Goes inside of the next condition - Does it make sense???

	case number == 2:
		weekday = "Tuesday"

	case number == 3:
		weekday = "Wednesday"

	case number == 4:
		weekday = "Thursday"

	case number == 5:
		weekday = "Friday"

	default:
		weekday = "Holiday"
	}

	return weekday
}

func main() {
	fmt.Println("Switch Functions")

	fmt.Println(weekday(2))

	fmt.Println(weekday2(2))

	fmt.Println(weekday3(1))

}
