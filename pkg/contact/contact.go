package contact

type Contact struct {
	Id    uint `json:"id" gorm:"primary_key"`
	Name  string
	Phone string
	Email string
	City  string
	Age   int
}
