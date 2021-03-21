package contact

type Repository interface {
	New(*Contact) (*Contact, error)
	Update(*Contact) (*Contact, error)
	Delete(*Contact) error

	Get(uint) (*Contact, error)
	List() ([]*Contact, error)

	Close()
}
