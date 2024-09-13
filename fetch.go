package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"time"
)

type IdleMailClient struct {
	Client *client.Client
	Index  uint32
}

func (ec *IdleMailClient) ListenForEmailsTick() {
	mb, err := ec.Client.Select("INBOX", false)
	if err != nil {
		log.Fatal("Selecting mailbox error: ", err)
		return
	}

	if mb.UnseenSeqNum <= 0 {
		return
	}

	ec.Index = mb.UnseenSeqNum

	(&ReadClient{Client: ec.Client}).Read(imap.MailboxStatus{Messages: ec.Index})
}

func (ec *IdleMailClient) ListenForEmails() {
	var tick = time.Second * config.ReadTimeout

	ec.ListenForEmailsTick()

	for range time.NewTicker(tick).C {
		ec.ListenForEmailsTick()
	}

	log.Fatal("idle main client exited")
}
