package main

import (
  "fmt"
  "net/smtp"
  "os"
  "strings"
)

var (
  hostname = "mailhog" // docker-composeで指定したサービス名
  port     = 1025 // mailhogのSMPTポート
  username = "test@test.com"
  password = "password"
  from     = "from@test.com"
  subject  = "hello"
  body     = "Hello World!"
  receiver = []string{"test@test.com"}
)

func main() {
  SendEmail()
}

func SendEmail() {
  smtpServer := fmt.Sprintf("%s:%d", hostname, port)
  auth := smtp.CRAMMD5Auth(username, password)
  msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(receiver, ","), subject, body))

  if err := smtp.SendMail(smtpServer, auth, from, receiver, msg); err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}