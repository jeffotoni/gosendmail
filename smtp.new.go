package gosendmail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strings"
)

func New() *Sender {
	if len(host) == 0 ||
		len(username) == 0 ||
		len(password) == 0 ||
		len(port) == 0 {
		return nil
	}

	auth := smtp.PlainAuth("", username, password, host)
	return &Sender{auth}
}

func (s *Sender) Send(m *Message) error {
	if insecureSkip == "true" {
		c, err := smtp.Dial(host + ":" + port)
		if err != nil {
			return errors.New("error Dial:" + err.Error())
		}
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			//ServerName:         host,
		}
		if err := c.StartTLS(tlsconfig); err != nil {
			return err
		}
	}

	if len(from) == 0 {
		from = username
	}
	return smtp.SendMail(fmt.Sprintf("%s:%s", host, port), s.auth, from, m.To, m.ToBytes())
}

func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	// fmt.Println("fileName ", fileName)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	//mimeHeaders := "MIME-Version: 1.0;\n"
	buf.WriteString(mimeHeaders)
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		// buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	// log.Println("body:", m.Body)
	buf.WriteString(m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}
		buf.WriteString("--")
	}
	// log.Println(string(buf.Bytes()))
	return buf.Bytes()
}
