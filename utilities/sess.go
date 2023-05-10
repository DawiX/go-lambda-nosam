package utilities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitSession(profile string, region string) (*session.Session, error) {
	if profile != "" && region != "" {
		return session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable, // Must be set to enable for SSO handled profiles
			Profile:           profile,
			Config: aws.Config{
				Region:                        aws.String(region),
				CredentialsChainVerboseErrors: aws.Bool(true),
			},
		})
	} else {
		return session.NewSession()
	}
}
