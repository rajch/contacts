package main

import (
	"context"
	"fmt"
	"io"

	cgrpc "github.com/rajch/contacts/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func addcontact() {
	fmt.Println("Add new contact")
	fmt.Println("===============")
	var name, phone, email, city string
	var age int32

	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Email: ")
	fmt.Scanf("%s", &email)
	fmt.Print("Phone: ")
	fmt.Scanf("%s", &phone)
	fmt.Print("City: ")
	fmt.Scanf("%s", &city)
	fmt.Print("Age: ")
	fmt.Scanf("%d", &age)
	fmt.Println()

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not connect to contacts GRPC server: %v.", err)
	}
	defer conn.Close()

	client := cgrpc.NewContactServiceClient(conn)

	newcontact, err := client.NewContact(
		context.Background(),
		&cgrpc.Contact{
			Name:  name,
			Phone: phone,
			Email: email,
			City:  city,
			Age:   age,
		},
	)
	if err != nil {
		fmt.Printf("Could not add new contact via rpc: %v\n", err)
		return
	}

	showcontactheader()
	showcontact(newcontact)
}

func getallcontacts() {
	fmt.Println("List contacts")
	fmt.Println("=============")

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not connect to contacts GRPC server: %v.", err)
	}
	defer conn.Close()
	client := cgrpc.NewContactServiceClient(conn)

	allcontacts, err := client.GetAllContacts(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	showcontactheader()

	for {
		contact, err := allcontacts.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		showcontact(contact)
	}
}

func getcontactbyid() {
	fmt.Println("Get contact by id")
	fmt.Println("=================")
	var choice int32
	fmt.Print("Id: ")
	fmt.Scanf("%d", &choice)
	fmt.Println()

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not connect to contacts GRPC server: %v.", err)
	}
	defer conn.Close()

	client := cgrpc.NewContactServiceClient(conn)
	fetchedcontact, err := client.GetContactById(
		context.Background(),
		&cgrpc.GetContactInput{
			Id: choice,
		},
	)
	if err != nil {
		fmt.Printf("Could not get contact via rpc: %v\n", err)
		return
	}

	showcontactheader()
	showcontact(fetchedcontact)

}

func showcontactheader() {
	fmt.Printf(
		"%-3v %-20v %-10v %-10v %-10v %-3v\n",
		"Id",
		"Name",
		"Email",
		"Phone",
		"City",
		"Age",
	)
	fmt.Println("--- -------------------- ---------- ---------- ---------- ---")
}

func showcontact(contact *cgrpc.Contact) {
	fmt.Printf(
		"%-3v %-20.20v %-10.10v %-10.10v %-10.10v %-3d\n",
		contact.Id,
		contact.Name,
		contact.Email,
		contact.Phone,
		contact.City,
		contact.Age,
	)
}
