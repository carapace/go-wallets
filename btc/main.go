package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response []UTXO

type UTXO struct {
	Address string					`json:"address"`
	TXID string  					`json:"txid"`
	VOut float64					`json:"vout"`
	ScriptPubKey string				`json:"scriptPubKey"`
	Amount float64					`json:"amount"`
	Satoshi float64					`json:"satoshi"`
	Height float64					`json:"height"`
	Confirmations float64			`json:"confirmations"`
}

func main() {
	response, err := http.Get("https://blockexplorer.com/api/addr/168bQ9CBHgn285MrZCxovxKaeXLNZiAHhd/utxo?noCache=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	balance := 0.0
	for i := 0; i < len(responseObject); i++ {
		balance += responseObject[i].Amount
	}
	fmt.Println(balance)
}