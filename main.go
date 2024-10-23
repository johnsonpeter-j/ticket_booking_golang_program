package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// array
// var bookings [50]string

// slice
// var bookings = []string{}
// var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isEmailValid, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isEmailValid && isValidTicketNumber {
			bookTicket(userTicket, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("These are all our bookings: %v\n.", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Printf("%v ticket is booked out. come back next year.\n", conferenceName)
				break
			}
		} else {
			// fmt.Printf("We only have %v ticket remaining, so you can't book %v tickets.\n", remainingTickets, userTicket)
			// continue
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}

			if !isEmailValid {
				fmt.Println("email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}
	wg.Wait()
}

func greetUser() {

	fmt.Printf("Welcome to %v ticket booking app.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	println("Enter your first name : ")
	fmt.Scan(&firstName)

	println("Enter your last name : ")
	fmt.Scan(&lastName)

	println("Enter your email name : ")
	fmt.Scan(&email)

	println("Enter number of tickets : ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets -= userTicket
	// for array
	// bookings[0] = firstName + " " + lastName

	// create a map for user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	// for slice
	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v.", bookings)

	// fmt.Printf("The whole array : %v\n", bookings)
	// fmt.Printf("The first value : %v\n", bookings[0])
	// fmt.Printf("Array Type : %T\n", bookings)
	// fmt.Printf("Array Length : %v\n", len(bookings))

	// fmt.Printf("The whole Slice : %v\n", bookings)
	// fmt.Printf("The first value : %v\n", bookings[0])
	// fmt.Printf("Slice Type : %T\n", bookings)
	// fmt.Printf("Slice Length : %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	wg.Done()

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTicket, firstName, lastName)

	fmt.Println("###################")
	fmt.Printf("Sending ticket :\n%v \nto email address \n%v", ticket, email)
	fmt.Println("###################")
}
