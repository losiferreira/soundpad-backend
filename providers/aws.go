package providers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type Aws struct {
	S3Client *s3.S3
}

func NewAws() *Aws {
	return &Aws{}
}

func (a *Aws) Setup() *Aws {
	// Set up AWS credentials
	awsConfig := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_KEY"),
			os.Getenv("AWS_SECRET"),
			""),
	}

	// Create an AWS session and S3 client
	awsSession := session.Must(session.NewSession(awsConfig))
	a.S3Client = s3.New(awsSession)

	return a
}
