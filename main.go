package main

import (
	"log"
	"os"
	"time"
)

func notify() {
	srv := getService()

	calID := os.Getenv("GOOGLE_CALENDAR_ID")

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calID).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	if len(events.Items) > 0 {
		for _, i := range events.Items {
			lineNotify(i)
			return
		}
	} else {
		lineNotify(nil)
	}
}

func main() {
	subCmd := "notify"

	if len(os.Args) > 1 {
		subCmd = os.Args[1]
	}

	switch subCmd {
	case "notify":
		notify()
	case "auth":
		auth()
	}
}
