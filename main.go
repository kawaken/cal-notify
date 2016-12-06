package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/now"
	calendar "google.golang.org/api/calendar/v3"
)

func init() {
	now.FirstDayMonday = true
}

func getEventTime(e *calendar.Event) time.Time {
	if e.Start.DateTime != "" {
		t, _ := time.Parse(time.RFC3339, e.Start.DateTime)
		return t
	}

	t, _ := time.Parse("2006-01-02", e.Start.Date)
	return t
}

func notify() {
	srv := getService()

	calID := os.Getenv("GOOGLE_CALENDAR_ID")
	if calID == "" {
		log.Fatal("GOOGLE_CALENDAR_ID is not defined")
	}

	timeNow := time.Now()
	strToday := timeNow.Format(time.RFC3339)
	timeTomorrow := now.BeginningOfDay().AddDate(0, 0, 1)
	timeFriday := now.EndOfWeek().AddDate(0, 0, -2)
	strFriday := timeFriday.Format(time.RFC3339)

	events, err := srv.Events.List(calID).ShowDeleted(false).
		SingleEvents(true).TimeMin(strToday).TimeMax(strFriday).MaxResults(15).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	if len(events.Items) == 0 {
		lineNotify(nil)
		return
	}

	for _, i := range events.Items {
		et := getEventTime(i)
		if et.Before(timeNow) {
			continue
		}
		var when string
		switch et.Day() {
		case timeNow.Day():
			when = "今日"
		case timeTomorrow.Day():
			when = "明日"
		case timeFriday.Day():
			when = fmt.Sprintf("%d(金)", et.Day())
		default:
			continue
		}

		lineNotify(&notifyEvent{Event: i, when: when})
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
