package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
)

type ReadClient struct {
	Client *client.Client
	Ch     chan imap.MailboxStatus
}

func (rc *ReadClient) Loop() {
	for upd := range rc.Ch {
		rc.Read(upd)
	}
}

func (rc *ReadClient) Read(mbox imap.MailboxStatus) {
	messages := make(chan *imap.Message, 10)
	go rc.ReadMessages(messages)

	seq := new(imap.SeqSet)
	seq.AddNum(mbox.Messages)
	err := rc.Client.Fetch(seq, []imap.FetchItem{imap.FetchRFC822}, messages)
	if err != nil {
		// Deal with error
		log.Println("Fetch error: ", err)
	}
}

func (rc *ReadClient) ReadMessages(messages chan *imap.Message) {
	for email := range messages {
		log.Println("New mail: ")
		for _, body := range email.Body {
			go MailProcessing(rc.EmailBodyParse(body), email)
		}
	}
}

func (rc *ReadClient) EmailBodyParse(r io.Reader) []byte {
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			log.Println("got text: ", string(b))
			return b
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			// Unused in my case
			log.Println("got attachment: ", filename)
		}
	}
	return []byte{}
}
