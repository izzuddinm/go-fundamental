package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months)

	sliceOne := months[3:5]
	fmt.Println(sliceOne)

	sliceTwo := months[:7]
	fmt.Println(sliceTwo)

	sliceThree := months[7:]
	fmt.Println(sliceThree)

	sliceFour := months[:]
	fmt.Println(sliceFour)

	var sliceFive []string = months[:]
	fmt.Println(sliceFive)

	// function slice
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice := days[5:]
	daySlice[0] = "New Saturday"
	daySlice[1] = "New Sunday"

	fmt.Println(days)
	fmt.Println(daySlice)

	daySlice = append(daySlice, "New Day")
	fmt.Println(daySlice)

	lenDaySlice := len(daySlice)
	fmt.Println(lenDaySlice)

	capDaySlice := cap(daySlice)
	fmt.Println(capDaySlice)

	makeSlice := make([]string, 3, 6)
	makeSlice[0] = "Java"
	makeSlice[1] = "Golang"
	makeSlice[2] = "Ruby On Rails"
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	fmt.Println(makeSlice)

	// makeSlice[3] = "Kotlin" will return error, must use append to add data
	makeSlice = append(makeSlice, "Kotlin")
	makeSlice = append(makeSlice, "JavaScript")
	makeSlice = append(makeSlice, "PHP")
	makeSlice = append(makeSlice, "Rust")
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	fmt.Println(makeSlice)

	fromSlice := days
	toSlice := make([]string, len(days), cap(days))
	fmt.Println("var fromSlice Before Copy : ", fromSlice)
	fmt.Println("var toSlice Before Copy : ", toSlice)
	copy(toSlice, fromSlice[:])
	fmt.Println("var fromSlice After Copy : ", fromSlice)
	fmt.Println("var toSlice After Copy : ", toSlice)

	// what a different array between slice
	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
