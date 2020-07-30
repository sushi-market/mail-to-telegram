package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"time"
)

type IdleMailClient struct {
	Client    *client.Client
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

			log.Println("First time index: ", ec.Index)
		}

		if mb.Messages <= ec.Index {
			continue
		}

		for ec.Index < mb.Messages {
			ec.Index += 1
			log.Println("Index update: ", ec.Index)

			(&ReadClient{Client: ec.Client}).Read(imap.MailboxStatus{Messages: ec.Index})
			//ec.UpdatesCh <- imap.MailboxStatus{Messages: ec.Index}
		}
	}

	log.Fatal("idle main client exited")
}
