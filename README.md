# Sync-It

Sync-It is a Go application that uploads files from a specified local directory to a Google Cloud Storage (GCP) bucket. The application checks for file modifications and only uploads files that have been updated.

## Features

- Authenticate with Google Cloud Storage
- List all objects in the specified GCP bucket
- Walk through the specified local directory and check for modified files
- Upload only modified files to the GCP bucket

## Prerequisites

- Go 1.16 or higher
- Google Cloud SDK installed and authenticated
- A Google Cloud Storage bucket


## Usage

1. Set up your Google Cloud authentication:

    Ensure you have the Google Cloud SDK installed and authenticated. You can authenticate using the following command:

    ```bash
    gcloud auth application-default login
    ```
    Using a Service Account
    Create a service account in the Google Cloud Console and download the JSON key file.

    ```bash
    Set the GOOGLE_APPLICATION_CREDENTIALS environment variable to point to the downloaded key file:
    ```

2. Run the application:

    ```bash
    go run main.go gcp.go -GCP_BUCKET=<your-gcp-bucket-name> -DIR=<your-local-directory>
    ```

    Replace `<your-gcp-bucket-name>` with the name of your GCP bucket and `<your-local-directory>` with the path to the local directory you want to sync.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or issues, please open an issue on the [GitHub repository](https://github.com/yourusername/sync-it/issues).
