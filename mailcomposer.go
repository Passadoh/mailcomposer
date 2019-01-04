// Package mailcomposer takes "html/template" formated html template + data, and returns a composed mail.
package mailcomposer

import (
	"bytes"
	"html/template"
)

// Parameters consists of dynamic data used to a composed an email.
type Parameters struct {
	Data interface{}
	File string
}

// Email is the end-result.
type Email struct {
	Body    string
	Subject string
}

// Create turns process data and returns a composed email template.
func Create(p Parameters) (Email, error) {
	e := Email{}

	err := e.parseTemplate(p)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (e *Email) parseTemplate(p Parameters) error {
	t, err := template.ParseFiles(p.File)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, p.Data); err != nil {
		return err
	}
	e.Body = buf.String()
	return nil
}
