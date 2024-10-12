package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName string ="Go Conference"
const conferenceTickets int =50  //can predict datatype 
var remainingTickets int =50;
var bookings= make([]UserData,0)


type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

func main(){
	
	greetUsers()

	for{
		//var username  //cant predict datatpe
		firstName,lastName,email,userTickets:=getUserInput()

		isValidName,isValidEmail,isValidTicketNumber:=helper.ValidateUserInput(firstName,lastName,email,int(userTickets),remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(int(userTickets),firstName,lastName,email)

			fmt.Printf("The first names of bookings are: %v\n",printFirstName()) 

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("SOLD OUT")
				break
			}
		} else{
			if !isValidName{
				fmt.Println("First name or last you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Your email is invalid")
			}
			if !isValidTicketNumber{
				fmt.Println("Invalid ticket number")
			}
			fmt.Println("Your input is invalid")

		}

		
		
	}
 
	
}


func greetUsers(){
	fmt.Printf("Welcome to %v  Conference\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remainingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")
}

func printFirstName() []string{
	firstNames:=[]string{}
	for _,booking := range bookings{
		var firstName = booking.firstName
		firstNames=append(firstNames, firstName)
	}
	return firstNames
}


func getUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets needed:")
	fmt.Scan(&userTickets)
	return firstName,lastName,email,userTickets
}

func bookTickets(userTickets int,firstName string,lastName string,email string){
	remainingTickets-=userTickets
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n",bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email);
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
}