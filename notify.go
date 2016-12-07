package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	calendar "google.golang.org/api/calendar/v3"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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

func lineNotify(events []*notifyEvent) {
	token := os.Getenv("LINE_NOTIFY_TOKEN")

	messages := []string{""}
	for _, e := range events {
		messages = append(messages, e.String())
	}

	v := url.Values{}
	v.Set("stickerPackageId", "2")
	v.Set("stickerId", randomSticker())
	v.Set("message", strings.Join(messages, "\n"))

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

func randomSticker() string {
	max, min := 179, 140
	return strconv.Itoa(rand.Intn(max-min) + min)
}
