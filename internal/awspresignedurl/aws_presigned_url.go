package awspresignedurl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Run(show bool) {
	if show {
		fmt.Println("-- AWS Pre-signed URL")

		// Load the bucket name from OS variable
		s3bucket := os.Getenv("S3_BUCKET")
		if s3bucket == "" {
			fmt.Println("Unable to load s3 bucket name from os variables")
			return
		}

		// S3 request
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})
		if err != nil {
			log.Fatal("failed to create an aws session")
		}

		svc := s3.New(sess)
		r, _ := svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(s3bucket),
			Key:    aws.String("leo.jpg"),
		})

		// Create the pre-signed URL
		url, err := r.Presign(15 * time.Minute)
		if err != nil {
			fmt.Println("Failed to generate a pre-signed url:", err)
			return
		}

		// Display the pre-signed url
		fmt.Println("Pre-signed URL:", url)
	}
}
