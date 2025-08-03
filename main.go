package main

import (
	"fmt"

	"demo/app-4/account"

	"github.com/fatih/color"
)

func getMenu() {
	fmt.Println("Введите 1 для создания аккаунта")
	fmt.Println("Введите 2 для поиска аккаунта")
	fmt.Println("Введите 3 для удаления аккаунта")
	fmt.Println("Введите 4 для выхода")
	fmt.Println("")
}

func main() {
	vault := account.NewVault()

	for {
		getMenu()

		var inputValue int
		fmt.Scanln(&inputValue)

		if inputValue == 1 {
			CreateAccount(vault)
		}
		if inputValue == 2 {
			FindAccount(vault)
		}
		if inputValue == 3 {
			DeleteAccount(vault)
		}
		if inputValue == 4 {
			break
		}
	}
}

func FindAccount(vault *account.Vault) {
	url := promptData("Bведите URL для поиска")
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.Output()
	}
}

func DeleteAccount(vault *account.Vault) {
	url := promptData("Bведите URL для поиска")
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Аккаунт удален")
	} else {
		color.Red("Аккаунт не найден")
	}
}

func CreateAccount(vault *account.Vault) {
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

func promptData(prompt string) string {
	var res string
	fmt.Println(prompt)

	fmt.Scanln(&res)
	return res
}
