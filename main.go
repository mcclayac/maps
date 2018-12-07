package main

import "fmt"

func main() {

	dayMonths := make(map[string] int)  // map = dictionary
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in Feburary :%d\n", dayMonths["Feb"])
	//fmt.Printf("Days in Janurary :%d\n", dayMonths["Janurary"])  // new item in map

	january := "January"
	days, ok := dayMonths[january]
	if !ok {
		fmt.Printf("Can't get days for %s\n", january)
	} else {
		fmt.Printf("days: %d\n", days)
	}

	january = "Jan"
	days, ok = dayMonths[january]
	if !ok {
		fmt.Printf("Can't get days for %s\n", january)
	} else {
		fmt.Printf("days: %d\n", days)
	}

	for key, value := range dayMonths {  //  un-ordered
		fmt.Printf("%s has %d days\n", key, value)
	}
 	has31 := 0
	for _, days := range dayMonths {  //  don't care about the key
		if days == 31 {
			has31++
		}
	}
	fmt.Printf("%d months have 31 days\n", has31)

	fmt.Printf("1 - Number of Months in Year : %d\n", len(dayMonths))
	delete(dayMonths, "Feb")
	fmt.Printf("2 - Number of Months in Year : %d\n", len(dayMonths))
	delete(dayMonths, "Feb")
	fmt.Printf("3 - Number of Months in Year : %d\n", len(dayMonths))

	// static map
	fmt.Printf("Static Map\n")
	dayMonths2 := map[string] int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}
	for key, value := range dayMonths2 {  //  un-ordered
		fmt.Printf("Static: %s has %d days\n", key, value)
	}

}
