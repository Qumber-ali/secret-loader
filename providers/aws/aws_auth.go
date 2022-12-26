package aws_secrets_manager

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Auth(profile string) *session.Session {

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	return sess
}
