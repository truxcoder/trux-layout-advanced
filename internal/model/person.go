package model

type Person struct {
	Base
}

func (m *Person) TableName() string {
	return "person"
}
