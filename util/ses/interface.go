package ses

import "github.com/aws/aws-sdk-go/service/ses"

type ISES interface {
	SendEmail(req *EmailRequest) (*ses.SendEmailOutput, error)
}
