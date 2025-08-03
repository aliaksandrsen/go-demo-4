package account

import (
	"encoding/json"
	"strings"
	"time"

	"demo/app-4/files"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db files.JsonDB
}

func NewVault(db *files.JsonDB) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}
	var vault Vault

	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
	}

	return &VaultWithDB{
		Vault: vault,
		db:    *db,
	}
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}

	vault.db.Write(data)
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDB) FindAccountsByURL(url string) []Account {
	var result []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)

		if isMatched {
			result = append(result, account)
		}
	}

	return result
}

func (vault *VaultWithDB) DeleteAccountsByURL(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)

		if !isMatched {
			accounts = append(accounts, account)
			continue
		}

		isDeleted = true
	}

	vault.Accounts = accounts
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}

	vault.db.Write(data)

	return isDeleted
}
