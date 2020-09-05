package flows

import (
	htmlform "github.com/ory/kratos/selfservice/form"
)

type Page struct {
	Title	string
	ID         string `json:"id"`
	ExpiresAt  string `json:"expires_at"`
	IssueAt    string `json:"issued_at"`
	RequestUrl string `json:"request_url"`
	Force      bool   `json:"forced"`
	Methods    Method
}

type Method struct {
	Oidc     MethodConfig `json:"oidc,omitempty"`
	Password MethodConfig `json:"password,omitempty"`
}

type MethodConfig struct {
	Method string     `json:"method"`
	Config FormConfig `json:"config"`
}

type FormConfig struct {
	Action string           `json:"action"`
	Method string           `json:"method"`
	Fields []htmlform.Field `json:"fields"`
	Errors []ErrorMessage   `json:"errors,omitempty"`
	Messages []AttachedMessages `json:"messages,omitempty`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type AttachedMessages struct {
	Id int `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	Type string `json:"type,omitempty"`
}

type Profile struct {
	Title string
	ID         string     `json:"id"`
	ExpiresAt  string     `json:"expires_at"`
	IssueAt    string     `json:"issued_at"`
	RequestUrl string     `json:"request_url"`
	Form       FormConfig `json:"form"`
	Identity   Identity   `json:"identity`
}

type Identity struct {
	Id              string        `json:"id,omitempty"`
	TraitsSchemaId  string        `json:"traits_schema_id,omitempty"`
	TraitsSchemaUrl string        `json:"traits_schema_url,omitempty"`
	Traits          Traits `json:"traits,omitempty"`
}

type Traits struct {
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"email,omitempty"`
}