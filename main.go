package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T , conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("we have total %v tickets and %v  remainingTickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string
	var FirstName string
	var LastName string
	var Email string
	var userTicket uint

	fmt.Println("Enter your First Name")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&LastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&Email)

	fmt.Println("How Many Tickets do you want")
	fmt.Scan(&userTicket)

	// ask for user input

	remainingTickets = remainingTickets - userTicket
	bookings[0] = FirstName + " " + LastName

	fmt.Printf("The Whole Array is : %v\n", bookings)
	fmt.Printf("The First Value is : %v\n", bookings[0])
	fmt.Printf("Array type  : %T\n", bookings)
	fmt.Printf("Array length : %v\n", len(bookings))

	fmt.Printf("Thank You %v %v for booking %v tickets You Will receive a confirmation email at %v\n", FirstName, LastName, userTicket, Email)
	fmt.Printf("Now %v ticket are left out", remainingTickets)

}
