package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func gcp_authenicate() (context.Context, *storage.Client) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, client
}

func gcp_list_bucket_controls(ctx context.Context, client *storage.Client, bucketName string) {

	bucket := client.Bucket(bucketName)

	it := bucket.Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(attrs.Name, attrs.ContentType)
	}
}

func gcp_needUpload(ctx context.Context, client *storage.Client, localFilePath string, bucketName string) bool {
	bucket := client.Bucket(bucketName)

	fileInfo, err := os.Stat(localFilePath)
	if err != nil {
		log.Printf("[gcp_needUpload] Failed to get local file info: %v", err)
		return false
	}
	localModTime := fileInfo.ModTime()

	attrs, err := bucket.Object(localFilePath).Attrs(ctx)
	if err != nil {
		return true
	}

	gcpModTime := attrs.Updated

	return localModTime.After(gcpModTime)
}

func gcp_uploadFile(ctx context.Context, client *storage.Client, localFilePath string, bucketName string) {
	bucket := client.Bucket(bucketName)

	file, err := os.Open(localFilePath)
	if err != nil {
		log.Printf("[gcp_uploadFile] Failed to open local file: %v", err)
		return
	}
	defer file.Close()

	wc := bucket.Object(localFilePath).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Printf("[gcp_uploadFile] Failed to upload file: %v", err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Printf("[gcp_uploadFile] Failed to close writer: %v", err)
		return
	}

	log.Printf("File '%s' uploaded successfully", localFilePath)
}
