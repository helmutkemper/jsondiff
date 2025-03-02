package diffServer

import (
	"github.com/rivo/tview"
	"strconv"
)

type formToken struct {
	token     *tview.InputField
	tokenType *tview.InputField
	status    *tview.InputField
	issuedAt  *tview.InputField
	expiresIn *tview.InputField
}

func (e *formToken) Init(token TokenData, fieldWidth int) {
	issuedAt := strconv.FormatInt(token.GetIssuedAt(), 10)
	expiresIn := strconv.FormatInt(token.GetExpiresIn(), 10)

	e.token = tview.NewInputField().
		SetLabel("Token").
		SetFieldWidth(fieldWidth).
		SetText(token.GetAccessToken()).
		SetChangedFunc(
			func(text string) {
				token.SetAccessToken(text)
			},
		)

	e.tokenType = tview.NewInputField().
		SetLabel("Token type").
		SetFieldWidth(fieldWidth).
		SetText(token.GetTokenType()).
		SetChangedFunc(
			func(text string) {
				token.SetTokenType(text)
			},
		)

	e.status = tview.NewInputField().
		SetLabel("Status").
		SetFieldWidth(fieldWidth).
		SetText(token.GetStatus()).
		SetChangedFunc(
			func(text string) {
				token.SetStatus(text)
			},
		)

	e.issuedAt = tview.NewInputField().
		SetLabel("Issued at").
		SetFieldWidth(fieldWidth).
		SetText(issuedAt).SetChangedFunc(
		func(text string) {
			var err error
			var i int64
			i, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				i = 0
			}
			token.SetIssuedAt(i)
		},
	)

	e.expiresIn = tview.NewInputField().
		SetLabel("Expires in").
		SetFieldWidth(fieldWidth).
		SetText(expiresIn).SetChangedFunc(
		func(text string) {
			var err error
			var i int64
			i, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				i = 0
			}
			token.SetExpiresIn(i)
		},
	)
}

func (e *formToken) Mount(form *tview.Form) {
	form.AddFormItem(e.token)
	form.AddFormItem(e.tokenType)
	form.AddFormItem(e.status)
	form.AddFormItem(e.issuedAt)
	form.AddFormItem(e.expiresIn)
}
