package main

import (
	"github.com/emersion/go-imap"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var (
	b      *tb.Bot
	userID = &tb.Chat{ID: config.TelegramUserID}
)

type MessageFmt struct {
	Subject string
	Link    string
}

func init() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		Token:  config.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("telegram: ", err)
		return
	}
}

func MessageFormatting(msgFmt MessageFmt, mail *imap.Message) string {
	from := ""
	if mail.Envelope != nil {
		arr := mail.Envelope.From
		if len(arr) > 0 && arr[0] != nil {
			from = arr[0].Address()
		}
	}

	text := ""
	text += "Subject: " + msgFmt.Subject
	text += "\nFrom: " + from
	text += "\n"
	text += "\n" + msgFmt.Link
	return text
}
