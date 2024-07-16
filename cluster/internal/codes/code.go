package codes

import (
	"github.com/dobyte/due/v2/codes"
)

var (
	OK                         = codes.OK
	Canceled                   = codes.Canceled
	Unknown                    = codes.Unknown
	InvalidArgument            = codes.InvalidArgument
	DeadlineExceeded           = codes.DeadlineExceeded
	NotFound                   = codes.NotFound
	InternalError              = codes.InternalError
	Unauthorized               = codes.Unauthorized
	IllegalInvoke              = codes.IllegalInvoke
	IllegalRequest             = codes.IllegalRequest
	AccountExists              = codes.NewCode(10, "account exists")
	AccountNotExists           = codes.NewCode(11, "account not exists")
	IncorrectAccountOrPassword = codes.NewCode(12, "incorrect account or password")
)
