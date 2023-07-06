package tr

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Name     string
	Password string
	Next     *User
}

type Admins struct {
	Total int
	Datas *User
}

func NewPersonil(data User) *User {
	return &User{
		Id:       data.Id,
		Name:     data.Name,
		Password: data.Password,
	}
}

type Linkliest interface {
	Add(data User)
	GetAll() ([]User, error)
}

func (a *Admins) Add(data User) {

	a.Total += 1
	data.Id = a.Total
	dat := Bcript(data.Password)
	data.Password = dat

	if a.Datas == nil {
		a.Datas = &data
		return
	}
	br := a.Datas

	for br.Next != nil {
		br = br.Next
	}
	if br.Next == nil {
		br.Next = &data
	}

}
func Bcript(y string) string {
	password := []byte(y)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)

}
func (a *Admins) GetAll() ([]User, error) {
	var users []User

	br := a.Datas

	for br != nil {
		users = append(users, *br)
		br = br.Next
	}

	// var dat = Modell(datas)
	return users, nil

}
func main() {
	users := User{
		Id:       1,
		Name:     "fajar",
		Password: "fajaraj",
	}

	users2 := User{
		Id:       1,
		Name:     "fajar",
		Password: "fajaraj",
	}
	var admin Admins
	var dats Linkliest = &admin
	dats.Add(users)
	dats.Add(users2)

	data, _ := dats.GetAll()
	fmt.Print(data)

}
