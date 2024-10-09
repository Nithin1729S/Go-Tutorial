package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName string ="Go Conference"
	const conferenceTickets int =50  //can predict datatype 
	var remainingTickets uint =50;

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)
	fmt.Printf("Welcome to %v booking application",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remainingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")
	var bookings[]string
	for{
		//var username  //cant predict datatpe
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

		if userTickets <= remainingTickets{
			remainingTickets-=userTickets
			bookings=append(bookings,firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email);
			fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

			firstNames:=[]string{}
			for _,booking := range bookings{
				var names=strings.Fields(booking)
				var firstName = names[0]
				firstNames=append(firstNames, firstName)
			}

			fmt.Printf("These are our bookings: %v\n",firstNames)

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("SOLD OUT")
				break
			}
		} else{
			fmt.Printf("We only have %v tickets remaining , so you cant book %v tickets.\n",remainingTickets,userTickets)

		}

		
		
	}
 
	
}
