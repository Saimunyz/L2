package pattern

import (
	"errors"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Email structure
type Email struct {
	email string
}

// Check - checking the correctness of the incoming login email
func (e *Email) Check(incomingEmail string) bool {
	if e.email != incomingEmail {
		return false
	}

	return true
}

// newEmail - creates new Email instance
func newEmail(email string) *Email {
	return &Email{
		email: email,
	}
}

// Password structure
type Password struct {
	password string
}

// Check - checking the correctness of the incoming login password
func (p *Password) Check(incomingPassword string) bool {
	if p.password != incomingPassword {
		return false
	}

	return true
}

// newPassword - creates new Password instance
func newPassword(password string) *Password {
	return &Password{
		password: password,
	}
}

// SecuritiCode structure
type SecurityCode struct {
	code int
}

// newSecurityCide - creates new SecurityCode instance
func newSecurityCode() *SecurityCode {
	return &SecurityCode{
		code: 0,
	}
}

// SendCode - returns random generated code for security check
func (s *SecurityCode) SendCode() int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Int()
	s.code = code
	return code
}

// Check - checking the correctness of the incoming login code
func (s *SecurityCode) Check(incomingCode int) bool {
	if s.code != incomingCode {
		return false
	}
	return true
}

// User - facade for login session
type User struct {
	email    *Email
	password *Password
	security *SecurityCode
}

// newUser - creates new user instance
func newUser(email, password string) *User {
	return &User{
		email:    newEmail(email),
		password: newPassword(password),
		security: newSecurityCode(),
	}
}

// Login - facade function that hides all internal processes
func (u *User) Login(email, password string) error {
	if u.email.Check(email) && u.password.Check(password) {
		code := u.security.SendCode()
		if !u.security.Check(code) {
			return errors.New("Wrong security code")
		}

		return nil
	}

	return errors.New("wrong email or password")
}
