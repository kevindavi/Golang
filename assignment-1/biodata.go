package main

import (
	"fmt"
	"os"
	"strconv"
)

// Friend represents a friend's data
type Friend struct {
	Name        string
	Address     string
	Job         string
	ClassReason string
}

// getFriendData returns a list of friends
func getFriendData() []Friend {
	return []Friend{
		{"Alice", "123 Pine St", "Software Engineer", "To learn Go."},
		{"Bob", "456 Maple Ave", "Data Scientist", "For career advancement."},
		{"Charlie", "789 Oak Blvd", "DevOps Specialist", "Interested in Go's simplicity."},
		// Add more entries as needed
	}
}

// displayFriendData prints friend data based on the given absen number
func displayFriendData(friends []Friend, absen int) {
	if absen > 0 && absen <= len(friends) {
		friend := friends[absen-1]
		fmt.Println("Name:", friend.Name)
		fmt.Println("Address:", friend.Address)
		fmt.Println("Job:", friend.Job)
		fmt.Println("Class Reason:", friend.ClassReason)
	} else {
		fmt.Println("No data found for the given number.")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an absen number as argument.")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number:", os.Args[1])
		return
	}

	friends := getFriendData()
	displayFriendData(friends, absen)
}
