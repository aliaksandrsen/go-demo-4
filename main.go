package main

import (
	"fmt"

	"demo/app-4/account"
	"demo/app-4/files"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": CreateAccount,
	"2": FindAccount,
	"3": DeleteAccount,
}

func main() {
	vault := account.NewVault(files.NewJsonDB("data.json"))

	for {
		inputValue := promptData([]string{
			"",
			"Введите 1 для создания аккаунта",
			"Введите 2 для поиска аккаунта",
			"Введите 3 для удаления аккаунта",
			"Введите 4 для выхода",
			"Выберите вариант",
		})

		menuFunc := menu[inputValue]
		if menuFunc == nil {
			break
		}

		menuFunc(vault)
	}
}

func FindAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Bведите URL для поиска"})
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.Output()
	}
}

func DeleteAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Bведите URL для поиска"})
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Аккаунт удален")
	} else {
		color.Red("Аккаунт не найден")
	}
}

func CreateAccount(vault *account.VaultWithDB) {
	login := promptData([]string{"Веедите логин"})
	password := promptData([]string{"Веедите пароль (оставьте пустым для генерации)"})
	url := promptData([]string{"Веедите url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData[T any](list []T) string {
	for i, v := range list {
		if i == len(list)-1 {
			fmt.Printf("%v: ", v)
		} else {
			fmt.Println(v)
		}
	}

	var res string

	fmt.Scanln(&res)
	return res
}
