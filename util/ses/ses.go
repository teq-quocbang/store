package ses

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	mySES "github.com/teq-quocbang/store/client/ses"
)

type SES struct {
	session *session.Session
	svc     *ses.SES
}

func NewSES() ISES {
	return SES{
		session: mySES.GetSession(),
		svc:     mySES.GetService(),
	}
}
