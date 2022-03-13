package send

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

// general account info, always get it from environment variables
type ValEmail struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}


var (
	servername string // port is needed
	username   string
	password   string
	name       string
)

// compose message according to "from, to, subject, body"
func composeMsg(from string, to string, subject string, body string) (message string) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["Content-Type"] = "text/plain; charset=UTF-8"
	// Setup message
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	return
}

// send email over SSL
func TlsSend(toAddr string, subject string, body string, emailConfig *ValEmail) (err error) {
	servername = emailConfig.Host
	username = emailConfig.Account
	password = emailConfig.Password
	name = emailConfig.Name

	host, _, _ := net.SplitHostPort(servername)
	// get SSL connection
	conn, err := tls.Dial("tcp", servername, nil)
	if err != nil {
		return
	}
	// create new SMTP client
	smtpClient, err := smtp.NewClient(conn, host)
	if err != nil {
		return
	}
	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, host)
	// auth the smtp client
	err = smtpClient.Auth(auth)
	if err != nil {
		return
	}
	// set To && From address, note that from address must be same as authorization user.
	from := mail.Address{name, username}
	to := mail.Address{"", toAddr}
	err = smtpClient.Mail(from.Address)
	if err != nil {
		return
	}
	err = smtpClient.Rcpt(to.Address)
	if err != nil {
		return
	}
	// Get the writer from SMTP client
	writer, err := smtpClient.Data()
	if err != nil {
		return
	}
	// compose message body
	message := composeMsg(from.String(), to.String(), subject, body)
	// write message to recp
	_, err = writer.Write([]byte(message))
	if err != nil {
		return
	}
	// close the writer
	err = writer.Close()
	if err != nil {
		return
	}
	// Quit sends the QUIT command and closes the connection to the server.
	smtpClient.Quit()
	return nil
}

func AsyncTlsSend(toAddr string, subject string, body string, emailConfig *ValEmail, handle func(err error)) error {
	go func() {
		err := TlsSend(toAddr, subject, body, emailConfig)
		handle(err)
	}()
	return nil
}
