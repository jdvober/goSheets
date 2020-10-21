/* package main
 *
 * import (
 *     auth "github.com/jdvober/goGoogleAuth"
 *     sh "github.com/jdvober/goSheets/values"
 * )
 *
 *
 * func main() {
 *     spreadsheetID := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
 *     r := "H4"
 *     // With range = "A1"
 *     vals := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}}
 *     client := auth.Authorize()
 *     sh.BatchUpdate(client, spreadsheetID, r, vals)
 *
 * } */

package main

import (
	"fmt"

	auth "github.com/jdvober/goGoogleAuth"
	sh "github.com/jdvober/goSheets/values"
)

type body struct {
	Data struct {
		Range  string     `json:"range"`
		Values [][]string `json:"values"`
	} `json:"data"`
	ValueInputOption string `json:"valueInputOption"`
}

func main() {
	client := auth.Authorize()

	spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
	rangeData := "sheet2!I1:J3"
	values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}}

	sh.BatchUpdate(client, spreadsheetId, rangeData, values)
	fmt.Println("Finished main()")
}
