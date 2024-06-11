package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Printf("[Main] Starting App")

	localDirectory := flag.String("DIR", os.Getenv("DIR"), "Directory to watch")
	gcpBucket := flag.String("GCP_BUCKET", os.Getenv("GCP_BUCKET"), "GCP Bucket Name")

	flag.Parse()

	if *gcpBucket != "" {
		ctx, client := gcp_authenicate()
		gcp_list_bucket_controls(ctx, client, *gcpBucket)

		err := filepath.Walk(*localDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("[Main] Error accessing directory path %s: %v", path, err)
				return nil
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			if gcp_needUpload(ctx, client, path, *gcpBucket) {
				gcp_uploadFile(ctx, client, path, *gcpBucket)
			}
			return nil
		})
		if err != nil {
			log.Printf("[Main] Failed to walk directory: %v", err)
		}
	}
}
