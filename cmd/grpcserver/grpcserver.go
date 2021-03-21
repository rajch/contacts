package main

import (
	"context"

	"github.com/rajch/contacts/pkg/contact"
	// "github.com/rajch/contacts/pkg/filerepo"
	"github.com/rajch/contacts/pkg/gormrepo"
	"github.com/rajch/contacts/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type contactServer struct {
	grpc.UnimplementedContactServiceServer
}

func (cs *contactServer) NewContact(_ context.Context, c *grpc.Contact) (*grpc.Contact, error) {
	newrecord := contact.Contact{
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
		City:  c.City,
		Age:   int(c.Age),
	}
	g, err := getrepo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer g.Close()

	newcontact, err := g.New(&newrecord)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpc.Contact{
		Id:    int32(newcontact.Id),
		Name:  newcontact.Name,
		Email: newcontact.Email,
		Phone: newcontact.Phone,
		City:  newcontact.City,
		Age:   int32(newcontact.Age),
	}, nil

}

func (cs *contactServer) GetAllContacts(_ *emptypb.Empty, stream grpc.ContactService_GetAllContactsServer) error {
	g, err := getrepo()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	defer g.Close()

	allcontacts, err := g.List()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, c := range allcontacts {
		err := stream.Send(&grpc.Contact{
			Id:    int32(c.Id),
			Name:  c.Name,
			Email: c.Email,
			Phone: c.Phone,
			City:  c.City,
			Age:   int32(c.Age),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs *contactServer) GetContactById(_ context.Context, ci *grpc.GetContactInput) (*grpc.Contact, error) {
	g, err := getrepo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer g.Close()

	contact, err := g.Get(uint(ci.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &grpc.Contact{
		Id:    int32(contact.Id),
		Name:  contact.Name,
		Email: contact.Email,
		Phone: contact.Phone,
		City:  contact.City,
		Age:   int32(contact.Age),
	}, nil
}

func getrepo() (contact.Repository, error) {
	//return filerepo.New("testdb.db.json")
	return gormrepo.New("testdb.db")
}
