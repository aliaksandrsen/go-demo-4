package main

import (
	"fmt"
	"strings"

	"demo/app-4/account"
	"demo/app-4/files"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDB("data.json"))

	for {
		inputValue := promptData(
			"",
			"Введите 1 для создания аккаунта",
			"Введите 2 для поиска аккаунта по URL",
			"Введите 3 для поиска аккаунта по имени",
			"Введите 4 для удаления аккаунта",
			"Введите 5 для выхода",
			"Выберите вариант",
		)

		menuFunc := menu[inputValue]
		if menuFunc == nil {
			break
		}

		menuFunc(vault)
	}
}

var menu = map[string]func(*account.VaultWithDB){
	"1": CreateAccount,
	"2": FindAccount("URL"),
	"3": FindAccount("login"),
	"4": DeleteAccount,
}

var checkMap = map[string]func(*account.Account, string) bool{
	"URL": func(account *account.Account, str string) bool {
		return strings.Contains(account.Url, str)
	},
	"login": func(account *account.Account, str string) bool {
		return strings.Contains(account.Login, str)
	},
}

func FindAccount(value string) func(*account.VaultWithDB) {
	return func(vault *account.VaultWithDB) {
		str := promptData(fmt.Sprintf("Введите %s для поиска", value))
		accounts := vault.FindAccounts(str, checkMap[value])

		if len(accounts) == 0 {
			color.Red("Аккаунтов не найдено")
		}

		for _, account := range accounts {
			account.Output()
		}
	}
}

func DeleteAccount(vault *account.VaultWithDB) {
	url := promptData("Bведите URL для поиска")
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Аккаунт удален")
	} else {
		color.Red("Аккаунт не найден")
	}
}

func CreateAccount(vault *account.VaultWithDB) {
	login := promptData("Веедите логин")
	password := promptData("Веедите пароль (оставьте пустым для генерации)")
	url := promptData("Веедите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData(list ...string) string {
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
