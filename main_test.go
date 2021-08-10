package gosendemail

import (
	"net/smtp"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Sender
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_Send(t *testing.T) {
	type fields struct {
		auth smtp.Auth
	}
	type args struct {
		m *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				auth: tt.fields.auth,
			}
			if err := s.Send(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Sender.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	type args struct {
		s string
		b string
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.s, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_AttachFile(t *testing.T) {
	type fields struct {
		To          []string
		CC          []string
		BCC         []string
		Subject     string
		Body        string
		Attachments map[string][]byte
	}
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				To:          tt.fields.To,
				CC:          tt.fields.CC,
				BCC:         tt.fields.BCC,
				Subject:     tt.fields.Subject,
				Body:        tt.fields.Body,
				Attachments: tt.fields.Attachments,
			}
			if err := m.AttachFile(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Message.AttachFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessage_ToBytes(t *testing.T) {
	type fields struct {
		To          []string
		CC          []string
		BCC         []string
		Subject     string
		Body        string
		Attachments map[string][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				To:          tt.fields.To,
				CC:          tt.fields.CC,
				BCC:         tt.fields.BCC,
				Subject:     tt.fields.Subject,
				Body:        tt.fields.Body,
				Attachments: tt.fields.Attachments,
			}
			if got := m.ToBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run ^TestMessage_SendEmail$
func TestMessage_SendEmail(t *testing.T) {
	sender := New()
	if sender == nil {
		t.Error("Error New() check the required fields: EMAIL_HOST,EMAiL_USERNAME,EMAIL_PASSWORD,EMAIL_PORT")
		return
	}

	type fields struct {
		To          []string
		CC          []string
		BCC         []string
		Subject     string
		Body        string
		Attachments string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{"sendemail_", fields{
			To:          []string{"jeff.otoni@gmail.com", "jeff.otoni@s3wf.com.br"},
			CC:          []string{"jeffotoni@gmail.com"},
			BCC:         []string{"jefferson.lima@engdb.com.br"},
			Subject:     "Test email pkg gosendmail",
			Body:        "Este teste é somente plain/text e não usa ainda HTML",
			Attachments: "./websocket.png",
		}, nil},
	}
	nill := make(map[string][]byte)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				To:          tt.fields.To,
				CC:          tt.fields.CC,
				BCC:         tt.fields.BCC,
				Subject:     tt.fields.Subject,
				Body:        tt.fields.Body,
				Attachments: nill,
			}
			m.AttachFile(tt.fields.Attachments)
			if got := sender.Send(m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
