package aws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	S3Client   *s3.S3
	bucketName string
	region     string
)

func NewS3(bname, re string) error {
	bucketName = bname
	region = re

	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		return err
	}

	config := &aws.Config{
		Endpoint:         aws.String(fmt.Sprintf("https://s3-%s.amazonaws.com", region)),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
	}

	sess, err := session.NewSession(config)
	if err != nil {
		fmt.Println("failed to create session")
		return err
	}

	S3Client = s3.New(sess)

	return nil
}

func UploadToS3(filePath string, key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)

	// read file content to buffer
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer) // convert to io.ReadSeeker type
	fileType := http.DetectContentType(buffer)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName), // required
		Key:           aws.String(key),        // required
		ACL:           aws.String("public-read"),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"), //required
		},
		// see more at http://godoc.org/github.com/aws/aws-sdk-go/service/s3#S3.PutObject
	}

	_, err = S3Client.PutObject(params)
	if err != nil {
		return err
	}

	return nil

}

func UploadToS3ByMultiPart(file multipart.File, fileName string) error {

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName), // required
		Key:           aws.String(fileName),   // required
		ACL:           aws.String("public-read"),
		Body:          file,
		ContentLength: aws.Int64(int64(len(bs))),
		ContentType:   aws.String("application/octet-stream"),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"), //required
		},
		// see more at http://godoc.org/github.com/aws/aws-sdk-go/service/s3#S3.PutObject
	}

	_, err = S3Client.PutObject(params)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFromS3(key string) error {

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName), // Required
		Key:    aws.String(key),        // Required
	}
	_, err := S3Client.DeleteObject(params)

	return err
}
