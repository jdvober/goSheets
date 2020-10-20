package main

import (
	auth "github.com/jdvober/goGoogleAuth"
	sh "github.com/jdvober/goSheets/values"
)

var spreadsheetID string = "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
var r string = "H4"
var vals []interface{} = []interface{}{"One", "Two"}

func main() {
	client := auth.Authorize()
	sh.Update(client, spreadsheetID, r, vals)

}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
