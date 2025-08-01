package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"ceatedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login cannot be empty")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	acc := &Account{
		Login:    login,
		Password: password,
		Url:      urlString,
	}

	// field, _ := reflect.TypeOf(acc).Elem().FieldByName("login")
	// fmt.Println(string(field.Tag))

	if password == "" {
		acc.generatePassword(5)
	}

	return acc, nil
}
