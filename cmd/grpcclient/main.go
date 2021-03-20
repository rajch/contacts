package main

import (
	"fmt"
)

func main() {
mainloop:
	for {
		fmt.Println("Menu")
		fmt.Println("====")
		fmt.Printf("1. Add New Contact\n2. List Contacts\n3. Get Contact by ID\n4. Exit\n\n")
		fmt.Print("Your choice: ")
		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			addcontact()
		case 2:
			getallcontacts()
		case 3:
			getcontactbyid()
		case 4:
			break mainloop
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
