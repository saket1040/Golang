package main

import "fmt"

type User struct {
	Name  string
	Email string
	Phone string
	Age   int16
}

type UserBuilder struct {
	// Right now, UserBuilder.User is exported, which breaks encapsulation.
	// User *User
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{},
	}
}

func (u *UserBuilder) Name(name string) *UserBuilder {
	u.user.Name = name
	return u
}

func (u *UserBuilder) Email(email string) *UserBuilder {
	u.user.Email = email
	return u
}

func (u *UserBuilder) Phone(phone string) *UserBuilder {
	u.user.Phone = phone
	return u
}

func (u *UserBuilder) Age(age int16) *UserBuilder {
	u.user.Age = age
	return u
}

func (u *UserBuilder) Build() *User {
	return u.user
}

// better way to print
func (u *User) String() string {
	return fmt.Sprintf("User{Name: %s, Email: %s, Phone: %s, Age: %d}", u.Name, u.Email, u.Phone, u.Age)
}

func main() {
	fmt.Println("Hello, World!")
	builder := NewUserBuilder()
	user := builder.Name("Alice").Email("alice@example.com").Phone("12345").Build()
	fmt.Println(user)
}