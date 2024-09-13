package main

import (
	"flag"
	"github.com/emersion/go-imap/client"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var config Config

func init() {
	// Email config
	flag.StringVar(&config.EmailServer, "email-server", os.Getenv("EMAIL_SERVER"), "Email server (example: mail.domain.com:993)")
	flag.StringVar(&config.EmailLogin, "email-login", os.Getenv("EMAIL_LOGIN"), "The login of your email account")
	flag.StringVar(&config.EmailPassword, "email-password", os.Getenv("EMAIL_PASSWORD"), "The password of your email account")

	// Telegram config
	flag.Int64Var(&config.TelegramUserID, "telegram-userid", 0, "Please ask to telegram bot @myidbot")
	flag.DurationVar(&config.ReadTimeout, "read-timeout", 0, "Mailbox read timeout")
	flag.StringVar(&config.TelegramToken, "telegram-token", os.Getenv("TELEGRAM_TOKEN"), "Telegram bot token (https://core.telegram.org/bots/api)")

	// Other
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose/debug")

	flag.Parse()
	if config.ReadTimeout == 0 {
		var err error
		var intReadTimeout int64
		intReadTimeout, err = strconv.ParseInt(os.Getenv("READ_TIMEOUT"), 10, 64)
		if err != nil {
			log.Fatal("READ_TIMEOUT invalid int")
		}

		config.ReadTimeout = time.Duration(intReadTimeout)
	}

	if config.TelegramUserID == 0 {
		var err error
		config.TelegramUserID, err = strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)
		if err != nil {
			log.Fatal("TELEGRAM_USER_ID invalid int")
		}
	}

	if config.Verbose {
		checkConfig := reflect.ValueOf(config)
		typeOfS := checkConfig.Type()
		for i := 0; i < checkConfig.NumField(); i++ {
			log.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, checkConfig.Field(i).Interface())
		}
	}

	userID.ID = config.TelegramUserID
}

func dialClient() (*client.Client, func()) {
	// Let's assume config is an IMAP client
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(config.EmailServer, nil)
	if err != nil {
		log.Fatal(err)
	}

	if config.Verbose {
		c.SetDebug(os.Stdout)
	}

	log.Println("Connected")

	// Login
	if err := c.Login(config.EmailLogin, config.EmailPassword); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select a mailbox
	if _, err := c.Select("INBOX", false); err != nil {
		log.Fatal(err)
	}

	return c, func() {
		err := c.Logout() // Don't forget to logout
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	c, cleanup1 := dialClient()
	defer cleanup1()

	mc := IdleMailClient{Client: c}
	mc.ListenForEmails()
}
