package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"google.golang.org/api/calendar/v3"
)

const lineNotifyURL = "https://notify-api.line.me/api/notify"

func cal2message(event *calendar.Event) string {
	if event == nil {
		return "直近のイベントが設定されていません"
	}
	var when string
	if event.Start.DateTime != "" {
		when = event.Start.DateTime
	} else {
		when = event.Start.Date
	}

	return fmt.Sprintf(`%s は %s の日です。
%s`, when, event.Summary, event.Description)
}

func lineNotify(event *calendar.Event) {
	token := os.Getenv("LINE_NOTIFY_TOKEN")

	v := url.Values{"message": {cal2message(event)}}

	req, err := http.NewRequest(http.MethodPost, lineNotifyURL, strings.NewReader(v.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(string(b))
	}
}
