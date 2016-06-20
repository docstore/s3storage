package s3docstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/docstore/service"
)

var _ docstore.Storer = AwsS3{}

// AwsS3 defines a document service backed by by aws s3
type AwsS3 struct {
	AwsConfig  aws.Config
	BucketName string
	Session    *session.Session
	s3service  *s3.S3
}

// BasicAws loads aws config from system variables
func BasicAws(region, bucketName string) AwsS3 {
	config := aws.Config{
		Region: aws.String(region),
	}

	awsSession := session.New()

	svc := s3.New(awsSession, &config)

	return AwsS3{
		AwsConfig:  config,
		BucketName: bucketName,
		Session:    awsSession,
		s3service:  svc,
	}
}

// NewAws instantiates AwsS3 without system variables
func NewAws(region, bucketName string, awsSession *session.Session, awsConfig aws.Config) AwsS3 {

	svc := s3.New(awsSession, &awsConfig)

	return AwsS3{
		AwsConfig:  awsConfig,
		BucketName: bucketName,
		Session:    awsSession,
		s3service:  svc,
	}
}

func (a AwsS3) Put(obj docstore.CreateObj) (string, error) {
	putReq := s3.PutObjectInput{
		Bucket: &a.BucketName,
		Body:   obj,
		Key:    &obj.Identifier,
	}

	_, err := a.s3service.PutObject(&putReq)

	if err != nil {
		return "", err
	}

	return obj.Identifier, nil
}

func (a AwsS3) Get(fileName string) (docstore.RetrieveObj, error) {
	getRequest := s3.GetObjectInput{
		Bucket: &a.BucketName,
		Key:    &fileName,
	}
	objOutput, getError := a.s3service.GetObject(&getRequest)

	if getError != nil {
		return docstore.RetrieveObj{}, getError
	}
	doc := docstore.RetrieveObj{
		ReadCloser: objOutput.Body,
		Identifier: fileName,
	}

	return doc, nil
}
