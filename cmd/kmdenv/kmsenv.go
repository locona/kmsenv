package main

import (
	"encoding/json"
	"os"

	"github.com/locona/kmsenv"
)

func main() {
	kmsenv, err := kmsenv.New()
	if err != nil {
		panic(err)
	}

	res, err := kmsenv.Encrypt()
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./kms-encrypt.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}
