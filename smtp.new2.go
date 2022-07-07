package gosendmail

import (
	"crypto/tls"
	"errors"
	"net"
	"net/smtp"
)

var smtpHost string

func New2() *Sender {
	if len(host) == 0 ||
		len(username) == 0 ||
		len(password) == 0 ||
		len(port) == 0 ||
		len(from) == 0 {
		return nil
	}
	smtpHost = host + ":" + port
	auth := LoginAuth(username, password)
	return &Sender{auth}
}

type loginAuth struct {
	username string
	password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func (s *Sender) Send2(m *Message) error {
	conn, err := net.Dial("tcp", smtpHost)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{}
	if insecureSkip != "true" {
		tlsconfig = &tls.Config{
			ServerName: host,
		}
	} else if insecureSkip != "true" {
		tlsconfig = &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: true,
		}
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		return err
	}

	if err = c.Auth(s.auth); err != nil {
		return err
	}

	return smtp.SendMail(smtpHost, s.auth, from, m.To, m.ToBytes())
}
