package model

type User struct {
	name   string
	flight Flight
}

func NewUser(name string, flight Flight) *User {
	return &User{name, flight}
}

func (u *User) Name() (string, error) {
	return u.name, nil
}

func (u *User) Flight() (string, error) {
	return u.name, nil
}
