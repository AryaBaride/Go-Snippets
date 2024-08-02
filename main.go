package main

import "fmt"

var total_tickets int = 100
var remaining_tickets int = total_tickets
var review_list []string

func main() {
	fmt.Println("Welcome to ticket booking application")
	fmt.Printf("We have %d tickets available:\n", total_tickets)

	for {
		fmt.Println("Select one of the options")
		fmt.Println("1. Buy a ticket")
		fmt.Println("2. Cancel ticket")
		fmt.Println("3. Post a review")
		fmt.Println("4. Check Reviews")
		fmt.Println("5. Exit")

		var choice int
		fmt.Println("Enter your choice : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			booking_logic()
		case 2:
			cancel_tickets()
		case 3:
			post_review()
		case 4:
			check_reviews()
		case 5:
			fmt.Println("Exited the application, thank you")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func booking_logic() {
	var tickets_sold int
	fmt.Println("Enter the number of tickets you want to book :")
	fmt.Scanln(&tickets_sold)
	if tickets_sold > remaining_tickets {
		fmt.Println("Not enough tickets available.")
		return
	}
	remaining_tickets = total_tickets - tickets_sold
	fmt.Printf("Thank you for booking %d tickets\n", tickets_sold)
	fmt.Printf("Remaining tickets: %d\n", remaining_tickets)
}

func cancel_tickets() {
	var cancelled_tickets int
	fmt.Println("Enter the number of tickets to cancel")
	fmt.Scanln(&cancelled_tickets)
	if cancelled_tickets > (total_tickets - remaining_tickets) {
		fmt.Println("You cannot cancel more tickets than you have booked.")
		return
	}
	new_remaining_tickets := remaining_tickets + cancelled_tickets
	fmt.Println("You have successfully cancelled your ticket")
	fmt.Printf("%d tickets remaining\n", new_remaining_tickets)
}

func check_reviews() {
	fmt.Println("Here are some of the reviews from our customers:")
	for _, value := range review_list {
		fmt.Println(value)
	}
}

func post_review() {
	fmt.Println("Enter your review")
	var review string
	fmt.Scanln(&review)
	review_list = append(review_list, review)
	fmt.Println("Thank you for posting a review!")
}
