package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	fmt.Println("Golang email app running")
	email()
}

func email() {
	// sender data
	from := os.Getenv("FromEmailAddr") // ex: "john.doe@gmail.com"
	pass := os.Getenv("SMTPpwd")
	// receiver address
	toEmail := os.Getenv("ToEmailAddr")
	to := []string{toEmail} // we're just using one email, but this is a slice string, so we can send to multiple emails
	// SMTP - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Our Golang Email\n"
	body := "Our first email!"
	message := []byte(subject + body) // slice of bytes
	// authentication data
	auth := smtp.PlainAuth("", from, pass, host)
	// send email
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Check your email!")

}

/* Notes:
- If you have 2 step authentication enabled, you should use "Application-specific passwords" that will be used to log in to your account
- Documentation: https://pkg.go.dev/net/smtp
*/
