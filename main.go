package main

import (
	"fmt"

	"demo/app-4/account"
	"demo/app-4/files"
)

func main() {
	CreateAccaunt()
}

func CreateAccaunt() {
	login := promptData("Веедите логин")
	password := promptData("Веедите пароль (оставьте пустым для генерации)")
	url := promptData("Веедите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в json")
		return
	}

	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	var res string
	fmt.Println(prompt)

	fmt.Scanln(&res)
	return res
}
