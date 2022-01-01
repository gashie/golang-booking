package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)
var bookings = make([]UserData, 0)
var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}
var wg = sync.WaitGroup{}
func main() {
	

	
	greetUsers(conferenceName, conferenceTickets, remainingTickets)


		//ask user for their name

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			//Logic for remaining tickets
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1) //add new thread to wait group
		go	sendTicket(userTickets, firstName, lastName, email) //go word start new thread
			//create map for a user
			// var userData = make(map[string]string)
			// userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["email"] = email
	
		
		//if a function will return something, assign it to a variable
			firstNames := getFirstName()
 
			fmt.Printf("The first names of bookings are %v\n", firstNames)
			fmt.Printf("These are all our bookings: %v\n", bookings)

			//conditional statement
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year")
				//break //stop execution of the code
			}
		} else {
			if !isValidName {
				fmt.Printf("Your Name is not valid \n")
			}
			if !isValidEmail {
				fmt.Printf("Your email is not valid \n")
			}

			if !isValidTicketNumber {
				fmt.Printf("ticket number is not valid \n")
			}

			//fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets \n", remainingTickets, userTickets)
			//continue //skip to next iteration of the code
		}
   wg.Wait() //wait for all threads to complete   
	
}

func greetUsers(confName string, confTicket uint, remTicket uint) {
	fmt.Printf("Welcome to %v booking app", confName)
	fmt.Printf("We have total of %v tickets and %v are still available", confTicket, remTicket)
	fmt.Println("Get your tickets from here to attend")
}
func getFirstName() []string {
	firstNames := []string{}           //slice
	for _, booking := range bookings { //foreach
		//split strings
		//var names = strings.Fields(booking)
		
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets  for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() //removes the thread from waiting list
}
