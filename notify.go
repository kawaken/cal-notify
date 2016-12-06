package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	calendar "google.golang.org/api/calendar/v3"
)

type notifyEvent struct {
	*calendar.Event
	when string
}

func (n *notifyEvent) String() string {
	if n == nil {
		return "直近のイベントが設定されていません"
	}

	return fmt.Sprintf(`%s は %s の日です。`, n.when, n.Summary)
}

func (n *notifyEvent) LongString() string {
	if n == nil {
		return "直近のイベントが設定されていません"
	}

	return fmt.Sprintf(`%s は %s の日です。
%s`, n.when, n.Summary, n.Description)
}

const lineNotifyURL = "https://notify-api.line.me/api/notify"

func lineNotify(event *notifyEvent) {
	token := os.Getenv("LINE_NOTIFY_TOKEN")

	v := url.Values{"message": {event.String()}}

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
