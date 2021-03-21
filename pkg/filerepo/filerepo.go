package filerepo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rajch/contacts/pkg/contact"
)

type Filerepo struct {
	filepath string
	LastId   uint
	Contacts []*contact.Contact
}

func New(filename string) (contact.Repository, error) {
	f, err := os.Open(filename)

	if os.IsNotExist(err) {
		return &Filerepo{
			filepath: filename,
			LastId:   0,
			Contacts: []*contact.Contact{},
		}, nil
	}

	if err != nil {
		return nil, err
	}

	defer f.Close()

	filedata, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var result Filerepo

	err = json.Unmarshal(filedata, &result)
	if err != nil {
		return nil, err
	}

	result.filepath = filename

	return &result, nil
}

func (f *Filerepo) Update(c *contact.Contact) (*contact.Contact, error) {
	record, err := f.Get(c.Id)
	if err != nil {
		return c, err
	}

	record.Name = c.Name
	record.Email = c.Email
	record.Phone = c.Phone
	record.City = c.City
	record.Age = c.Age

	return record, nil
}

func (f *Filerepo) Delete(c *contact.Contact) error {
	resultindex, _ := f.find(c.Id)
	if resultindex == -1 {
		return os.ErrNotExist
	}

	newlength := len(f.Contacts) - 1
	f.Contacts[resultindex] = f.Contacts[newlength]
	f.Contacts = f.Contacts[:newlength]

	return nil
}

func (f *Filerepo) Get(id uint) (*contact.Contact, error) {
	resultindex, result := f.find(id)
	if resultindex > -1 {
		return result, nil
	}

	return nil, os.ErrNotExist
}

func (f *Filerepo) find(id uint) (int, *contact.Contact) {
	for index, item := range f.Contacts {
		if item.Id == id {
			return index, item
		}
	}
	return -1, nil
}

func (f *Filerepo) List() ([]*contact.Contact, error) {
	return f.Contacts, nil
}

func (f *Filerepo) Close() {
	f.Flush()
}

func (f *Filerepo) Flush() error {
	frdata, err := json.Marshal(f)
	if err != nil {
		return err
	}

	fl, err := os.Create(f.filepath)
	if err != nil {
		return err
	}

	defer fl.Close()

	_, err = fl.Write(frdata)
	if err != nil {
		return err
	}

	return nil
}

func (f *Filerepo) New(c *contact.Contact) (*contact.Contact, error) {
	f.LastId += 1
	c.Id = f.LastId
	f.Contacts = append(f.Contacts, c)
	return c, f.Flush()
}
