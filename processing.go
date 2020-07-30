package main

import (
	"github.com/emersion/go-imap"
	"gopkg.in/tucnak/telebot.v2"
	"html"
	"jaytaylor.com/html2text"
	"log"
	"strings"
	"time"
)

func MailProcessing(msg []byte, mail *imap.Message) {
	if mail == nil {
		mail = &imap.Message{}
	}

	msgFmt := MessageFmt{
		Subject: "",
		Link:    "N/A",
	}
	if mail.Envelope != nil {
		msgFmt.Subject = mail.Envelope.Subject
	}
	msgFmt.Link = MailBodyProcessing(string(msg))

	emailBody := MessageFormatting(msgFmt, mail)
	arr := splitBodyByLimit(emailBody, 4000)

	for _, tgBody := range arr {
		tgBody = html.EscapeString(tgBody)
		log.Println("Sending message to telegram, len=", len(tgBody))
		_, err := b.Send(userID, tgBody, telebot.ModeHTML)
		time.Sleep(time.Second)
		if err != nil {
			log.Fatal("telegram: ", err)
			return
		}
	}
}

func splitBodyByLimit(text string, limit int) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	var res []string

	current := ""

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		shortLines := splitLongLine(line, limit)

		for _, short := range shortLines {
			newCurrent := current + "\n" + line
			if len(newCurrent) > limit {
				res = append(res, current)
				current = short
			} else {
				current = newCurrent
			}
		}
	}

	if current != "" {
		res = append(res, current)
	}

	return res
}

func splitLongLine(line string, limit int) []string {
	tmp := []rune(line)

	var shortLines []string
	for len(tmp) > 0 {
		r := limit
		if len(tmp) < r {
			r = len(tmp)
		}

		shortLines = append(shortLines, string(tmp[:r]))
		tmp = tmp[r:]
	}

	return shortLines
}

func MailBodyProcessing(msg string) string {
	text, err := html2text.FromString(msg, html2text.Options{PrettyTables: true})
	if err != nil {
		log.Println("failed to do html2text: ", err)
		return msg
	}
	return text
}
