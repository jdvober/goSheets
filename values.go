package gsheets

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/sheets/v4"
)

// Update adds values to a Google Sheet
func UpdateValues(client *http.Client, spreadsheetID string, r string, dimension string, values []interface{}) {
	// Authenticate
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create sheets Client %v", err)
	}

	var vr sheets.ValueRange
	/* values := []interface{}{"One", "Two"} */
	vr.Values = append(vr.Values, values)
	vr.MajorDimension = dimension
	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, r, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to post data to sheet. %v", err)
	}
}

// https://stackoverflow.com/questions/46230624/google-sheets-api-golang-batchupdatevaluesrequest
func BatchUpdateValues(client *http.Client, spreadsheetID string, rangeData string, majorDimension string, values [][]interface{}) {
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:          rangeData,
		Values:         values,
		MajorDimension: majorDimension,
	})
	_, err = srv.Spreadsheets.Values.BatchUpdate(spreadsheetID, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func GetValues(client *http.Client, spreadsheetID string, readRange string) [][]interface{} {
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

func Clear(client *http.Client, spreadsheetID string, r string) {
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	rb := &sheets.ClearValuesRequest{}
	_, err = srv.Spreadsheets.Values.Clear(spreadsheetID, r, rb).Do()
	if err != nil {
		log.Fatal(err)
	}

}

func BatchClear(client *http.Client, spreadsheetID string, ranges []string) {
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	rb := &sheets.BatchClearValuesRequest{
		Ranges: ranges,
	}
	_, err = srv.Spreadsheets.Values.BatchClear(spreadsheetID, rb).Do()
	if err != nil {
		log.Fatal(err)
	}

}
