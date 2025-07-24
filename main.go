package main

import (
	"fmt"

	"demo/app-4/account"
	"demo/app-4/files"
)

func main() {
	files.ReadFile()
	files.WriteFile("Hello, World!", "file.txt")
	login := promptData("Веедите логин")
	password := promptData("Веедите пароль (оставьте пустым для генерации)")
	url := promptData("Веедите url")

	account1, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	account1.OutputPassword()

	fmt.Printf("account1: %v\n", account1)
}

func promptData(prompt string) string {
	var res string
	fmt.Println(prompt)

	fmt.Scanln(&res)
	return res
}
