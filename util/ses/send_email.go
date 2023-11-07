package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/pkg/errors"

	"github.com/teq-quocbang/store/util/myerror"
)

const CharSetUTF8 string = "UTF-8"

type EmailRequest struct {
	Sender  *string
	To      []*string
	Cc      []*string
	Bcc     []*string
	Content *string
	Subject *string
	IsHTML  bool
}

func (s SES) SendEmail(req *EmailRequest) (*ses.SendEmailOutput, error) {
	var (
		body    = ses.Body{}
		content = ses.Content{
			Charset: aws.String(CharSetUTF8),
			Data:    req.Content,
		}
	)

	if req.IsHTML {
		body.Html = &content
	} else {
		body.Text = &content
	}

	input := &ses.SendEmailInput{
		Source: req.Sender,
		Destination: &ses.Destination{
			ToAddresses:  req.To,
			CcAddresses:  req.Cc,
			BccAddresses: req.Bcc,
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String(CharSetUTF8),
				Data:    req.Subject,
			},
			Body: &body,
		},
	}

	result, err := s.svc.SendEmail(input)
	if err != nil {
		var aErr awserr.Error

		if ok := errors.As(err, &aErr); ok {
			return nil, myerror.ErrSendEmail(aErr)
		}

		return nil, myerror.ErrSendEmail(err)
	}

	return result, nil
}
