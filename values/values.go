package values

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/sheets/v4"
)

type Request struct {
	SpreadsheetID string
	Range         string
	RequestBody   struct {
		Range          string
		MajorDimension string
		Values         [][]string
	}
}
type body struct {
	Data struct {
		Range          string     `json:"range"`
		Values         [][]string `json:"values"`
		MajorDimension string     `json:"majorDimension"`
	} `json:"data"`
	ValueInputOption string `json:"valueInputOption"`
}

// Update adds values to a Google Sheet
func Update(client *http.Client, spreadsheetID string, r string, dimension string, values []interface{}) {

	/* ctx := context.Background() */

	// Authenticate
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create sheets Client %v", err)
	}

	var vr sheets.ValueRange
	/* values := []interface{}{"One", "Two"} */
	vr.Values = append(vr.Values, values)
	vr.MajorDimension = dimension
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, r, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to post data to sheet. %v", err)
	}
}

// https://stackoverflow.com/questions/46230624/google-sheets-api-golang-batchupdatevaluesrequest
func BatchUpdate(client *http.Client, spreadsheetID string, rangeData string, majorDimension string, values [][]interface{}) {
	/* client := auth.Authorize() */
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	/* spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk" */
	/* rangeData := "sheet2!I1:J3" */
	/* values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}} */
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:          rangeData,
		Values:         values,
		MajorDimension: majorDimension,
	})
	_, err = sheetsService.Spreadsheets.Values.BatchUpdate(spreadsheetID, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")
}

func Get(client *http.Client, spreadsheetID string, readRange string) [][]interface{} {

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	}
	return resp.Values

}
