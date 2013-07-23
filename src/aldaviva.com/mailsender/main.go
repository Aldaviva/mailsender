package main

import "net/smtp"
import "fmt"
import "time"
import "os"

const SMTP_SERVER = "email.aldaviva.com:25"
const RFC_2822_DATE_FORMAT = "2 Jan 2006 15:04:05 -0700"

func main() {
	torrentName := os.Args[1]

	subject := "Downloaded " + torrentName
	from := "forseti@aldaviva.com"
	to := "ben@aldaviva.com"
	date := time.Now().Format(RFC_2822_DATE_FORMAT)
	body := "From: "+from+"\r\nTo: "+to+"\r\nDate: "+date+"\r\nSubject: "+subject+"\r\n\r\n'"+torrentName+"' finished torrenting"
	fmt.Printf("Sending mail (subject = %s, from = %s, to = %s, body = %s)\n", subject, from, to, body)

	sendError := smtp.SendMail(SMTP_SERVER, nil, from, []string{to}, []byte(body))

	if sendError != nil {
		fmt.Printf("Unable to send: %s\n", sendError)
	} else {
		fmt.Printf("Sent.\n")
	}
}