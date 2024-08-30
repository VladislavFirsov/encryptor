package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gitlab-private.wildberries.ru/wbpay-go/packages/crypto"
)

func main() {
	usage := `
Программа для шифрования строки "health-check" с использованием пакета crypto.
Использование:
  primary <путь к файлу с primary ключом>
  secondary <путь к файлу с secondary ключом>

Пример:
  ./encryptor /path/to/primary_key /path/to/secondary_key
`

	fmt.Println(usage)

	if len(os.Args) < 3 {
		log.Fatal("Необходимо указать оба пути: к primary и secondary ключам.")
	}

	primaryKeyPath := os.Args[1]
	secondaryKeyPath := os.Args[2]

	primaryKey, err := os.ReadFile(primaryKeyPath)
	if err != nil {
		log.Fatal("Ошибка чтения primary key: ", err)
	}

	secondaryKey, err := os.ReadFile(secondaryKeyPath)
	if err != nil {
		log.Fatal("Ошибка чтения secondary key: ", err)
	}

	cryptor, err := crypto.NewCryptorAES(primaryKey, secondaryKey, nil)
	if err != nil {
		log.Fatal("Ошибка создания CryptorAES: ", err)
	}

	dataToEncrypt := "health-check"

	encryptedBase64, err := cryptor.EncodeB64(context.Background(), dataToEncrypt)
	if err != nil {
		log.Fatal("Ошибка шифрования: ", err)
	}

	log.Println("Зашифрованные данные (Base64):", encryptedBase64)
}
