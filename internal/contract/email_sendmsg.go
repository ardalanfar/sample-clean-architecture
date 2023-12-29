package contract

import (
	"context"
)

type SendMassage interface {
	SendEmail(context.Context, string) error
	BuildMessage()string
}