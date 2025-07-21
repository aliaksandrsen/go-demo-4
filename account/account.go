package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Account struct {
	login    string
	password string
	url      string
}

// type accountWirWithTimeStamp struct {
// 	createdAt time.Time
// 	updatedAt time.Time
// 	account
// }

func (acc Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
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
		login:    login,
		password: password,
		url:      urlString,
	}
	if password == "" {
		acc.generatePassword(5)
	}

	return acc, nil
}
