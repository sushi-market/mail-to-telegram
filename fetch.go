package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"time"
)

type IdleMailClient struct {
	Client    *client.Client
	UpdatesCh chan imap.MailboxStatus
	Index     uint32
}

func (ec *IdleMailClient) ListenForEmails() {
	const tick = time.Second * 15

	firstTime := true

	for range time.NewTicker(tick).C {
		mb, err := ec.Client.Select("INBOX", false)
		if err != nil {
			log.Println("Selecting mailbox error: ", err)
			continue
		}

		if firstTime {
			firstTime = false
			ec.Index = mb.Messages
		}

		ec.Index -= 1

		if mb.Messages <= ec.Index {
			continue
		}

		for ec.Index < mb.Messages {
			ec.Index += 1
			ec.UpdatesCh <- imap.MailboxStatus{Messages: ec.Index}
		}
	}

	log.Fatal("idle main client exited")
}
