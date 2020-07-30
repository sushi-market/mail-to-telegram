package main

import (
	"github.com/emersion/go-imap"
	"gopkg.in/tucnak/telebot.v2"
	"html"
	"log"
)

func MailProcessing(msg []byte, mail *imap.Message) {
	msgFmt := MessageFmt{
		Subject: "",
		Link:    "N/A",
	}
	if mail.Envelope != nil {
		msgFmt.Subject = mail.Envelope.Subject
	}
	msgFmt.Link = MailBodyProcessing(string(msg))
	tgBody := MessageFormatting(msgFmt)
	tgBody = html.EscapeString(tgBody)
	_, err := b.Send(userID, tgBody, telebot.ModeHTML)
	if err != nil {
		log.Fatal("telegram: ", err)
		return
	}
}

func MailBodyProcessing(msg string) string {
	return msg
}
