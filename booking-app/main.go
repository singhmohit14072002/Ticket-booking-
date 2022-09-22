package main

import (
	"fmt"
	"sync"
	"time"

	"strings"
)
 
var conferenceName = "Go Conference" // variable define in function
const conferenceTicket int = 50
var remainingTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
  firstName string
  lastName string
  email   string
  numberOfTickets uint
  
}

var wg = sync.WaitGroup{}

func main() {
	 
	 greetUsers()
	
	      firstName, lastName, email, userTicket :=  getUserInput()
		  isValidName,isValidEmail,isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)

           if isValidName && isValidEmail && isValidTicketNumber {
				
			bookTicket( userTicket , firstName , lastName  , email)
            
			wg.Add(1)
			go sendTicket( userTicket , firstName , lastName  , email)
           

		            // call function print name 
                    firstName := getFirstName() // here booking is a parameter 
				     fmt.Printf("The first name of the bookings are : %v\n" ,firstName)
					
		             if  remainingTicket == 0 {
			          // end the program
			        fmt.Printf("Our conference is booked out. Come back next year.")
					 //  break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is to short ")
			} 
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			 }
			  if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			 }
		}
	      wg.Wait() // it wait untile the application are excute 
}

// simply we replace print line with the function call by using of greetUsers 
func greetUsers() {
  fmt.Printf("Welcome to %v booking application\n" , conferenceName)
  fmt.Printf("We have total of %v tickets and %v are still available .\n" , conferenceTicket, remainingTicket)
  fmt.Println("Get your tickets here to attend")
}

func getFirstName()[]string{
	firstNames := []string{}
					for _, booking := range bookings {
				firstNames = append( firstNames,  booking.firstName)
		}
		return firstNames
}

 func validateUserInput(firstName string , lastName string , email string, userTicket uint)  (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail :=  strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

     return isValidName,isValidEmail,isValidTicketNumber
 }

 func getUserInput() (string, string,string,uint) {
	var firstName string
	var lastName  string
	var email string 
	var userTicket uint 
// ask user  for their name 
	 fmt.Println("Enter your first name: ") 
	fmt.Scan(&firstName)

   fmt.Println("Enter your last name: ") 
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ") 
	fmt.Scan(&email)

	  fmt.Println("Enter your no of ticket: ") 
	 fmt.Scan(&userTicket)

	 return firstName, lastName, email, userTicket
 }

 func bookTicket( userTicket uint, firstName string , lastName string , email string ) {
 remainingTicket  = remainingTicket - userTicket
	
// create a map for a user data
var userData =  UserData {
	firstName: firstName,
	lastName: lastName,
	email: email,
	numberOfTickets: userTicket,
}

bookings = append(bookings, userData ) // now here we working dynamically using slice 
	  fmt.Printf("List of booking is %v\n " , bookings)
// fmt.Println(&remainingTicket) // it will print the memory location
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n" , firstName,lastName, userTicket,email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket,conferenceName )

 }

 func sendTicket(userTicket uint, firstName string,lastName string , email string)  {
	time.Sleep(30 * time.Second)
      var ticket = fmt.Sprintf("%v tickets for %v %v,", userTicket, firstName, lastName)
    fmt.Println("########")
	  fmt.Printf("Sending ticket:\n %v to email %v\n", ticket ,email)   
    fmt.Println("########")
	wg.Done()
 }