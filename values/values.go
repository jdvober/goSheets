package values

import (
	"context"
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

// Sample Values for testing.  If being called from
/* var majorDimension string = "COLUMNS"
 * var spreadsheetID string = "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
 * var r string = "F4:F5" */

// Update adds values to a Google Sheet
func Update(client *http.Client, spreadsheetID string, r string) {

	ctx := context.Background()

	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create sheets Client %v", err)
	}

	var vr sheets.ValueRange
	values := []interface{}{"One", "Two"}
	vr.Values = append(vr.Values, values)
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, r, &vr).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to post data to sheet. %v", err)
	}
}
