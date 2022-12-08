package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"3390142252@qq.com",
		"axgppxpfftgfdbcg",
		"smtp.qq.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	fmt.Println("finish1")
	msg := []byte("From: 3390142252@qq.com\n" +
		"To: kangningwang254@gmail.com\n" +
		"Subject: 测试!\n" +
		"\n" +
		"This is the email body.\n")
	err := smtp.SendMail(
		"smtp.qq.com:25",
		auth,
		"3390142252@qq.com",
		[]string{"kangningwang254@gmail.com"},
		msg,
	)
	fmt.Println("finish2")
	if err != nil {
		log.Fatal(err)
	}
}
