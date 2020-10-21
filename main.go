package main

import (
	"fmt"

	auth "github.com/jdvober/goGoogleAuth"
	sh "github.com/jdvober/goSheets/values"
)

func main() {
	client := auth.Authorize()

	spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
	rangeData := "sheet2!I1:J3"
	values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}}

	sh.BatchUpdate(client, spreadsheetId, rangeData, values)
	fmt.Println("Finished main()")
}
