package main

import (
	"fmt"

	"demo/app-4/account"
)

func getMenu() {
	fmt.Println("Введите 1 для создания аккаунта")
	fmt.Println("Введите 2 для поиска аккаунта")
	fmt.Println("Введите 3 для удаления аккаунта")
	fmt.Println("Введите 4 для выхода")
	fmt.Println("")
}

func main() {
	for {
		getMenu()

		var inputValue int
		fmt.Scanln(&inputValue)

		if inputValue == 1 {
			CreateAccount()
		}
		if inputValue == 2 {
			FindAccount()
		}
		if inputValue == 3 {
			DeleteAccount()
		}
		if inputValue == 4 {
			break
		}
	}
}

func FindAccount() {
	fmt.Println("Аккаунт найден")
	fmt.Println("")
}

func DeleteAccount() {
	fmt.Println("Аккаунт удален")
	fmt.Println("")
}

func CreateAccount() {
	login := promptData("Веедите логин")
	password := promptData("Веедите пароль (оставьте пустым для генерации)")
	url := promptData("Веедите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	var res string
	fmt.Println(prompt)

	fmt.Scanln(&res)
	return res
}
