package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/makecall", makeCallHandler)
	http.ListenAndServe(":8080", nil)
}

func makeCallHandler(w http.ResponseWriter, r *http.Request) {
	// Get the caller ID from the query string
	callerID := r.URL.Query().Get("callerID")
# http://localhost:8080/makecall?callerID=My+CallerID+<1234567890>
	// Write the call file
	err := writeCallFile(callerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Call file created")
}

func writeCallFile(callerID string) error {
	// Open the call file for writing
	file, err := os.Create("/var/spool/asterisk/outgoing/mycall.call")
	if err != nil {
		return err
	}
	defer file.Close()

func writeCallFile(callerID string) (string, error) {
	// Open the call file for writing
	file, err := os.Create("/var/spool/asterisk/outgoing/mycall.call")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write the contents of the call file, including the caller ID
	fmt.Fprintf(file, "Channel: SIP/mytrunk/mydestination\n")
	fmt.Fprintf(file, "CallerID: %s\n", callerID)
	fmt.Fprintf(file, "Application: Playback\n")
	fmt.Fprintf(file, "Data: sound_file_to_play\n")
	fmt.Fprintf(file, "MaxRetries: 2\n")
	fmt.Fprintf(file, "RetryTime: 60\n")
	fmt.Fprintf(file, "WaitTime: 30\n")
	fmt.Fprintf(file, "Context: mycontext\n")
	fmt.Fprintf(file, "Extension: myextension\n")
	fmt.Fprintf(file, "Priority: 1\n")
	fmt.Fprintf(file, "\n")

	// Return the caller ID that was written to the file
	return callerID, nil
}
