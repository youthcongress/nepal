package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Connection(){
	// MinIO Configuration
	endpoint := "94.136.185.141:7000" 								// Change to your MinIO server address
	accessKey := "WpxEmNBd7N33ldrYtBai"                				// Change to your MinIO access key
	secretKey := "9MOBEv6jUGaWL0IjDG2lHmRJtVKvnccarzYjnSu2"         // Change to your MinIO secret key
	bucketName := "youthcongressnepal"           					// Change to your desired bucket name

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // Change to true if using HTTPS
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to MinIO:", err)
	}	

	// Check if the bucket exists
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatal("❌ Error checking bucket:", err)
	}

	if exists {
		fmt.Println("✅ Successfully connected to bucket:", bucketName)
	} else {
		fmt.Println("⚠️ Bucket does not exist:", bucketName)
	}
}