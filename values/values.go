package values

import (
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

// Update adds values to a Google Sheet
func Update(client *http.Client, spreadsheetID string, r string, values []interface{}) {

	/* ctx := context.Background() */

	// Authenticate
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create sheets Client %v", err)
	}

	var vr sheets.ValueRange
	/* values := []interface{}{"One", "Two"} */
	vr.Values = append(vr.Values, values)
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, r, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to post data to sheet. %v", err)
	}
}
