package send_email

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

// 发email demo
func TestSendDemo(t *testing.T) {
	e := email.NewEmail()
	e.From = "AAA <sendhanee@outlook.com>"
	e.To = []string{"worldelitecao@foxmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//e.Attach()
	err := e.SendWithStartTLS("smtp.office365.com:587",
		smtp.PlainAuth("", "sendhanee@outlook.com", "FAfafayoujian123", "smtp.office365.com"),
		&tls.Config{
			Rand:                        nil,
			Time:                        nil,
			Certificates:                nil,
			GetCertificate:              nil,
			GetClientCertificate:        nil,
			GetConfigForClient:          nil,
			VerifyPeerCertificate:       nil,
			RootCAs:                     nil,
			NextProtos:                  nil,
			ServerName:                  "smtp.office365.com",
			ClientAuth:                  0,
			ClientCAs:                   nil,
			InsecureSkipVerify:          true,
			CipherSuites:                nil,
			PreferServerCipherSuites:    false,
			SessionTicketsDisabled:      false,
			SessionTicketKey:            [32]byte{},
			ClientSessionCache:          nil,
			MinVersion:                  0,
			MaxVersion:                  0,
			CurvePreferences:            nil,
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                nil,
		},
		//&tls.Config{ServerName: "smtp.office365.com", InsecureSkipVerify: true},
	)
	if err != nil {
		t.Fatal(err)
	}
}
